package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Go")
	// fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice is ",diceNumber)


	switch diceNumber{
	case 1:
		fmt.Println("Dice Number is 1 you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again")
	default:
		fmt.Println("what was that!")
	}
}
