package connection

// PacketStream describes packet sender and receiver
type PacketStream interface {
	SendPacket(cmd uint8, data []byte) (err error)
	RecvPacket() (cmd uint8, buf []byte, err error)
}
