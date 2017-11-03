package main

import (
	"../../../src/wifi"
	"../../../src/feu"
)

func main(){
	shutdownWifi := make(chan struct{})

	go wifi.StartWifi(shutdownWifi)
	go feu.NewFeu("localhost:1234")

	select { }
}