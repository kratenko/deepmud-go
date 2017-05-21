package server

import (
	"bufio"
	"net"

	"sync"

	"github.com/Sirupsen/logrus"
)

type Server struct {
	mu              sync.Mutex
	currentClientId uint
	clients         map[net.Conn]*client
}

func ListenAndServe(addr string) error {
	server := &Server{clients: make(map[net.Conn]*client)}
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
		s.currentClientId++
		client := &client{conn: conn, server: s, id: s.currentClientId}
		s.mu.Lock()
		s.clients[conn] = client
		s.mu.Unlock()
		logrus.WithField("id", client.id).WithField("addr", conn.RemoteAddr()).Info("New connection")
		// Handle connections in a new goroutine.
		go handleClient(client)
	}
	return nil
}

func handleClient(client *client) {
	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		logrus.WithField("id", client.id).WithField("text", scanner.Text()).Info("Got text")
	}
	logrus.WithField("id", client.id).Info("Client connection broke")
	client.server.removeClient(client)
}

func (s *Server) removeClient(client *client) {
	s.mu.Lock()
	delete(s.clients, client.conn)
	s.mu.Unlock()
}
