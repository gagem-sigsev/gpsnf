package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

// captureLive captures packets from the provided handle
func captureLive(handle *pcap.Handle, packetCount int) {
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	count := 0
	for packet := range packetSource.Packets() {
		decodePacket(packet)
		count++

		// If packetCount is greater than 0, stop after capturing the specified number of packets
		if packetCount > 0 && count >= packetCount {
			log.Printf("Captured %d packets, stopping.\n", count)
			break
		}
	}
}
