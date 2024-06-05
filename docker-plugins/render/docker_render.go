package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	wharf "github.com/Makepad-fr/wharf/core"
	"github.com/Makepad-fr/wharf/docker-plugins/commons"
)

var metadata = commons.PluginMetadata{
	SchemaVersion:    "0.1.0",
	Vendor:           "MAKEPAD",
	Version:          wharf.Version,
	ShortDescription: "Render Dockerfiles from templates",
	URL:              "https://github.com/Makepad-fr/wharf",
	Experimental:     true,
}

var defaultTemplateFileName = "Dockerfile.template"

func init() {
	v, ok := os.LookupEnv("DOCKER_RENDER_DEFAULT_TEMPLATE_FILE_NAME")
	if ok {
		defaultTemplateFileName = v
	}
}

func run() error {
	os.Args = os.Args[1:]
	valuesFilePath := flag.String("values", "./docker-values.yaml", "The path for the values file to use")
	outputFilePath := flag.String("output", "", "The path for the output file")
	templateFileName := flag.String("file-name", defaultTemplateFileName, "The name of the template file")
	flag.Parse()
	if flag.NArg() != 1 {
		return errors.New("the path for the template file is required\n")
	}
	contextPath := flag.Arg(0)
	// If output filepath is an

	var file *os.File = os.Stdout

	var err error

	if outputFilePath != nil && (len(strings.TrimSpace(*outputFilePath)) > 0) {
		file, err = os.OpenFile(*outputFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	err = wharf.Render(contextPath, *templateFileName, *valuesFilePath, file)
	return err
}

func main() {
	if commons.ShowPluginMetaData(metadata) {
		return
	}
	err := run()
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}
}
