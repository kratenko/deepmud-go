package server

import "net"

type Stage uint

const (
	STAGE_USERNAME = iota
	STAGE_PASSWORD
	STATE_PLAYING
)

type client struct {
	conn   net.Conn
	server *Server
	id     uint
	stage Stage
	username string
	password string
}


func (c *client)SendTextLn(s string) {
	c.conn.Write([]byte(s + "\n"))
}

func (c *client)SendText(s string) {
	c.conn.Write([]byte(s ))
}