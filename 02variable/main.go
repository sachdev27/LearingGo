package main

import "fmt"

const LoginToken string = "IFBEWIUNA" // Public


func main(){
	var username string = "Sandesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n",isLoggedIn)

	var smallValue uint = 256
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n",smallValue)

	var smallFloatValue float32 = -256.0123456789
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type : %T \n",smallFloatValue)

	// default Values and aliases
	var declared int;
	fmt.Printf("Variable is of Intialized Int type : %T \n",declared)
	fmt.Println(declared)

	var declaredString string;
	fmt.Printf("Variable is of Intialized String type : %T \n",declaredString)
	fmt.Println(declaredString)

	// implicit type

	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n",website)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n",LoginToken)

	// no var style

	dog := "Dog"
	fmt.Println(dog)

}