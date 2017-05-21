package server

import "strings"

type Message struct {
	client *client
	Text   string
}

func (m Message) GetAction() string {
	words := strings.SplitN(m.Text, " ", 1)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}
