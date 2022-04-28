package main

import (
	"github.com/stretchr/testify/tree/master/require/internal/router"
	"net/http"
)

func main() {
	r := router.New()
	http.ListenAndServe(":8080", r.RootHandler())
}
