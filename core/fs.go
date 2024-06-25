package wharf

import "os"

// IsDirectory checks if the given path is a directory
func isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
