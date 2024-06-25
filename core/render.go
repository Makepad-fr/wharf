package wharf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

var funcMap = map[string]any{
	"include": Include,
}

func Include(path string, values map[string]any) string {
	isDir, err := isDirectory(path)

	if err != nil {
		return fmt.Sprintf("Error while importing %s: %s", path, err.Error())
	}
	// Create a generic error message template
	const errorMessage = "Error: Include should be called either with a Dockerfile.template path or with a Dockerfile path"
	if isDir {
		return errorMessage
	} else {
		baseFileName := filepath.Base(path)
		if baseFileName == "Dockerfile" {
			// If the path is a Dockerfile path
			file, err := os.Open(path)
			if err != nil {
				return fmt.Sprintf("Error: error while opening Dockerfile from %s", path)
			}
			content, err := io.ReadAll(file)
			if err != nil {
				return fmt.Sprintf("Error: error while reading file content %s", err.Error())
			}
			return string(content)
		}
		if baseFileName == "Dockerfile.template" {
			// IF the path is a Dockerfile.template
			contextPath := filepath.Dir(path)
			template, err := getTemplate(contextPath, baseFileName)
			if err != nil {
				return fmt.Sprintf("Error: error while getting template from %s: %s", path, err.Error())
			}
			var renderedContent bytes.Buffer
			// Don't use include in included templates as it creates a cyclic dependency
			// TODO: Maybe we can copy the funcMap before using it?
			err = render(template, values, &renderedContent, map[string]any{})
			if err != nil {
				return fmt.Sprintf("Error: error while rendering template %s: %s", path, err.Error())
			}
			return renderedContent.String()
		}
		return errorMessage
	}
}

// Render renders the template file in the given contextPath using the values files from the given path
// It writes the rendered Dockerfile to the io.Writer passed in parameters. It returns an error if something goes wrong
func Render(contextPath, templateFileName, valuesFilePath string, output io.Writer) error {
	template, err := getTemplate(contextPath, templateFileName)
	if err != nil {
		return err
	}
	values, err := readValues(contextPath, valuesFilePath)
	if err != nil {
		return err
	}

	err = render(template, values, output, funcMap)
	return err
}

// render renders the template and values to the file with the given path.
// If something goes wrong it returns an error
func render(tmpl *template.Template, values map[string]any, file io.Writer, funcMap template.FuncMap) error {
	// Include function map to the current template just before rendering the values
	tmpl = tmpl.Funcs(funcMap)
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
