package main

import "fmt"

func main() {
	fmt.Println("Welcome to Class on Pointers...")

	// var ptr *int
	// fmt.Println("Value of pointer is :",ptr)

	myNumber := 54
	var ptr = &myNumber
	fmt.Println("Value of Pointer is :",ptr)
	fmt.Println("Pointer is pointing to ( Whats inside the memory address that it is pointing to) :",*ptr)

	*ptr = *ptr + 2
	fmt.Println("After incrementing the value of pointer by 2, the value of pointer is :",ptr)
	fmt.Println("After incrementing the value of pointer by 2, the value inside pointer is :",*ptr)

}
