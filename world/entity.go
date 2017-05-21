package world

type EntityId uint64

type Entity interface {
	GetId() EntityId
	GetName() string
	GetDescription() string
}

type BaseEntity struct {
	id          EntityId
}

func (e *BaseEntity) GetId() EntityId {
	return e.id
}


