package p2p

type HandshakeFunc func(Peer) error

func NOPhandshakeFunc(Peer) error { return nil }
