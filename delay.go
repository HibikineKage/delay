package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)
func viewHandler( w http.ResponseWriter, r *http.Request ) {
	w.Header().Set( "Content-Type","text/plain; charset=utf-8")
	fmt.Fprintln(w,"-------- Standard Library --------")
	fmt.Fprintln(w,"Method:", r.Method)
	fmt.Fprintln(w,"URL:",r.URL.String())
	fmt.Fprintln(w,"URL.Path:",r.URL.Path)
	n, ok := r.URL.Query()["name"]
	if ok {
		fmt.Fprintln(w,"hello:", n)
	}
}
func main() {
	http.HandleFunc("/gocgi/",viewHandler)
	cgi.Serve(nil)
}