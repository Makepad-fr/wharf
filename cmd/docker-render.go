package main

import (
	"flag"
	"fmt"
	"os"

	docker_render "github.com/Makepad-fr/docker-render"
)

func main() {
	if docker_render.ShowPluginMetaData() {
		return
	}
	err := docker_render.Run()
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}
}
