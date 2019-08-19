package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/wangjiaming0909/GoP/server"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	for headerName, headerValue := range r.Header {
		fmt.Println(headerName)
		fmt.Println(headerValue)
	}

	w.Write([]byte("Hello World!"))
}

func main() {
	dispather := server.NewDispatcher()
	err := dispather.RegisterHandler("/", HttpHandler, false)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

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
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":9090", nil)

}
