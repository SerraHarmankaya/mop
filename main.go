package main

import (
	"bufio"
	"fmt"
	"mop/module"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the API method (GET, POST, PUT, DELETE): ")
	method, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Method input error:", err)
		return
	}

	fmt.Print("Enter the URL (e.g. http://example.com): ")
	URL, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("URL input error:", err)
		return
	}

	api, err := module.ApiAnalyzer(method, URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status Code: ", api.StatusCode)
	fmt.Println("Response Time: ", api.ResponseTime)
	fmt.Println("Content Type: ", api.ContentType)
	fmt.Println("Body Size: ", api.BodySize)

}
