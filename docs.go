//go:build ignore

package main

import "net/http"

func main() {
	http.ListenAndServe(`127.0.0.1:8080`, http.FileServer(http.Dir(`./docs`)))
}
