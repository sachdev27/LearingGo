package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("JSON in GO")
	DecodedJson()

}

type course struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"password"`
}

func DecodedJson() {
	//temp json data
	jsonDataFromWeb := []byte(`
	{
	"name": "Sandesh",
	"price": 25,
	"platform": "unix"
	}
	`)

	checkValid := json.Valid(jsonDataFromWeb)
	var lcocourse course
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Println(lcocourse)
		fmt.Printf("%#v", lcocourse)
	} else {
		fmt.Println("Json was not valid")
	}
}
