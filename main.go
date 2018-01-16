package main

import (
	"fmt"
	"net/http"
)

func main() {
	const t = "http://10.129.0.4"
	resp, err := http.Get(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GET request to %s and the result was %s\n", t, resp.Status)
}
