package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPhandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	listenAddr := ":4000"
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	// select {}
}
