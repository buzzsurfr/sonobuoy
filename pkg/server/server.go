package server

import "net"

type EchoServer interface {
	Serve(net.Listener) error
}
