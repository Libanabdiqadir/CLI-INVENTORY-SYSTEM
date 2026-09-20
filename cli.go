package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	system   *InventorySystem
	filename string
}

func NewCLI(system *InventorySystem, filename string) *CLI {
	return &CLI{
		system:   system,
		filename: filename,
	}
}

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	if err := LoadInventory(cli.system, cli.filename); err != nil {
		fmt.Printf("Error loading inventory: %v\n", err)
	} else {
		fmt.Println("Inventory loaded successfully.")
	}

	for {
		fmt.Println("\n=== Inventory Management System ===")
		fmt.Println("1. Add Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Add Customer")
		fmt.Println("4. View Customers")
		fmt.Println("5. Create Order")
		fmt.Println("6. View Orders")
		fmt.Println("7. Save & Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			cli.handleAddProduct(reader)
		case "2":
			cli.handleViewProducts()
		case "3":
			cli.handleAddCustomer(reader)
		case "4":
			cli.handleViewCustomers()
		case "5":
			cli.handleCreateOrder(reader)
		case "6":
			cli.handleViewOrders()
		case "7":
			if err := SaveInventory(cli.system, cli.filename); err != nil {
				fmt.Printf("Error saving inventory: %v\n", err)
			} else {
				fmt.Println("Inventory saved successfully. Goodbye!")
			}
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func (cli *CLI) handleAddProduct(reader *bufio.Reader) {
	fmt.Print("Enter product name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter product price: ")
	priceStr, _ := reader.ReadString('\n')
	price, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil {
		fmt.Println("Invalid price.")
		return
	}

	fmt.Print("Enter stock quantity: ")
	stockStr, _ := reader.ReadString('\n')
	stock, err := strconv.Atoi(strings.TrimSpace(stockStr))
	if err != nil {
		fmt.Println("Invalid stock quantity.")
		return
	}

	p := cli.system.ProductManager.AddProduct(0, name, "", price, int64(stock), time.Now())
	fmt.Printf("Product added successfully! ID: %d\n", p.ID)
}

func (cli *CLI) handleViewProducts() {
	fmt.Println("\n--- Products List ---")
	if len(cli.system.ProductManager.Products) == 0 {
		fmt.Println("No products found.")
		return
	}
	for _, p := range cli.system.ProductManager.Products {
		fmt.Printf("ID: %d | Name: %s | Price: $%.2f | Stock: %d\n", p.ID, p.Name, p.Price, p.Quantity)
	}
}

func (cli *CLI) handleAddCustomer(reader *bufio.Reader) {
	fmt.Print("Enter customer name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter customer email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter customer phone number: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	c := cli.system.CustomerManager.AddCustomer(name, email, phone)
	fmt.Printf("Customer added successfully! ID: %d\n", c.ID)
}

func (cli *CLI) handleViewCustomers() {
	fmt.Println("\n--- Customers List ---")
	if len(cli.system.CustomerManager.Customers) == 0 {
		fmt.Println("No customers found.")
		return
	}
	for _, c := range cli.system.CustomerManager.Customers {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s\n", c.ID, c.Name, c.Email, c.PhoneNumber)
	}
}

func (cli *CLI) handleCreateOrder(reader *bufio.Reader) {
	fmt.Print("Enter Customer ID: ")
	custIDStr, _ := reader.ReadString('\n')
	custID, err := strconv.Atoi(strings.TrimSpace(custIDStr))
	if err != nil {
		fmt.Println("Invalid Customer ID.")
		return
	}

	var items []OrderItem
	for {
		fmt.Print("Enter Product ID to order (or type 'done' to finish): ")
		prodInput, _ := reader.ReadString('\n')
		prodInput = strings.TrimSpace(prodInput)
		if strings.ToLower(prodInput) == "done" {
			break
		}

		prodID, err := strconv.Atoi(prodInput)
		if err != nil {
			fmt.Println("Invalid Product ID.")
			continue
		}

		var foundPrice float64
		found := false
		for _, p := range cli.system.ProductManager.Products {
			if p.ID == prodID {
				foundPrice = p.Price
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Product ID not found. Try again.")
			continue
		}

		fmt.Print("Enter quantity: ")
		qtyStr, _ := reader.ReadString('\n')
		qty, err := strconv.Atoi(strings.TrimSpace(qtyStr))
		if err != nil || qty <= 0 {
			fmt.Println("Invalid quantity.")
			continue
		}

		items = append(items, OrderItem{
			ProductID: prodID,
			Quantity:  qty,
			PriceEach: foundPrice,
		})
	}

	if len(items) == 0 {
		fmt.Println("No items added. Order creation cancelled.")
		return
	}

	order := cli.system.OrderManager.CreateOrder(custID, items, "Completed")
	fmt.Printf("Order created successfully! Order ID: %d | Total Amount: $%.2f\n", order.ID, order.TotalAmount)
}

func (cli *CLI) handleViewOrders() {
	fmt.Println("\n--- Orders History ---")
	if len(cli.system.OrderManager.Orders) == 0 {
		fmt.Println("No orders found.")
		return
	}
	for _, o := range cli.system.OrderManager.Orders {
		fmt.Printf("Order ID: %d | Customer ID: %d | Status: %s | Total: $%.2f | Date: %s\n",
			o.ID, o.CustomerID, o.Status, o.TotalAmount, o.OrderDate.Format("2006-01-02 15:04"))
		for _, item := range o.Items {
			fmt.Printf("   -> Product ID: %d | Qty: %d | Price Each: $%.2f\n", item.ProductID, item.Quantity, item.PriceEach)
		}
	}
}