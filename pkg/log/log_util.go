package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func _formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

// try create folder to file
func createFile(output string) (string, error) {
	workingDir, _ := os.Getwd()
	// fmt.Println(workingDir)
	// fmt.Println(filepath.Dir(output))
	// fmt.Println(filepath.Join(".", output))
	// fmt.Println(filepath.Join(workingDir, output))

	absPath := filepath.Join(workingDir, output)
	dirPath := filepath.Dir(absPath)

	// detect exists
	_, err := os.Stat(dirPath)

	// create new if not
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return absPath, nil
}

// split method called to only file name instead of full path
func callerOnlyFileName(frame *runtime.Frame) (function string, file string) {
	return "", fmt.Sprintf("%s:%d", _formatFilePath(frame.File), frame.Line)
}
