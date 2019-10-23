package view_page

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var tempBasePath string
var curPath string
var SetTempRelativePath func(path string) = func (path string) {
	_, f, _, _ := runtime.Caller(0)
	i := strings.LastIndex(f, string(os.PathSeparator))
	curPath = f[0:i+1]

	tempBasePath = curPath + path
	fmt.Println(tempBasePath)

	if !filepath.IsAbs(tempBasePath) {
		tempBasePath, _ = filepath.Abs(tempBasePath)
	}
	fmt.Println(tempBasePath)
}

func getTempPath(file string) string {
	return filepath.Join(tempBasePath, file)
}

func GetCurrentPath() string {
	return curPath
}

func GetTempBasePath() string {
	return tempBasePath
}
