package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	header        = "MAMAMIA"
	packetMaxSize = 1024
)

func main() {
	// Listen for TCP connections on port 1337
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}

	fmt.Println("Listening for connections on port 1337...")

	// Accept incoming connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a buffer and reader for the connection
	// buffer := make([]byte, packetMaxSize)
	// reader := bytes.NewReader(buffer)
	// JACOB: ^^^ those were generated and then never used. funny.

	// Process incoming packets
	for {
		// Read the packet header bytes
		headerBytes := make([]byte, len(header))
		if _, err := io.ReadFull(conn, headerBytes); err != nil {
			if err == io.EOF {
				// End of stream
				break
			} else {
				fmt.Printf("Error reading packet header: %v\n", err)
				return
			}
		}

		// Verify the header bytes
		if !bytes.Equal(headerBytes, []byte(header)) {
			fmt.Printf("Invalid packet header: %v\n", headerBytes)
			return
		}

		// Read the content length
		var contentLength uint32
		if err := binary.Read(conn, binary.BigEndian, &contentLength); err != nil {
			fmt.Printf("Error reading content length: %v\n", err)
			return
		}

		// Read the body data
		if contentLength > uint32(packetMaxSize-len(header)-4) {
			fmt.Printf("Packet body too large: %d bytes\n", contentLength)
			return
		}

		body := make([]byte, contentLength)
		if _, err := io.ReadFull(conn, body); err != nil {
			fmt.Printf("Error reading packet body: %v\n", err)
			return
		}

		// Process the body data
		fmt.Printf("Packet received: body=%s\n", body)
		// JACOB: ^ had to change that to %s to display string rather than bytes

		// JACOB: and for this I had to reply for the client to terminate successfully.
		conn.Write([]byte(`OK`))
	}

	fmt.Println("Connection closed")
}
