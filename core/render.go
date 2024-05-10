package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

var defaultTemplateFileName = "Dockerfile.template"

func init() {
	v, ok := os.LookupEnv("DOCKER_RENDER_DEFAULT_TEMPLATE_FILE_NAME")
	if ok {
		defaultTemplateFileName = v
	}
}

// const defaultContextPath =

func Run() error {
	os.Args = os.Args[1:]
	valuesFilePath := flag.String("values", "./docker-values.yaml", "The path for the values file to use")
	outputFilePath := flag.String("output", "", "The path for the output file")
	templateFileName := flag.String("file-name", defaultTemplateFileName, "The name of the template file")
	flag.Parse()
	if flag.NArg() != 1 {
		return errors.New("the path for the template file is required\n")
	}
	contextPath := flag.Arg(0)
	template, err := getTemplate(contextPath, *templateFileName)
	if err != nil {
		return err
	}
	values, err := readValues(contextPath, *valuesFilePath)
	if err != nil {
		return err
	}
	err = render(*outputFilePath, template, values)
	return nil
}

// render renders the template and values to the file with the given path.
// If something goes wrong it returns an error
func render(path string, tmpl *template.Template, values map[string]any) error {
	var file *os.File = os.Stdout
	var err error
	if len(strings.TrimSpace(path)) > 0 {
		file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	err = tmpl.Execute(file, values)
	if err != nil {
		return err
	}
	return nil
}

// readValues read the values yaml file from the given context path. If the provided values file path is an absolute path
// the context path will be ignored.
func readValues(contextPath, valuesFilePath string) (map[string]any, error) {
	if !filepath.IsAbs(valuesFilePath) {
		// Join the given valuesFilePath with the contextPath
		valuesFilePath = filepath.Join(contextPath, valuesFilePath)
	}
	f, err := os.ReadFile(valuesFilePath)
	if err != nil {
		return nil, err
	}
	var config map[string]any
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// getTemplate returns the template.Template from given context path and templatefile name
func getTemplate(contextPath, templateFileName string) (*template.Template, error) {
	info, err := os.Stat(contextPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("Context path %s should be a directory path. To pass template file name use -file-name option", contextPath)
	}

	contextPath = filepath.Join(contextPath, templateFileName)

	templ, err := template.ParseFiles(contextPath)
	if err != nil {
		return nil, err
	}

	return templ, nil
}
