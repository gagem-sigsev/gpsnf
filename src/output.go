package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
)

func printLayer(layerType gopacket.LayerType, layer gopacket.Layer) {
    switch layerType {
    case layers.LayerTypeEthernet:
        if ethernetLayer, ok := layer.(*layers.Ethernet); ok {
            fmt.Println("Ethernet Packet:")
            fmt.Printf("  Source MAC: %s\n", ethernetLayer.SrcMAC)
            fmt.Printf("  Destination MAC: %s\n", ethernetLayer.DstMAC)
            fmt.Printf("  EtherType: %s\n", ethernetLayer.EthernetType)
        }

    case layers.LayerTypeIPv4:
        if ipLayer, ok := layer.(*layers.IPv4); ok {
            fmt.Println("IPv4 Packet:")
            fmt.Printf("  Source IP: %s\n", ipLayer.SrcIP)
            fmt.Printf("  Destination IP: %s\n", ipLayer.DstIP)
            fmt.Printf("  Protocol: %s\n", ipLayer.Protocol)
        }

    case layers.LayerTypeTCP:
        if tcpLayer, ok := layer.(*layers.TCP); ok {
            fmt.Println("TCP Packet:")
            fmt.Printf("  Source Port: %d\n", tcpLayer.SrcPort)
            fmt.Printf("  Destination Port: %d\n", tcpLayer.DstPort)
            fmt.Printf("  Sequence Number: %d\n", tcpLayer.Seq)
            fmt.Printf("  Acknowledgment Number: %d\n", tcpLayer.Ack)
            fmt.Printf("  SYN: %t\n", tcpLayer.SYN)
            fmt.Printf("  ACK: %t\n", tcpLayer.ACK)
            fmt.Printf("  FIN: %t\n", tcpLayer.FIN)
            fmt.Printf("  RST: %t\n", tcpLayer.RST)
            fmt.Printf("  PSH: %t\n", tcpLayer.PSH)
            fmt.Printf("  URG: %t\n", tcpLayer.URG)
            fmt.Printf("  ECE: %t\n", tcpLayer.ECE)
            fmt.Printf("  CWR: %t\n", tcpLayer.CWR)
        }

    case layers.LayerTypeUDP:
        if udpLayer, ok := layer.(*layers.UDP); ok {
            fmt.Println("UDP Packet:")
            fmt.Printf("  Source Port: %d\n", udpLayer.SrcPort)
            fmt.Printf("  Destination Port: %d\n", udpLayer.DstPort)
            fmt.Printf("  Length: %d\n", udpLayer.Length)
        }

    default:
        fmt.Printf("Unsupported Layer Type: %s\n", layerType)
    }
}

