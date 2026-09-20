package main

import (
	"time"
)

type OrderItem struct {
	ProductID 		int
	Quantity 			int
	PriceEach			float64
}

type Order struct {
	ID         		int
	CustomerID		int
	Items      		[]OrderItem
	OrderDate  		time.Time
	Status     		string
	TotalAmount 	float64  
}

type OrderManager struct {
	Orders []Order
}

func (o *OrderManager) CreateOrder(customerId int, items []OrderItem, status string) Order {
	newID := len(o.Orders) + 1

	var total float64
	for _, item := range items {
		total += item.PriceEach * float64(item.Quantity)
	}

	newOrder := Order{
		ID: newID,
		CustomerID: customerId,
		OrderDate: time.Now(),
		Status: status,
		TotalAmount: total,
	}

	o.Orders = append(o.Orders, newOrder)
	return newOrder 
}