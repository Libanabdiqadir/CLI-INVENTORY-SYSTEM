package main

import (
	"time"
)

type Product struct {
	ID int
	Name string
	Category string
	Price float64
	Quantity int64
	CreatedAt time.Time
}

type ProductManager struct {
	Products []Product
}

func (m *ProductManager) AddProduct(
		id int, name string, category string, 
		price float64, quantity int64, 
		createdAt time.Time) Product {
	
	newID := len(m.Products) + 1
	newProduct := Product{
		ID: newID,
		Name: name,
		Category: category,
		Price: price,
		Quantity: quantity,
		CreatedAt: createdAt,
	}

	m.Products = append(m.Products, newProduct)
	return newProduct
}