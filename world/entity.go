package world

type EntityId uint64

type Entity interface {
	GetId() EntityId
	GetName() string
	GetDescription() string
}

type BaseEntity struct {
	id          EntityId
	Name        string
	Description string
}

func (e *BaseEntity) GetId() EntityId {
	return e.id
}

func (e *BaseEntity) GetName() string {
	return e.Name
}

func (e *BaseEntity) GetDescription() string {
	return e.Description
}
