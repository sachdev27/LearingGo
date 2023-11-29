package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning About Time")
	fmt.Println("Time Now is :",time.Now().Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2018,time.February,12,11,17,00,0,time.UTC)
	fmt.Println("Created Time : ",createdDate)
	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))
}
