package server

import (
	"log"
	"net"
	"net/http"
)

type HttpServer struct{}

func (s *HttpServer) Serve(lis net.Listener) error {
	return http.Serve(lis, &OkHandler{})
}

type OkHandler struct{}

func (h OkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Received ping")
}
