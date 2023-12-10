package main

import "fmt"

func main() {
	fmt.Println("ifelse in Go")
	var loginCount int

	loginCount = 10

	if loginCount < 10 {
		fmt.Println("Regular User")
	} else if loginCount > 10 {
		fmt.Println("Something Else")
	} else {
		fmt.Println("idk")
	}


	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Weird Syntax
	if num := 3; num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than 10")
	}

	var err error;

	if  err != nil {
		fmt.Println("Err is not nil")
	} else {
		fmt.Println("err is nil")
	}

}
