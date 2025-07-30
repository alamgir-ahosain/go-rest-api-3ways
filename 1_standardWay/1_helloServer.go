package main

import (
	"net/http"
)

type helloPage struct{}
func (h *helloPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is hello page"))
}

func main() {

	mux := http.NewServeMux() //router
	mux.Handle("/hello", &helloPage{}) //register route
	http.ListenAndServe(":8080", mux) 

}
