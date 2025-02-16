package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Today we Will learn to use NET/URL")
	myUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of myUrl Is %T\n:", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Cant Parse Url ", err)
	}
	fmt.Println("Parsed Url Is :", parsedUrl)
	fmt.Printf("Type of parsedUrl Is %T\n:", parsedUrl)
	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host: ", parsedUrl.Host)
	fmt.Println("Path: ", parsedUrl.Path)
	fmt.Println("RawQuery: ", parsedUrl.RawQuery)

	// Modify Your Url

	parsedUrl.Path = "newPath"
	fmt.Println("Parsed Url Is :", parsedUrl) // Modified url

}
