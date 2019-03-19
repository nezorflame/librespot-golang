package connection

import (
	"encoding/binary"
	"io"
	"sync"
)

// PlainConnection represents an unencrypted connection to a Spotify AP
type PlainConnection struct {
	io.Reader
	io.Writer
	mtx sync.Mutex
}

func makePacketPrefix(prefix []byte, data []byte) []byte {
	size := len(prefix) + 4 + len(data)
	buf := make([]byte, 0, size)
	buf = append(buf, prefix...)
	sizeBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuf, uint32(size))
	buf = append(buf, sizeBuf...)
	return append(buf, data...)
}

// MakePlainConnection creates new PlainConnection
func MakePlainConnection(reader io.Reader, writer io.Writer) PlainConnection {
	return PlainConnection{
		Reader: reader,
		Writer: writer,
	}
}

// SendPrefixPacket sends new prefix packet with data
func (p *PlainConnection) SendPrefixPacket(prefix, data []byte) ([]byte, error) {
	packet := makePacketPrefix(prefix, data)

	p.mtx.Lock()
	_, err := p.Write(packet)
	p.mtx.Unlock()

	return packet, err
}

// RecvPacket receives a packet
func (p *PlainConnection) RecvPacket() ([]byte, error) {
	var size uint32
	if err := binary.Read(p, binary.BigEndian, &size); err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	binary.BigEndian.PutUint32(buf, size)
	if _, err := io.ReadFull(p, buf[4:]); err != nil {
		return nil, err
	}

	return buf, nil
}
