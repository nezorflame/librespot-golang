package librespotmobile

import (
	"github.com/nezorflame/librespot-golang/pkg/spotify"
	"github.com/nezorflame/librespot-golang/pkg/librespot/core"
	"github.com/nezorflame/librespot-golang/pkg/librespot/player"
)

// MobilePlayer is a gomobile-compliant subset of the Player struct.
type MobilePlayer struct {
	player *player.Player
}

func createMobilePlayer(session *core.Session) *MobilePlayer {
	return &MobilePlayer{
		player: session.Player(),
	}
}

func (p *MobilePlayer) LoadTrack(fileID []byte, format int, trackID []byte) (*MobileAudioFile, error) {
	// Make a copy of the fileID and trackID byte arrays, as they may be freed/reused on the other end,
	// causing the fileID and/or trackID to change abruptly when the player actually request chunks.
	safeFileID := make([]byte, len(fileID))
	safeTrackID := make([]byte, len(trackID))
	copy(safeFileID, fileID)
	copy(safeTrackID, trackID)

	track, err := p.player.LoadTrackWithIdAndFormat(safeFileID, spotify.AudioFile_Format(format), safeTrackID)
	if err != nil {
		return nil, err
	}

	return createMobileAudioFile(track), nil
}
