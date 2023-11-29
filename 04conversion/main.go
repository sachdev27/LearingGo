package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ",rating)

	numrating,err := strconv.ParseFloat(strings.TrimSpace(rating),64) // Have to remove trailing \n from input

	if err != nil {
	fmt.Println("Your Error is ", err)
	// panic(err)
	} else {
		fmt.Println("Your rating after adding is ", numrating+1)
	}

}
