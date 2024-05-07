package commons

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

func ShowPluginMetaData(metadata PluginMetadata) bool {
	if len(os.Args) > 1 && os.Args[1] == "docker-cli-plugin-metadata" {
		json.NewEncoder(os.Stdout).Encode(metadata)
		return true
	}
	return false
}
