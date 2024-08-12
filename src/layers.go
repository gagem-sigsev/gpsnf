package main

import (
    "errors"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
)

// ExtractLayers extracts various layers from a given packet.
// It returns each layer and an error if any layer could not be extracted.
func ExtractLayers(packet gopacket.Packet) (ethLayer *layers.Ethernet, ipLayer *layers.IPv4, tcpLayer *layers.TCP, udpLayer *layers.UDP, err error) {
    // Extract Ethernet layer
    if ethLayerIface := packet.Layer(layers.LayerTypeEthernet); ethLayerIface != nil {
        if ethLayer, ok := ethLayerIface.(*layers.Ethernet); ok {
            // Successfully casted to *layers.Ethernet
        } else {
            return nil, nil, nil, nil, errors.New("failed to cast to *layers.Ethernet")
        }
    }

    // Extract IP layer
    if ipLayerIface := packet.Layer(layers.LayerTypeIPv4); ipLayerIface != nil {
        if ipLayer, ok := ipLayerIface.(*layers.IPv4); ok {
            // Successfully casted to *layers.IPv4
        } else {
            return nil, nil, nil, nil, errors.New("failed to cast to *layers.IPv4")
        }
    }

    // Extract TCP layer
    if tcpLayerIface := packet.Layer(layers.LayerTypeTCP); tcpLayerIface != nil {
        if tcpLayer, ok := tcpLayerIface.(*layers.TCP); ok {
            // Successfully casted to *layers.TCP
        } else {
            return nil, nil, nil, nil, errors.New("failed to cast to *layers.TCP")
        }
    }

    // Extract UDP layer
    if udpLayerIface := packet.Layer(layers.LayerTypeUDP); udpLayerIface != nil {
        if udpLayer, ok := udpLayerIface.(*layers.UDP); ok {
            // Successfully casted to *layers.UDP
        } else {
            return nil, nil, nil, nil, errors.New("failed to cast to *layers.UDP")
        }
    }

    // If all layers are successfully extracted, return them along with a nil error
    return ethLayer, ipLayer, tcpLayer, udpLayer, nil
}
