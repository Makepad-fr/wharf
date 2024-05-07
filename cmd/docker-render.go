package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PluginMetadata struct {
	SchemaVersion    string `json:"SchemaVersion"`
	Vendor           string `json:"Vendor"`
	Version          string `json:"Version"`
	ShortDescription string `json:"ShortDescription"`
	URL              string `json:"Url"`
	Experimental     bool   `json:"Experimental"`
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "docker-cli-plugin-metadata" {
		metadata := PluginMetadata{
			SchemaVersion:    "0.1.0",
			Vendor:           "MAKEPAD",
			Version:          "0.0.1",
			ShortDescription: "A Helmlike templating engine for Docker",
			URL:              "https://github.com/Makepad-fr/docker-template",
			Experimental:     true,
		}
		json.NewEncoder(os.Stdout).Encode(metadata)
		return
	}
	fmt.Println("Hello world")
}
