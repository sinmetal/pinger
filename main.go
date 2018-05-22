package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const t = "http://instance-1.c.long-parity-204323.internal"

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
