package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println(123)
	s := add(1, 2)
	fmt.Println(s)

	// resp, err := http.Get("https://www.baidu.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	var contents []byte
	// 	resp.Body.Read(contents)
	// 	fmt.Println(resp.ContentLength)
	// 	fmt.Println(contents)
	// }

	// client := new(http.Client)
	// response, err := client.Get("https://www.google.com")

	// if err != nil {
	// 	fmt.Println("error: ", err.Error())
	// } else {
	// 	fmt.Println(response.ContentLength)
	// 	for key, value := range response.Header {
	// 		fmt.Println(key)
	// 		fmt.Println(value)
	// 	}
	// }

	var uri = new(url.URL)
	uri.Path = "/"
	uri.Scheme = "http"
	uri.Host = "www.baidu.com"
	fmt.Println(uri.String())
}
