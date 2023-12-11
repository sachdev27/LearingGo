package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greeter()
	// greeter1()

	// result := adder(3,5)
	result,message := proAdder(3,5,1,2,4,6,6,7,78,8)
	fmt.Println("Addtion of all the values is :",result)
	fmt.Println("ProMessage is : ",message)
}

func adder(a int, b int) int{
	return a+b
}

func proAdder(values ...int) (int,string) {
	total := 0
	for index := range values {
		total += values[index]
	}
	return total,"Added successfully"
}

func greeter(){
	fmt.Println("Welcome to Earth!")
}

func greeter1(){
	fmt.Println("Welcome to Earth Again!")
}