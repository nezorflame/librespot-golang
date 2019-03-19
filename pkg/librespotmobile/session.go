package librespotmobile

import "github.com/nezorflame/librespot-golang/pkg/librespot/core"

// MobileSession exposes a simplified subset of the core.Session struct that is compatible with the subset
// of types accepted by gomobile. Most calls are proxied to the underlying core.Session pointer, which we
// cannot expose directly as it uses types incompatible with gomobile.
type MobileSession struct {
	session *core.Session
	player  *MobilePlayer
	mercury *MobileMercury
}

// NewMobileSession creates new MobileSession
func NewMobileSession() (*MobileSession, error) {
	sess, err := core.NewSession()
	if err != nil {
		return nil, err
	}

	return &MobileSession{
		session: sess,
		player:  createMobilePlayer(sess),
		mercury: createMobileMercury(sess),
	}, nil
}

// Login wraps core.Session Login method
func (s *MobileSession) Login(username, password, deviceName string) error {
	return s.session.Login(username, password, deviceName)
}

// LoginSaved wraps core.Session LoginSaved method
func (s *MobileSession) LoginSaved(username string, authData []byte, deviceName string) error {
	return s.session.LoginSaved(username, authData, deviceName)
}

// Username returns core.Session Username
func (s *MobileSession) Username() string {
	return s.session.Username()
}

// DeviceID returns core.Session DeviceID
func (s *MobileSession) DeviceID() string {
	return s.session.DeviceID()
}

// ReusableAuthBlob returns core.Session ReusableAuthBlob
func (s *MobileSession) ReusableAuthBlob() []byte {
	return s.session.ReusableAuthBlob()
}

// Country returns core.Session Country
func (s *MobileSession) Country() string {
	return s.session.Country()
}

// Player returns core.Session Player
func (s *MobileSession) Player() *MobilePlayer {
	return s.player
}

// Mercury returns core.Session Mercury
func (s *MobileSession) Mercury() *MobileMercury {
	return s.mercury
}
