package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define flags
	interfaceFlag := flag.String("interface", "", "Specify the network interface name")
	packetCountFlag := flag.Int("count", 0, "Specify the number of packets to capture (0 for unlimited)")

	// Parse command-line flags
	flag.Parse()

	// Check if the required flag is provided
	if *interfaceFlag == "" {
		fmt.Println("Usage: ./progName --interface <interface name> [--count <number>]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Open the network device
	handle, err := openDev(*interfaceFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Print out the packet count flag to verify it's being read correctly
	fmt.Printf("Capturing packets on interface %s with count %d\n", *interfaceFlag, *packetCountFlag)

	// Capture packets with the specified count
	captureLive(handle, *packetCountFlag)

	fmt.Println("Packet capture completed.")
}
