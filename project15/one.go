package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("In this lesson we are going to learn web services")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting Get Request", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of Response:%v\n", res.StatusCode)
	// fmt.Printf("Response:%s\n", res.Body)      this is the wrong way to read response

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Reading Response :", err)
	}
	fmt.Println("Response Is:", string(data))
}
