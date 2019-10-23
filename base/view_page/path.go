package view_page

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var basePath string

func getPath(file string) string {
	if !filepath.IsAbs(basePath) {
		basePath, _ = filepath.Abs(basePath)
	}

	return filepath.Join(basePath, file)
}

func GetCurrent() {
	_, f, _, _ := runtime.Caller(0)
	i := strings.LastIndex(f, string(os.PathSeparator))
	basePath = f[0:i+1] + "../../template"
}
