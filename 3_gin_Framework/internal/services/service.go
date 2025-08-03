package services

import (
	"fmt"

	"github.com/alamgir-ahosain/go-rest-api-3ways/3_gin_Framework/internal/models"
)

var products = []models.Product{
	{ID: 1, Name: "Apple", Price: 10},
	{ID: 2, Name: "Mango", Price: 20},
	{ID: 3, Name: "Jack Fruit", Price: 100},
}

// List return all product
func List() []models.Product {
	return products
}

// add product
func Add(p models.Product) {
	products = append(products, p)

}

//get product
func GetProductByID(id int) (*models.Product, error) {
	for idx, val := range products {
		if val.ID == id {
			return &products[idx], nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

//update product
func UpdateProduct(p *models.Product) {
	p.Price += 2
}

//put product
func PutProduct(id int, p *models.Product) (*models.Product, bool) {
	for i, val := range products {
		if val.ID == id {
			p.ID = id
			products[i] = *p
			return &products[i], true

		}
	}
	return nil, false
}

//Delete product
func DeleteProductByID(id int) error {
	for idx, val := range products {
		if val.ID == id {
			//remove product at index idx
			products = append(products[:idx], products[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}
