package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func decodePacket(packet gopacket.Packet) {
	ethLayer := packet.Layer(layers.LayerTypeEthernet) // ethLayer is of type gopacket.Layer
	if ethLayer != nil {
		printLayer(ethLayer.LayerType(), ethLayer)
	}

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		printLayer(ipLayer.LayerType(), ipLayer)
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		printLayer(tcpLayer.LayerType(), tcpLayer)
	}

	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		printLayer(udpLayer.LayerType(), udpLayer)
	}
}
