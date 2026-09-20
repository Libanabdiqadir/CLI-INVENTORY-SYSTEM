package main

import (
	"time"
)

type Customer struct {
	ID          int
	Name        string
	Email       string
  PhoneNumber string
	CreatedAt time.Time
}

type CustomerManager struct {
	Customers []Customer
}

func (cm *CustomerManager) AddCustomer(name, email, phone string) Customer {
    newID := len(cm.Customers) + 1
    newCustomer := Customer{
        ID:          newID,
        Name:        name,
        Email:       email,
        PhoneNumber: phone,
    }
    cm.Customers = append(cm.Customers, newCustomer)
    return newCustomer
}