module github.com/Makepad-fr/wharf/docker-cli-plugins/render

go 1.22.2

replace github.com/Makepad-fr/wharf/core => ../../core

replace github.com/Makepad-fr/wharf/docker-cli-plugins/commons => ../commons

require (
	github.com/Makepad-fr/wharf/core v0.0.0-00010101000000-000000000000
	github.com/Makepad-fr/wharf/docker-cli-plugins/commons v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v3 v3.0.1 // indirect
