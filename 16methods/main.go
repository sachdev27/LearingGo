package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	sandesh := User{"Sandesh", "sandeshsachdev27@gmail.com" , true, 23}
	fmt.Printf("Sandesh details are : %+v\n",sandesh)
	fmt.Printf("Name is %v ,Email is %v\n",sandesh.Name,sandesh.Email)
	sandesh.getStatus()
	sandesh.NewMail()
	fmt.Println("Email : ",sandesh.Email) // didn't change the mail
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus(){
	fmt.Println("Is User Actuve : ",u.Status)
}

func (u User) NewMail(){
	fmt.Println("Changing the mail!")
	u.Email = "test.broadcom.com"
	fmt.Println("New Email is : ",u.Email)
}