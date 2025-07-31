package products


//represent a product
type Product struct {
	Name  string  `json:"name"`
	Price []Price `json:"price"`
}

//represent individual price
type Price struct {
	Name int `json:"price"`
}
