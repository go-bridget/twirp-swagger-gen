package main

import (
	"errors"
	"flag"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-bridget/twirp-swagger-gen/internal/swagger"
	"google.golang.org/protobuf/compiler/protogen"
)

var _ = spew.Dump

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	var flags flag.FlagSet
	hostname := flags.String("hostname", "example.com", "")
	pathPrefix := flags.String("path_prefix", "/twirp", "")
	outputSuffix := flags.String("output_suffix", ".swagger.json", "")
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			in := f.Desc.Path()
			log.Debugf("generating: %q", in)

			if !f.Generate {
				log.Debugf("skip generating: %q", in)
				continue
			}

			writer := swagger.NewWriter(in, *hostname, *pathPrefix)
			if err := writer.WalkFile(); err != nil {
				if errors.Is(err, swagger.ErrNoServiceDefinition) {
					log.Debugf("skip writing file, %s: %q", err, in)
					continue
				}
				return err
			}

			out := f.GeneratedFilenamePrefix + *outputSuffix
			g := gen.NewGeneratedFile(out, f.GoImportPath)
			if _, err := g.Write(writer.Get()); err != nil {
				return err
			}
		}
		return nil
	})
}
