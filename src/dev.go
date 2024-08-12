package main

import (
	"errors"
	"github.com/google/gopacket/pcap"
)

func openDev(dev string) (*pcap.Handle, error) {
  // Open Device for capture
  handle, err := pcap.OpenLive(dev, 65535, true, pcap.BlockForever)
  if err != nil {
    return nil, errors.New("Unable to open device!")
  }

  return handle, nil
}
