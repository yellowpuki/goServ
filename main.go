package main

import (
	"/internal/router"
	"net/http"
)

func main() {
	r := router.New()
	http.ListenAndServe(":8080", r.RootHandler())
}
