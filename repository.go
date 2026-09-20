package main

import (
	"encoding/json"
	"os"
)

func SaveInventory(sys *InventorySystem, filename string) error {
	byteData, err := json.MarshalIndent(sys, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, byteData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadInventory(sys *InventorySystem, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, sys)
	if err != nil {
		return err
	}

	return nil
}