package wharf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	types "github.com/Makepad-fr/wharf/types"
)

// newFunctionFactory creates a new function factory witht the given context path
func newFunctionFactory(contextPath string) functionFactory {
	return functionFactory{
		contextPath: contextPath,
	}
}

// functionFactory crates a context aware functionMap for each template.
type functionFactory struct {
	contextPath string // THe path of the current template file that executes this function map
}

// Include includdes returns the content of as string. The [path] should either be
// a Dockerfile or a Dockerfile.template. If it's a Dockefile.template the [values] will
// be used to render the template using wharf. If it's a Dockerfile the cotnent will be rendered as it's.
// If any erorr happens it will returned as an error string as template function can not return errors as Go objects.
// FIXME: Users may use other filenames other then Dockerfile and Dockerifle.template especially when they are using from the same folder
func (f functionFactory) Include(path string, values types.ValueMap) string {
	// Update the passed path with the context path
	path = filepath.Join(f.contextPath, path)
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
			// Don't use include in included templates as it creates a cyclic dependency
			// TODO: Maybe we can copy the funcMap before using it?
			template, err := getTemplate(contextPath, baseFileName, map[string]any{})
			if err != nil {
				return fmt.Sprintf("Error: error while getting template from %s: %s", path, err.Error())
			}
			var renderedContent bytes.Buffer
			err = render(template, values, &renderedContent)
			if err != nil {
				return fmt.Sprintf("Error: error while rendering template %s: %s", path, err.Error())
			}
			return renderedContent.String()
		}
		return errorMessage
	}
}

// Creates the types.FuncMap for the related function factory
// FIXME: Check if it's possible to create this function map programmatically
func (f functionFactory) FuncMap() types.FuncMap {
	return map[string]any{
		"include": f.Include,
	}
}
