package main

import (
	"fmt"
	"log"
)

func main() {
	handle, err := openDev("wlan0")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	captureLive(handle)

	fmt.Println("Handle to device created. Starting Packet Capture...")
}
