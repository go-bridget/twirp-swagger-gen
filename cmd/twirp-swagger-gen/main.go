package main

import (
	"flag"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-bridget/twirp-swagger-gen/internal/swagger"
	"github.com/pkg/errors"
)

var _ = spew.Dump

func parse(hostname, filename, output, prefix string, camelCase bool) error {
	if filename == output {
		return errors.New("output file must be different than input file")
	}

	writer := swagger.NewWriter(filename, hostname, prefix, camelCase)
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
		camelCase  bool
	)
	flag.StringVar(&in, "in", "", "Input source .proto file")
	flag.StringVar(&out, "out", "", "Output swagger.json file")
	flag.StringVar(&host, "host", "api.example.com", "API host name")
	flag.StringVar(&pathPrefix, "pathPrefix", "/twirp", "Twirp server path prefix")
	flag.BoolVar(&camelCase, "camelCase", false, "Use camelCase for field names")
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

	if err := parse(host, in, out, pathPrefix, camelCase); err != nil {
		log.WithError(err).Fatal("exit with error")
	}
}
