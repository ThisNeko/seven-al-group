package main

import "../../src/wifi"

func main(){
	shutdownChan := make(chan struct{})
	wifi.StartWifi(shutdownChan)
}