package main

import (
	"encoding/json"
	"fmt"

	"github.com/sparrc/go-ping"
)

func main() {
	pinger, err := ping.NewPinger("10.128.0.3")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	b, err := json.Marshal(stats)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
