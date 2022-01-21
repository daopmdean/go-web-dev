package main

import (
	"fmt"
	"io"
	"net/http"
)

type simp int

func (s simp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "the dog")
	case "/cat":
		io.WriteString(res, "The cat")
	}
}

func main() {
	var s simp
	fmt.Println("Serve on port 1234")
	http.ListenAndServe(":1234", s)
}
