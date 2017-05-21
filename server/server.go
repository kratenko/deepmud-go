package server

import (
	"bufio"
	"net"

	"sync"

	"github.com/Sirupsen/logrus"
	"fmt"
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
		go handleClient(s, client)
	}
	return nil
}

func handleClient(server *Server, client *client) {
	client.SendTextLn("Willkommen in DeepMud. Wer bist du?")
	
	

	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		logrus.WithField("id", client.id).WithField("text", scanner.Text()).Info("Got text")
		msg := &Message{client, scanner.Text()}
		server.handleMessage(msg);
	}
	logrus.WithField("id", client.id).Info("Client connection broke")
	client.server.removeClient(client)
}

func (server Server) handleMessage(m *Message) {
	switch m.client.stage {
	case STAGE_USERNAME:
		m.client.username = m.Text
		m.client.SendTextLn(fmt.Sprintf("Hallo %s, nenne mir dein größtes Geheimnis", m.client.username))
		m.client.stage = STAGE_PASSWORD
	case STAGE_PASSWORD:
		m.client.password = m.Text
		m.client.SendTextLn("Willkommen in der Welt von DeepMud, schau dich in Ruhe um.")
		m.client.stage = STATE_PLAYING
	}
}

func (s *Server) removeClient(client *client) {
	s.mu.Lock()
	delete(s.clients, client.conn)
	s.mu.Unlock()
}
