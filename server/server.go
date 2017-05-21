package server

import (
	"bufio"
	"net"

	"github.com/Sirupsen/logrus"
)

type Server struct {
}

var Default *Server

func ListenAndServe(addr string) error {
	server := &Server{}
	return server.ListenAndServe(addr)
}

func (s *Server) ListenAndServe(addr string) error {

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to start server")
		return err
	}
	defer l.Close()

	logrus.Info("Listening on " + addr)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			logrus.WithError(err).Error("Error on accept")
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
	return nil
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		logrus.WithField("text", scanner.Text()).Info("Got text")
	}
}
