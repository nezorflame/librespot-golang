package player

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"sync"

	"github.com/nezorflame/librespot-golang/pkg/librespot/connection"
	"github.com/nezorflame/librespot-golang/pkg/librespot/mercury"
	"github.com/nezorflame/librespot-golang/pkg/spotify"
)

type Player struct {
	stream   connection.PacketStream
	mercury  *mercury.Client
	seq      uint32
	audioKey []byte

	chanLock    sync.Mutex
	seqChanLock sync.Mutex
	channels    map[uint16]*Channel
	seqChans    sync.Map
	nextChan    uint16
}

func CreatePlayer(conn connection.PacketStream, client *mercury.Client) *Player {
	return &Player{
		stream:   conn,
		mercury:  client,
		channels: map[uint16]*Channel{},
		seqChans: sync.Map{},
		chanLock: sync.Mutex{},
		nextChan: 0,
	}
}

func (p *Player) LoadTrack(file *spotify.AudioFile, trackID []byte) (*AudioFile, error) {
	return p.LoadTrackWithIdAndFormat(file.GetFileId(), file.GetFormat(), trackID)
}

func (p *Player) LoadTrackWithIdAndFormat(fileID []byte, format spotify.AudioFile_Format, trackID []byte) (*AudioFile, error) {
	// fmt.Printf("[player] Loading track audio key, fileId: %s, trackId: %s\n", utils.ConvertTo62(fileID), utils.ConvertTo62(trackID))

	// Allocate an AudioFile and a channel
	audioFile := newAudioFileWithIdAndFormat(fileID, format, p)

	// Start loading the audio key
	err := audioFile.loadKey(trackID)

	// Then start loading the audio itself
	audioFile.loadChunks()

	return audioFile, err
}

func (p *Player) loadTrackKey(trackID []byte, fileID []byte) ([]byte, error) {
	seqInt, seq := p.mercury.NextSeqWithInt()

	p.seqChans.Store(seqInt, make(chan []byte))

	req := buildKeyRequest(seq, trackID, fileID)
	err := p.stream.SendPacket(connection.PacketRequestKey, req)
	if err != nil {
		log.Println("Error while sending packet", err)
		return nil, err
	}

	channel, _ := p.seqChans.Load(seqInt)
	key := <-channel.(chan []byte)
	p.seqChans.Delete(seqInt)

	return key, nil
}

func (p *Player) AllocateChannel() *Channel {
	p.chanLock.Lock()
	channel := NewChannel(p.nextChan, p.releaseChannel)
	p.nextChan++

	p.channels[channel.num] = channel
	p.chanLock.Unlock()

	return channel
}

func (p *Player) HandleCmd(cmd byte, data []byte) {
	switch {
	case cmd == connection.PacketAesKey:
		// Audio key response
		dataReader := bytes.NewReader(data)
		var seqNum uint32
		binary.Read(dataReader, binary.BigEndian, &seqNum)

		if channel, ok := p.seqChans.Load(seqNum); ok {
			channel.(chan []byte) <- data[4:20]
		} else {
			fmt.Printf("[player] Unknown channel for audio key seqNum %d\n", seqNum)
		}

	case cmd == connection.PacketAesKeyError:
		// Audio key error
		fmt.Println("[player] Audio key error!")
		fmt.Printf("%x\n", data)

	case cmd == connection.PacketStreamChunkRes:
		// Audio data response
		var channel uint16
		dataReader := bytes.NewReader(data)
		binary.Read(dataReader, binary.BigEndian, &channel)

		// fmt.Printf("[player] Data on channel %d: %d bytes\n", channel, len(data[2:]))

		if val, ok := p.channels[channel]; ok {
			val.handlePacket(data[2:])
		} else {
			fmt.Printf("Unknown channel!\n")
		}
	}
}

func (p *Player) releaseChannel(channel *Channel) {
	p.chanLock.Lock()
	delete(p.channels, channel.num)
	p.chanLock.Unlock()
	// fmt.Printf("[player] Released channel %d\n", channel.num)
}
