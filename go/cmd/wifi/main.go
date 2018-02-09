package main

import (
	"../../src/wifi"
	"strconv"
	"os"
)

func main(){
	shutdownChan := make(chan struct{})

	var loss float64 = 0
	if len(os.Args) > 1 {
		l, err := strconv.ParseFloat(os.Args[1], 64)
		if err == nil {
			loss = l
		}
	}

	wifi.StartWifi(shutdownChan,loss)
}