package server

import "net"

type Server struct {

}


func (s *Server)ListenAndServe(addr string) {


	l, err := net.Listen("tcp", addr)
}