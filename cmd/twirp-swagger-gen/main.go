package main

import (
	"flag"
	"strings"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-bridget/twirp-swagger-gen/internal/swagger"
	"github.com/pkg/errors"
)

var _ = spew.Dump

func parse(hostname, filename, output, prefix, protoPath string) error {
	if filename == output {
		return errors.New("output file must be different than input file")
	}

	writer := swagger.NewWriter(filename, hostname, prefix, strings.Split(protoPath, ","))
	if err := writer.WalkFile(); err != nil {
		if !errors.Is(err, swagger.ErrNoServiceDefinition) {
			return err
		}
	}
	return writer.Save(output)
}

func main() {
	var (
		in         string
		out        string
		host       string
		pathPrefix string
		protoPath  string
	)
	flag.StringVar(&in, "in", "", "Input source .proto file")
	flag.StringVar(&out, "out", "", "Output swagger.json file")
	flag.StringVar(&host, "host", "api.example.com", "API host name")
	flag.StringVar(&pathPrefix, "pathPrefix", "/twirp", "Twirp server path prefix")
	flag.StringVar(&protoPath, "protoPath", "", "Path to .proto files (comma-separated))")
	flag.Parse()

	if in == "" {
		log.Fatalf("Missing parameter: -in [input.proto]")
	}
	if out == "" {
		log.Fatalf("Missing parameter: -out [output.proto]")
	}
	if host == "" {
		log.Fatalf("Missing parameter: -host [api.example.com]")
	}

	if err := parse(host, in, out, pathPrefix, protoPath); err != nil {
		log.WithError(err).Fatal("exit with error")
	}
}
