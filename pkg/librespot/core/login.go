package core

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/nezorflame/librespot-golang/pkg/librespot/connection"
	"github.com/nezorflame/librespot-golang/pkg/librespot/crypto"
	"github.com/nezorflame/librespot-golang/pkg/librespot/discovery"
	"github.com/nezorflame/librespot-golang/pkg/librespot/mercury"
	"github.com/nezorflame/librespot-golang/pkg/librespot/utils"
	"github.com/nezorflame/librespot-golang/pkg/spotify"

	"github.com/golang/protobuf/proto"
)

const (
	version = "master"
	buildID = "dev"
)

// NewSession creates new instance of Session
func NewSession() (*Session, error) {
	session := &Session{
		keys:               crypto.GenerateKeys(),
		mercuryConstructor: mercury.CreateMercury,
		shannonConstructor: crypto.CreateStream,
	}
	return session, session.doConnect()
}

// Login allows to log into Spotify using username and password
func (s *Session) Login(username, password, deviceName string) error {
	if err := s.setDeviceInfoAndConnect(deviceName); err != nil {
		return err
	}

	loginPacket := makeLoginPasswordPacket(username, password, s.deviceID)
	return s.doLogin(loginPacket, username)
}

// LoginSaved allows to log into Spotify using an existing authData blob
func (s *Session) LoginSaved(username string, authData []byte, deviceName string) error {
	if err := s.setDeviceInfoAndConnect(deviceName); err != nil {
		return err
	}

	packet := makeLoginBlobPacket(
		username,
		authData,
		spotify.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(),
		s.deviceID,
	)
	return s.doLogin(packet, username)
}

// LoginDiscovery registers librespot as a Spotify Connect device via mdns.
// When user connects, logs on to Spotify and saves credentials in file at cacheBlobPath.
// Once saved, the blob credentials allow the program to connect to other
// Spotify Connect devices and control them.
func (s *Session) LoginDiscovery(cacheBlobPath string, deviceName string) error {
	deviceID := utils.GenerateDeviceID(deviceName)
	disc := discovery.LoginFromConnect(cacheBlobPath, deviceID, deviceName)
	return s.sessionFromDiscovery(disc)
}

// LoginDiscoveryBlob allows to login using an authentication blob through
// Spotify Connect discovery system, reading an existing blob data.
// To read from a file, see LoginDiscoveryBlobFile.
func (s *Session) LoginDiscoveryBlob(username, blob, deviceName string) error {
	deviceID := utils.GenerateDeviceID(deviceName)
	disc := discovery.CreateFromBlob(utils.BlobInfo{
		Username:    username,
		DecodedBlob: blob,
	}, "", deviceID, deviceName)
	return s.sessionFromDiscovery(disc)
}

// LoginDiscoveryBlobFile allows to login from credentials at cacheBlobPath previously saved by LoginDiscovery.
// Similar to LoginDiscoveryBlob, except it reads it directly from a file.
func (s *Session) LoginDiscoveryBlobFile(cacheBlobPath, deviceName string) error {
	deviceID := utils.GenerateDeviceID(deviceName)
	disc := discovery.CreateFromFile(cacheBlobPath, deviceID, deviceName)
	return s.sessionFromDiscovery(disc)
}

// LoginOAuth allows to login to Spotify using the OAuth method
func (s *Session) LoginOAuth(deviceName, clientID, clientSecret string) error {
	token := getOAuthToken(clientID, clientSecret)
	return s.loginOAuthToken(token.AccessToken, deviceName)
}

func (s *Session) loginOAuthToken(accessToken string, deviceName string) error {
	if err := s.setDeviceInfoAndConnect(deviceName); err != nil {
		return err
	}

	packet := makeLoginBlobPacket(
		"",
		[]byte(accessToken),
		spotify.AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN.Enum(),
		s.deviceID,
	)
	return s.doLogin(packet, "")
}

func (s *Session) setDeviceInfoAndConnect(deviceName string) error {
	s.deviceID = utils.GenerateDeviceID(deviceName)
	s.deviceName = deviceName

	return s.startConnection()
}

func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(connection.PacketLogin, packet)
	if err != nil {
		log.Fatal("bad shannon write", err)
	}

	// Pll once for authentication response
	welcome, err := s.handleLogin()
	if err != nil {
		return err
	}

	// Store the few interesting values
	s.username = welcome.GetCanonicalUsername()
	if s.username == "" {
		// Spotify might not return a canonical username, so reuse the blob's one instead
		s.username = s.discovery.LoginBlob().Username
	}
	s.reusableAuthBlob = welcome.GetReusableAuthCredentials()

	// Poll for acknowledge before loading - needed for gopherjs
	// s.poll()
	go s.runPollLoop()

	return nil
}

func (s *Session) handleLogin() (*spotify.APWelcome, error) {
	cmd, data, err := s.stream.RecvPacket()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	if cmd == connection.PacketAuthFailure {
		return nil, fmt.Errorf("authentication failed")
	} else if cmd == connection.PacketAPWelcome {
		welcome := &spotify.APWelcome{}
		err := proto.Unmarshal(data, welcome)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %v", err)
		}
		fmt.Println("Authentication succeeded: Welcome,", welcome.GetCanonicalUsername())
		fmt.Println("Blob type:", welcome.GetReusableAuthCredentialsType())
		return welcome, nil
	} else {
		return nil, fmt.Errorf("authentication failed: unexpected cmd %v", cmd)
	}
}

func (s *Session) getLoginBlobPacket(blob utils.BlobInfo) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(blob.DecodedBlob)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(data)
	if _, err = buffer.ReadByte(); err != nil {
		return nil, err
	}
	readBytes(buffer)

	if _, err = buffer.ReadByte(); err != nil {
		return nil, err
	}
	authNum := readInt(buffer)

	authType := spotify.AuthenticationType(authNum)
	if _, err = buffer.ReadByte(); err != nil {
		return nil, err
	}
	authData := readBytes(buffer)

	return makeLoginBlobPacket(blob.Username, authData, &authType, s.deviceID), nil
}

func makeLoginPasswordPacket(username, password, deviceID string) []byte {
	return makeLoginBlobPacket(username, []byte(password), spotify.AuthenticationType_AUTHENTICATION_USER_PASS.Enum(), deviceID)
}

func makeLoginBlobPacket(username string, authData []byte, authType *spotify.AuthenticationType, deviceID string) []byte {
	versionString := "librespot-golang_" + version + "_" + buildID

	packet := &spotify.ClientResponseEncrypted{
		LoginCredentials: &spotify.LoginCredentials{
			Username: proto.String(username),
			Typ:      authType,
			AuthData: authData,
		},
		SystemInfo: &spotify.SystemInfo{
			CpuFamily:               spotify.CpuFamily_CPU_UNKNOWN.Enum(),
			Os:                      spotify.Os_OS_UNKNOWN.Enum(),
			SystemInformationString: proto.String("librespot-golang"),
			DeviceId:                proto.String(deviceID),
		},
		VersionString: proto.String(versionString),
	}

	packetData, err := proto.Marshal(packet)
	if err != nil {
		log.Fatal("login marshaling error: ", err)
	}

	return packetData
}
