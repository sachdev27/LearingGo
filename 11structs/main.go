package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	sandesh := User{"Sandesh", "sandeshsachdev27@gmail.com" , true, 23}
	fmt.Printf("Sandesh details are : %+v\n",sandesh)
	fmt.Printf("Name is %v ,Email is %v\n",sandesh.Name,sandesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
