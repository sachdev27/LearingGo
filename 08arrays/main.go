package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Go Lang")

	var strArray [4]string ;
	strArray[0] = "Sandesh"
	strArray[1] = "Crysys"
	strArray[3] = "Hev#"
	fmt.Println("First Element : ",strArray[0])
	fmt.Println("First Element : ",strArray[1])
	fmt.Println("First Element : ",strArray[2])
	fmt.Println("First Element : ",strArray[3])
	fmt.Println("Array Element : ",strArray)
	fmt.Println("Array Element len : ", len(strArray))

	var vegList = [3]string{"potato", "tomato", "onion"}
	fmt.Println("Array Element : ",vegList)
	fmt.Println("Array Element len : ", len(vegList))


}
