package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Makepad-fr/wharf/core"
	"github.com/Makepad-fr/wharf/docker-plugins/commons"
)

var metadata = commons.PluginMetadata{
	SchemaVersion:    "0.1.0",
	Vendor:           "MAKEPAD",
	Version:          core.Version,
	ShortDescription: "Render Dockerfiles from templates",
	URL:              "https://github.com/Makepad-fr/wharf",
	Experimental:     true,
}

func main() {
	if commons.ShowPluginMetaData(metadata) {
		return
	}
	err := core.Run()
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}
}
