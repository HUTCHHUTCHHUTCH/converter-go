package main

import (
	"github.com/asyncapi/converter-go/internal/cli"
	"github.com/docopt/docopt-go"

	v2 "github.com/asyncapi/converter-go/pkg/converter/v2"

	"fmt"
	"log"
	"os"
)

const version = "asyncapi-converter 0.2"

func main() {
	usage := fmt.Sprintf(`
  Convert AsyncAPI documents from version 1.x to %s. 

  Usage:
    asyncapi-converter <PATH> [--toYAML] [--id=<id>]
    asyncapi-converter -h | --help | --version

  Arguments:
    PATH        a path to asyncapi document (either url or local file, supports json and yaml format)  	

  Options:
    --toYAML    produces results in yaml format instead json
    --id=<id>   allows to specify application id`, v2.AsyncapiVersion)

	opts, err := docopt.ParseArgs(usage, nil, version)
	if err != nil {
		log.Fatal(err)
	}
	asyncapiCli := cli.New(opts)
	converter, reader, err := asyncapiCli.NewConverterAndReader()
	if err != nil {
		log.Fatal(err)
	}
	err = converter.Convert(reader, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
