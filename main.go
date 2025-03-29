package main

import (
	"html/template"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	files, _ := template.ParseFiles("view/index.html")
	files.Execute(res, 123)
}

func login(res http.ResponseWriter, req *http.Request) {
	files, _ := template.ParseFiles("view/login.html")
	files.Execute(res, 123)
}

func userList(res http.ResponseWriter, req *http.Request) {
	files, _ := template.ParseFiles("view/userList.html")
	files.Execute(res, 123)
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/index.html", index)
	http.HandleFunc("/login.html", login)
	http.HandleFunc("/userList.html", userList)
	server.ListenAndServe()
}
