package helper

import (
	"os"
	"strings"
)

// IsPathExists method for check path exists
func IsPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// IsFIleExists method for check file exists
func IsFileExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

// Mkdir method for make directory
func Mkdir(path string, perm os.FileMode) error {
	if ok := IsPathExists(path); !ok {
		if err := os.Mkdir(path, perm); err != nil {
			return err
		}
	}

	return nil
}

// GetWorkDir method
func GetWorkDir() string {
	rootPath := os.Getenv("PROJECT_DIR")
	sWorkDir := strings.Split(rootPath, "/")
	workDir := sWorkDir[len(sWorkDir)-1]

	return workDir
}

// CapitalizeWords method
func CapitalizeWords(s string) string {
	words := strings.Split(s, "-")
	var capitalizedWords []string

	for _, word := range words {
		if word != "" {
			capitalizedWord := strings.ToUpper(word[:1]) + word[1:]
			capitalizedWords = append(capitalizedWords, capitalizedWord)
		}
	}

	return strings.Join(capitalizedWords, "")
}
