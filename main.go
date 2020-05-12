package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	u := q.Get("username")
	p := q.Get("password")

	fmt.Fprintf(w, "hello %s, %s\n", u, p)
}

func getData(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/login", login)
	http.HandleFunc("/getdata", getData)

	http.ListenAndServe(":8090", nil)
}
