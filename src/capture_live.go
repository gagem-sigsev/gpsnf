package main

import (
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

func captureLive(handle *pcap.Handle) {
  // create a new packet source
  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    decodePacket(packet)
  }
}
