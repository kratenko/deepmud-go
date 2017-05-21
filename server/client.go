package server

import "net"

type client struct {
	conn   net.Conn
	server *Server
	id     uint
}
