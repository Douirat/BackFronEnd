package main

import (
	products "products/products"
)

func main() {
	// Declare a global variable of type Products
	products.AllProducts.AppendProduct("123", "T590", "ThinkPad", "Lenovo")
	products.AllProducts.AppendProduct("456", "4050", "latitude", "Dell")
	products.AllProducts.AppendProduct("789", "HQ45", "workbook", "HP")
	products.AllProducts.AppendProduct("987", "480", "YOGA", "Lenovo")
	products.AllProducts.AppendProduct("654", "SAM32", "GALAXY", "SAMSUNG")
	products.AllProducts.AppendProduct("321", "X", "Thinkcenter", "Lenovo")
	products.AllProducts.Display()
	products.Routing()
}
