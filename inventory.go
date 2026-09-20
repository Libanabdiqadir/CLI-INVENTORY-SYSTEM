package main

type InventorySystem struct {
	ProductManager
	CustomerManager
	OrderManager
}

func NewInventorySystem() *InventorySystem {
	return &InventorySystem{
		ProductManager:  ProductManager{},
		CustomerManager: CustomerManager{},
		OrderManager:    OrderManager{},
	}
}