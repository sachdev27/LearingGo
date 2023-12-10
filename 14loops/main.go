package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	days := []string{"Sunday","Monday","Tuesday","Wednesday","Friday","Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println("Day : ",days[i])
	}


	// for i := range days {
	// 	fmt.Println("i is : ",i)
	// }

	// for i,day := range days {
	// 	fmt.Printf("i is %v & day is %v\n",i,day)
	// }

	// for _ ,day := range days {
	// 	fmt.Printf("i is _ & day is %v\n",day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2{
			goto web
		}

		if rougeValue == 5 {
			rougeValue++
			continue       // can use continue , break
		}

		fmt.Println("Value is :",rougeValue)
		rougeValue++
	}

	web:
		fmt.Println("Jumping at sachdev27.github.io")

}
