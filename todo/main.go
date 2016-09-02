package main

import (
	"net/http"
	"github.com/charlires/todo/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)
}
