package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	header = "MAMAMIA"
)

func main() {
	// Connect to the server at localhost:1337
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}

	defer conn.Close()

	// Create the packet body
	body := []byte("Hello, World!")

	// Create the packet header
	headerBytes := []byte(header)

	// Calculate the content length
	contentLength := uint32(len(body))

	// Send the packet to the server
	if err := binary.Write(conn, binary.BigEndian, headerBytes); err != nil {
		fmt.Printf("Error sending packet header: %v\n", err)
		return
	}

	if err := binary.Write(conn, binary.BigEndian, contentLength); err != nil {
		fmt.Printf("Error sending content length: %v\n", err)
		return
	}

	if _, err := conn.Write(body); err != nil {
		fmt.Printf("Error sending packet body: %v\n", err)
		return
	}

	fmt.Println("Packet sent!")

	// Wait for a response from the server
	response := make([]byte, 1024)
	if _, err := conn.Read(response); err != nil {
		fmt.Printf("Error receiving response: %v\n", err)
		return
	}

	// JACOB: had to change this to %s to display string rather than bytes
	fmt.Printf("Response received: %s\n", response)
}
