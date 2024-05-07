package docker_render

import (
	"encoding/json"
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

func ShowPluginMetaData() bool {
	if len(os.Args) > 1 && os.Args[1] == "docker-cli-plugin-metadata" {
		metadata := PluginMetadata{
			SchemaVersion:    "0.1.0",
			Vendor:           "MAKEPAD",
			Version:          version,
			ShortDescription: "A Helmlike templating engine for Docker",
			URL:              "https://github.com/Makepad-fr/docker-render",
			Experimental:     true,
		}
		json.NewEncoder(os.Stdout).Encode(metadata)
		return true
	}
	return false
}
