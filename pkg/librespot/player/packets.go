package player

import (
	"bytes"
	"encoding/binary"
)

func buildAudioChunkRequest(channel uint16, fileID []byte, start uint32, end uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, channel)
	binary.Write(buf, binary.BigEndian, uint8(0x0))
	binary.Write(buf, binary.BigEndian, uint8(0x1))
	binary.Write(buf, binary.BigEndian, uint16(0x0000))
	binary.Write(buf, binary.BigEndian, uint32(0x00000000))
	binary.Write(buf, binary.BigEndian, uint32(0x00009C40))
	binary.Write(buf, binary.BigEndian, uint32(0x00020000))
	buf.Write(fileID)
	binary.Write(buf, binary.BigEndian, start)
	binary.Write(buf, binary.BigEndian, end)

	return buf.Bytes()
}

func buildKeyRequest(seq []byte, trackID []byte, fileID []byte) []byte {
	buf := new(bytes.Buffer)

	buf.Write(fileID)
	buf.Write(trackID)
	buf.Write(seq)
	binary.Write(buf, binary.BigEndian, uint16(0x0000))

	return buf.Bytes()
}
