package world

type Room struct {
	BaseEntity
	contents []Entity
}

func NewRoom(id EntityId) *Room {
	return &Room{
		BaseEntity: BaseEntity{id: id},
		contents:   make([]Entity, 0),
	}
}
