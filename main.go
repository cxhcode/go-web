package main

import (
	"net/http"
	"log"
	"fmt"
	"text/template"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key -> ", key)
		fmt.Println("value -> ", value)
	}
	fmt.Fprint(w, "这是陈新华的第一个go web 程序")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		if len(strings.TrimSpace(r.Form.Get("username"))) == 0 {
			w.Write([]byte("用户名不能为空"))
		} else {
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			by := []byte("登录成功")
			w.Write(by)
		}
	}
}

func main() {

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
