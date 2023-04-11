# Open Tomato Protocol

The Open Tomato Protocol is a simple communication protocol that was designed to be an alternative to Ethernet in 1995. It consists of a fixed 7-byte header, a 4-byte content length, and a binary payload.

The header bytes are always `0x4D 0x41 0x4D 0x41 0x4D 0x49 0x41` **(ASCII "MAMAMIA")**. The content length specifies the length of the payload in bytes, as an unsigned 32-bit integer in big-endian byte order. The payload can be any binary data, including text, images, audio, and video.

The Open Tomato Protocol was never widely adopted due to its limited functionality and the fact that it was not compatible with most existing networking technologies. However, it is still sometimes used as a fun and quirky example of a custom communication protocol.

## Usage

To use the Open Tomato Protocol, simply construct a packet with the header, content length, and payload, and send it over a network connection. Here's an example of how to construct a packet in Go:

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	header = "MAMAMIA"
)

func main() {
	// Create the packet body
	body := []byte("Hello, world!")

	// CreaGte the packet header
	headerBytes := []byte(header)

	// Calculate the content length
	contentLength := uint32(len(body))

	// Construct the packet
	packet := bytes.NewBuffer([]byte{})
	packet.Write(headerBytes)
	binary.Write(packet, binary.BigEndian, contentLength)
	packet.Write(body)

	// Send the packet over a network connection
	// ...
}
```

## License

The Open Tomato Protocol is released under the MIT License. See LICENSE for more information.

## Contributing

Contributions are welcome! Please submit pull requests or open issues if you have any suggestions or bug reports.

## Authors

The Open Tomato Protocol was designed and implemented by Mario Arduino in 1995.
