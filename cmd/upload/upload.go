package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var basePath string

func getCurrent() {
	_, basePath, _, _ = runtime.Caller(0)
	i := strings.LastIndex(basePath, string(os.PathSeparator))
	basePath = basePath[0:i+1]
}

func getTempPath(file string) string {
	//return basePath+"../../template/"+file
	return filepath.Join(basePath, "../../template", file)
}

func View(file string, w io.Writer, data interface{}) error {
	temp, err := template.ParseFiles(getTempPath(file))
	if err != nil {
		return err
	}

	return temp.Execute(w, data)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)	// 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		if err := View("upload/upload.html", w, token); err != nil {
			log.Println(err)
		}
	} else {
		_ = r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(basePath+"test/"+handler.Filename,
			os.O_WRONLY|os.O_CREATE, 0666)	// 此处假设当前目录下已存在test目录
			if err != nil {
				fmt.Println(err)
				return
			}
		defer f.Close()
		_, _ = io.Copy(f, file)
	}
}

func main() {
	getCurrent()
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

