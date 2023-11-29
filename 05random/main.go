package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to Maths in Go Lang")

	// var numberOne int = 4
	// var numberTwo float64 = 4.5

	// fmt.Print("Sum is : ",numberOne + int(numberTwo))


	// random number

	// fmt.Println(time.Now().Unix())
	// rand.Seed(time.Now().Unix())
	// randint := rand.Intn(5) // [0,101)
	// fmt.Println("Random Number between 0-100 : ",randint)

	// Crypto Random
	
	randCryInt,_ := rand.Int(rand.Reader,big.NewInt(100))
	fmt.Println(randCryInt)
}
