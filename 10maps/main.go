package main

import "fmt"

func main() {
	fmt.Println("Welcome, Today we will be learning about maps")

	var languagues = make(map[string]string)

	languagues["JS"] = "Javascript"
	languagues["GO"] = "Golang"
	languagues["PY"] = "Python"

	fmt.Println("JS shorts for :",languagues["JS"])

	fmt.Println(languagues)

	value, isPresent := languagues["JS"]
	println(value,"->",isPresent)

	delete(languagues,"GO")
	fmt.Println(languagues)



}
