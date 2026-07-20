package main

import (
	"bufio"
	"fmt"
	"mop/module"
	"net/http"
	"os"
	"strings"
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

	normalizedMethod := strings.ToUpper(strings.TrimSpace(method))

	var body string
	if strings.TrimSpace(method) != "" && (normalizedMethod == http.MethodPost || normalizedMethod == http.MethodPut) {
		fmt.Print("Enter request body (JSON): ")

		body, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Body input error:", err)
			return
		}

		body = strings.TrimSpace(body)
	}

	api, err := module.ApiAnalyzer(method, URL, body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(fmt.Sprintf(`Status Code: %s`, api.StatusMessage))
	fmt.Println("Response Time: ", api.ResponseTime)
	fmt.Println("Content Type: ", api.ContentType)
	fmt.Println("Body Size: ", api.BodySize)

}
