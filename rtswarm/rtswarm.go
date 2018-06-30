package rtswarm

import (
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/swarm/api"
	"github.com/ethereum/go-ethereum/swarm/api/client"
)

// SwarmManager is a helper interface
type SwarmManager struct {
	Config     *api.Config
	Client     *client.Client
	PrivateKey *ecdsa.PrivateKey
}

// NewSwarmManager is used to generate our swarm manager helper interface
func NewSwarmManager() (*SwarmManager, error) {
	sm := &SwarmManager{}
	sm.GenSwarmAPIConfig()
	if err := sm.GenSwarmPrivateKeys(); err != nil {
		return nil, err
	}
	sm.GenSwarmClient()
	return sm, nil
}

func (sm *SwarmManager) DownloadManifest(hash string) (*api.Manifest, bool, error) {
	if hash == "" {
		return nil, false, errors.New("hash is empty")
	}

	manifest, isEncrypted, err := sm.DownloadManifest(hash)
	if err != nil {
		return nil, false, err
	}
	return manifest, isEncrypted, nil
}

func (sm *SwarmManager) UploadRaw(file *os.File, manifest string, encrypt bool) (string, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	resp, err := sm.Client.UploadRaw(file, fileInfo.Size(), encrypt)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (sm *SwarmManager) Upload(filePath, manifest string, encrypt bool) (string, error) {
	f, err := client.Open(filePath)
	if err != nil {
		return "", err
	}
	resp, err := sm.Client.Upload(f, manifest, encrypt)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// GenSwarmAPIConfig is used to generate a default swarm api configuration
func (sm *SwarmManager) GenSwarmAPIConfig() {
	sm.Config = api.NewConfig()
}

// GenSwarmPrivateKeys is used to generate our swarm private keys
func (sm *SwarmManager) GenSwarmPrivateKeys() error {
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	sm.PrivateKey = key
	return nil
}

// GenSwarmClient is used to generate our swarm client
func (sm *SwarmManager) GenSwarmClient() {
	sm.Client = client.NewClient(client.DefaultGateway)
}
