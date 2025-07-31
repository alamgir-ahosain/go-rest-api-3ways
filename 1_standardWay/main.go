package main

import (
	"net/http"
	"regexp"
	
)

type homePageHandler struct{}
func (h *homePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Home page"))
}

type ProductsHandler struct{}
// func (b *bookStoresHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This is Book Store page"))
// }

func (h *ProductsHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {}
func (h *ProductsHandler) ListProducts(w http.ResponseWriter, r *http.Request) {}
func (h *ProductsHandler) GetProduct(w http.ResponseWriter, r *http.Request) {}
func (h *ProductsHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {}
func (h *ProductsHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {}

var(
	productRe=regexp.MustCompile(`^/bookStores/*$`) 
    productReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`) // /bookStores vs bookStores/<id>

)
//dispatch the request to correct handler
func(h *ProductsHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch{
	case r.Method==http.MethodPost && productRe.MatchString(r.URL.Path):
		h.CreateProduct(w,r)
		return
	case r.Method==http.MethodGet && productRe.MatchString(r.URL.Path):
		h.ListProducts(w,r)
		return
	case r.Method==http.MethodGet && productReWithID.MatchString(r.URL.Path) :
		h.GetProduct(w,r)
		return
	case r.Method==http.MethodPut && productReWithID.MatchString(r.URL.Path):
		h.UpdateProduct(w,r)
		return
	case r.Method==http.MethodDelete && productReWithID.MatchString(r.URL.Path):
		h.DeleteProduct(w,r)
		return
	default:
		return
	}
	
}

func main() {

	mux := http.NewServeMux() //router

	//register the route and handler
	mux.Handle("/", &homePageHandler{}) 
	mux.Handle("/bookStores", &ProductsHandler{}) 
	mux.Handle("/bookStores/", &ProductsHandler{}) 
	
	//Run the server
	http.ListenAndServe(":8080", mux) 

}
