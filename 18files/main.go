package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")
	content := "This needs to go in file - dev@go.ini"

	file ,err := os.Create("./testgo.txt")
	CheckNilErr((err))
	// if err != nil {
	// 	panic(err)
	// }

	msg,err := io.WriteString(file,content)
	CheckNilErr((err))
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("The lenght of msg is : ",msg)

	defer file.Close()

	readFile("./testgo.txt")
}


func readFile(filename string){
	databyte ,err := ioutil.ReadFile(filename)
	CheckNilErr((err))
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Text Data inside file is  : \n",string(databyte))

}

func CheckNilErr (err error){
	if err != nil {
		panic(err)
	}
}