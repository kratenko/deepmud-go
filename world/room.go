package world

type Room struct {
	BaseEntity
	contents    []Entity
	Name        string
	Description string
}

func NewRoom(id EntityId) *Room {
	return &Room{
		BaseEntity: BaseEntity{id: id},
		contents:   make([]Entity, 0),
	}
}

func (r *Room) GetName() string {
	return r.Name
}

func (r *Room) GetDescription() string {
	return r.Description
}
