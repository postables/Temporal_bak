package testutil

import (
	"testing"

	ma "gx/ipfs/QmYmsdtJ3HsodkePE3eU3TsCaP2YvPZJ4LoXnNkDE5Tpt7/go-multiaddr"
	peer "gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
	ci "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
)

type Identity interface {
	Address() ma.Multiaddr
	ID() peer.ID
	PrivateKey() ci.PrivKey
	PublicKey() ci.PubKey
}

// TODO add a cheaper way to generate identities

func RandIdentity() (Identity, error) {
	p, err := RandPeerNetParams()
	if err != nil {
		return nil, err
	}
	return &identity{*p}, nil
}

func RandIdentityOrFatal(t *testing.T) Identity {
	p, err := RandPeerNetParams()
	if err != nil {
		t.Fatal(err)
	}
	return &identity{*p}
}

// identity is a temporary shim to delay binding of PeerNetParams.
type identity struct {
	PeerNetParams
}

func (p *identity) ID() peer.ID {
	return p.PeerNetParams.ID
}

func (p *identity) Address() ma.Multiaddr {
	return p.Addr
}

func (p *identity) PrivateKey() ci.PrivKey {
	return p.PrivKey
}

func (p *identity) PublicKey() ci.PubKey {
	return p.PubKey
}

// NewIdentity constructs a new identity object with specific parameters
func NewIdentity(ID peer.ID, addr ma.Multiaddr, privk ci.PrivKey, pubk ci.PubKey) Identity {
	return &identity{PeerNetParams{ID: ID, Addr: addr, PrivKey: privk, PubKey: pubk}}
}
