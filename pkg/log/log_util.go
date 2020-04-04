package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func _formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func createFile(output string) (string, error) {
	// TODO:
	// if os.IsNotExist(output) {
	//
	// }
	wd, _ := os.Getwd()
	fmt.Println(wd)

	return output, nil
}

// split method called to only file name instead of full path
func callerOnlyFileName(frame *runtime.Frame) (function string, file string) {
	return "", fmt.Sprintf("%s:%d", _formatFilePath(frame.File), frame.Line)
}
