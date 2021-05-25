package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Batam", "Kepri", "Indonesia"}
	address2 := &address1 //pointer refer by reference
	// address2 := address1 // refer by value

	address2.City = "Nongsa"

	fmt.Println(address1)
	fmt.Println(address2)
}
