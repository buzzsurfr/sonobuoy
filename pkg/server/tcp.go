package server

import (
	"log"
	"net"
)

type TcpServer struct{}

func (s *TcpServer) Serve(lis net.Listener) error {
	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		// Handler
		go func(c net.Conn) {
			log.Print("Received ping")
			c.Close()
		}(conn)
	}
}
