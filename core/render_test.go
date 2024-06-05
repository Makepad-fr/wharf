package wharf

import (
	"io"
	"os"
	"strings"
	"testing"
)

const expectedRenderedDockerfile = `
# Use an official base image
FROM ubuntu:latest

# Set the exposed port
EXPOSE 9091

# Set environment variables

ENV DEBUG="false"

ENV LOG_LEVEL="info"


# Copy source code
COPY . /app

# Run command
CMD echo Hello World
`

func TestRenderToString(t *testing.T) {
	var stringBuilder strings.Builder
	err := Render("../example/", "Dockerfile.template", "docker-values.yaml", &stringBuilder)
	if err != nil {
		t.Error(err)
	}

	generated := strings.TrimSpace(stringBuilder.String())
	expected := strings.TrimSpace(expectedRenderedDockerfile)

	if generated != expected {
		t.Log("Strings are not equal")
		t.Logf("Generated length: %d", len(generated))
		t.Logf("Expected length: %d", len(expected))

		for i := 0; i < len(generated) && i < len(expected); i++ {
			if generated[i] != expected[i] {
				t.Errorf("Difference at char %d: '%c' != '%c'", i, generated[i], expected[i])
				break
			}
		}

		if len(generated) != len(expected) {
			t.Error("Generated and expected strings have different lengths")
		}
	}
}

func TestRenderToOutputFile(t *testing.T) {
	// Render the template to a file
	// Compare result with the expected one
	file, err := os.CreateTemp(os.TempDir(), "Dockerfile")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
	err = Render("../example", "Dockerfile.template", "docker-values.yaml", file)
	if err != nil {
		t.Error(err)
	}
	equals, err := compareFiles(file, "../example/Dockerfile.expected")
	if err != nil {
		t.Error(err)
	}
	if !equals {
		t.Error("Files are not equals")
	}
}

// CompareFiles checks if the contents of two files are the same.
func compareFiles(file1 *os.File, file2 string) (bool, error) {
	file1.Seek(0, io.SeekStart)
	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	const bufferSize = 4096
	buf1 := make([]byte, bufferSize)
	buf2 := make([]byte, bufferSize)

	for {
		n1, err1 := file1.Read(buf1)
		n2, err2 := f2.Read(buf2)

		if n1 != n2 || err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil // End of both files reached
			}
			return false, nil
		}

		if n1 == 0 && n2 == 0 {
			break // End of both files reached
		}

		if !compareChunks(buf1[:n1], buf2[:n2]) {
			return false, nil
		}
	}

	return true, nil
}

// compareChunks checks if two byte slices are equal.
func compareChunks(chunk1, chunk2 []byte) bool {
	if len(chunk1) != len(chunk2) {
		return false
	}
	for i := range chunk1 {
		if chunk1[i] != chunk2[i] {
			return false
		}
	}
	return true
}
