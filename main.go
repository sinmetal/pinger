package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://10.142.0.2")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
