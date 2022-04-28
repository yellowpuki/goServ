package router

import "net/http"

type rootHandler struct{}

func (r rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
