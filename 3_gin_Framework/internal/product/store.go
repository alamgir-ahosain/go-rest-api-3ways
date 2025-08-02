package product

import "fmt"

var products = []Product{
	{ID: 1, Name: "Apple", Price: 10},
	{ID: 2, Name: "Mango", Price: 20},
	{ID: 3, Name: "Jack Fruit", Price: 100},
}

// List return all product
func List() []Product {
	return products
}

// add product
func Add(p Product) {
	products = append(products, p)

}

//get product
func GetProductByID(id int) (*Product, error) {
	for idx, val := range products {
		if val.ID == id {
			return &products[idx], nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

//update product
func UpdateProduct(p *Product) {
	p.Price += 2
}

//Delete product
func DeleteProductByID(id int) error {
	for idx, val := range products {
		if val.ID == id {
			//remove product at index idx 
			products=append(products[:idx],products[idx+1:]... )
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}
