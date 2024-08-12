package main

import (
	_ "bytes"
	_ "fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/pcap"
)

func decodePacket(packet gopacket.Packet) {
  ethLayer := packet.Layer(layers.LayerTypeEthernet)  // ethLayer is of type gopacket.Layer
  if ethLayer != nil {
    printLayer(ethLayer.LayerType(), ethLayer) 
  }
}
