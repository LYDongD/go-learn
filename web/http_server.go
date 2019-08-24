package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("value", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["password"])
	}

}

func main() {

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9081", nil)
	if err != nil {
		log.Fatal("listen and serve: ", err)
	}
}
