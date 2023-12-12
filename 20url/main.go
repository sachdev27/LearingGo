package main

import (
	"fmt"
	"net/url"
)


const myurl = "https://sachdev27.github.io:8000/learn?coursename=reactjs&paymentid=dfg5678"

func main() {
	fmt.Println("Welcome to URL")
	fmt.Println(myurl)


	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)


	qparams := result.Query()
	fmt.Printf("qparams is of type : %T\n",qparams)

	fmt.Println(qparams)


	partsofurl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawPath: "user=sandesh",
	}

	anotherUrl := partsofurl.String()
	fmt.Println("Url is : ",anotherUrl)
}
