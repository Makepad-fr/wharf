package types

import "text/template"

// TODO: template.FuncMap uses map[string]any and any is a function with an arbitrary number of arguments func(args ....), we need to check if we can create a type for that
type FuncMap = template.FuncMap

// TODO: Change to more restrictive type as it will be used only for values
type ValueMap = map[string]any
