package world

import "github.com/kratenko/deepmud-go/server"

type Player struct {
	BaseEntity
	client server.Client
	name   string
}

func NewPlayer(world *World, client server.Client) *Player {
	p := &Player{
		BaseEntity: BaseEntity{id: world.NextId()},
		client:     client,
		name:       client.GetUsername(),
	}

	world.InsertEntity(p)
	return p
}

func (e *Player) GetName() string {
	return e.name
}

func (e *Player) GetDescription() string {
	return "Ein Spieler."
}

func (p *Player) handleMessage(m *server.Message) {
	// TODO: use Action from action.go to implement matching and support aliases
	switch m.GetAction() {
	case "sage":
	case "betrachte":
	}
}
