package main

import (
	"fmt"
	"log"
  "flag"
  "os"
)

func main() {

  // Define a string flag with a default value and a usage description.
  interfaceFlag := flag.String("interface", "", "Specify the network interface name")

  // Parse the command-line flags.
  flag.Parse()

  // Check if the required flag is provided.
  if *interfaceFlag == "" {
    // Print usage and exit if the flag is not provided.
    fmt.Println("Usage: ./progName --interface <interface name>")
    flag.PrintDefaults()
    os.Exit(1)
  }

	handle, err := openDev(*interfaceFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	captureLive(handle)

	fmt.Println("Handle to device created. Starting Packet Capture...")
}
