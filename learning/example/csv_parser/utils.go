package main

import (
	"io/ioutil"
	"net"
	"path"
	"strings"
)

func getFilesFromFolder(dirPath string) (lst []string, err error) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}

		lst = append(lst, path.Join(dirPath, file.Name()))
	}

	return
}

func isIPv4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	if ip.To4() == nil {
		return false
	}
	return true
}
