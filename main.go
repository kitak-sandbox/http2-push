package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var script []byte

// preparing script
func init() {
	var err error
	script, err = ioutil.ReadFile("./main.js")
	if err != nil {
		panic(err)
	}
}

// Send HTML and push script
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		fmt.Println("Push /main.js")
		pusher.Push("/main.js", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><script src="/main.js"></script></body></html>`)
}

// Send script as usual HTTP request
func handlerJavaScript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(script)
}
func main() {
	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/main.js", handlerJavaScript)
	fmt.Println("start http listening :8000")
	err := http.ListenAndServeTLS(":8000", "server.crt", "server.key", nil)
	fmt.Println(err)
}
