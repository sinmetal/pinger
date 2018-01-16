package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const t = "http://10.128.0.2"

	for {
		time.Sleep(time.Second * 3)
		resp, err := http.Get(t)
		if err != nil {
			fmt.Printf("GET request to %s and the result was failed\n", t)
			continue
		}
		fmt.Printf("GET request to %s and the result was %s\n", t, resp.Status)
	}
}
