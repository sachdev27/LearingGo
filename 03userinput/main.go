package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our Pizza : ")

	// Comma Ok || err Ok

	input, err := reader.ReadString('\n')

	fmt.Println("You entered ", input)
	fmt.Println("Error Shown ", err)
	fmt.Printf("type of rating %T", input)
}