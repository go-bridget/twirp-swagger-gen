package main

import (
	"flag"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"github.com/emicklei/proto"
	"github.com/go-bridget/twirp-swagger-gen/swagger"
	"github.com/pkg/errors"
)

var _ = spew.Dump

func loadProtoFile(filename, include string) (*proto.Proto, error) {
	reader, err := os.Open(filename)
	if err != nil {
		reader, err = os.Open(path.Join(include, filename))
		if err != nil {
			return nil, err
		}
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	return parser.Parse()
}

func parse(hostname, filename, output, include string) error {
	if filename == output {
		return errors.New("output file must be different than input file")
	}

	writer := swagger.NewSwaggerWriter(filename, hostname, include)

	definition, err := loadProtoFile(filename, include)
	if err != nil {
		return err
	}

	// main file for all the relevant info
	proto.Walk(definition, writer.Handlers()...)

	return writer.Save(output)
}

func main() {
	var (
		in      string
		out     string
		host    string
		include string
	)
	flag.StringVar(&in, "in", "", "Input source .proto file")
	flag.StringVar(&out, "out", "", "Output swagger.json file")
	flag.StringVar(&include, "I", "", "Extra include path for .proto files")
	flag.StringVar(&host, "host", "api.example.com", "API host name")
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

	if err := parse(host, in, out, include); err != nil {
		log.WithError(err).Fatal("exit with error")
	}
}
