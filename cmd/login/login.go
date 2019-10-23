package main

import (
	"fmt"
	. "github.com/MccGithub/GoWeb/base/view_page"
	"log"
	"net/http"
	"strings"
)

//var basePath = "/home/m/go/src/github.com/MccGithub/GoWeb/template"
//var basePath = "./template" // 相对路径中的当前路径(.)会随着程序在哪个路径运行而改变

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm() // 解析 url 传递的参数, 对于POST则解析响应包的主题 (request body)
	// 注意: 如果没有调用ParseForm方法, 下面无法获取表单的数据
	fmt.Println(r.Form)	// 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie!")	// 这个写入到 w 的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)	// 获取请求的方法
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("github.com/MccGithub/GoWeb/template/login/login.html")
		if err := View("login/login.html", w, nil); err != nil {
			log.Println(err)
		}
	} else {
		err := r.ParseForm()	// 解析 url 传递的参数, 对于 POST 则解析响应包的主题 (request body)
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据, 那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	SetTempRelativePath("../../template")
	http.HandleFunc("/", sayhelloName)	// 设置访问的路由
	http.HandleFunc("/login", login)	// 设置访问的路由
	err := http.ListenAndServe(":9090", nil)	// 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

