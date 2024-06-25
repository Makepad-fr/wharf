package wharf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Render renders the template file in the given contextPath using the values files from the given path
// It writes the rendered Dockerfile to the io.Writer passed in parameters. It returns an error if something goes wrong
func Render(contextPath, templateFileName, valuesFilePath string, output io.Writer) error {
	// Create a new function map
	funcMap := newFunctionFactory(contextPath).FuncMap()
	template, err := getTemplate(contextPath, templateFileName, funcMap)
	if err != nil {
		return err
	}
	values, err := readValues(contextPath, valuesFilePath)
	if err != nil {
		return err
	}

	err = render(template, values, output)
	return err
}

// render renders the template and values to the file with the given path.
// If something goes wrong it returns an error
func render(tmpl *template.Template, values map[string]any, file io.Writer) error {
	// Include function map to the current template just before rendering the values
	err := tmpl.Execute(file, values)
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
func getTemplate(contextPath, templateFileName string, funcMap template.FuncMap) (*template.Template, error) {
	info, err := os.Stat(contextPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("context path %s should be a directory path. To pass template file name use -file-name option", contextPath)
	}

	contextPath = filepath.Join(contextPath, templateFileName)

	templ, err := template.New(filepath.Base(templateFileName)).Funcs(funcMap).ParseFiles(contextPath)
	if err != nil {
		return nil, err
	}
	return templ, nil
}
