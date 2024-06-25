## Using Wharf as Go dependency
[![Go Reference](https://pkg.go.dev/badge/github.com/Makepad-fr/wharf/core.svg)](https://pkg.go.dev/github.com/Makepad-fr/wharf/core)

Wharf let you create Dockerfiles from Dockerfile templates programmatically with Go.

### Install the dependency

```bash
go get -u github.com/Makepad-fr/wharf/core@latest
```

## Example use cases 

### Render the a Dockerfile template to a string

```go
var stringBuilder strings.Builder
err := Render("../example/", "Dockerfile.template", "docker-values.yaml", &stringBuilder)
if err != nil {
    t.Error(err)
}
```
### Render the Dockerfile to a file

```go
file, err := os.CreateTemp(os.TempDir(), "Dockerfile")
	if err != nil {
		t.Error(err)
	}
	defer ile.Close()
	err = Render("../example", "Dockerfile.template", "docker-values.yaml", file)
```