package server

import (
	"net"

	"github.com/kratenko/deepmud-go/world"
)

type Stage uint

const (
	STAGE_USERNAME = iota
	STAGE_PASSWORD
	STATE_PLAYING
)

type client struct {
	conn     net.Conn
	server   *Server
	id       uint
	stage    Stage
	username string
	password string
	player   *world.Player // Only available after login, else nil
}

type Client interface {
	SendTextLn(s string)
	SendText(s string)
	GetUsername() string
}

func (c *client) SendTextLn(s string) {
	c.conn.Write([]byte(s + "\n"))
}

func (c *client) SendText(s string) {
	c.conn.Write([]byte(s))
}

func (c *client) GetUsername() string {
	return c.username
}
