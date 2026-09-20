package main

func main() {

	system := NewInventorySystem()

	cli := NewCLI(system, "inventory.json")

	cli.Run()
}