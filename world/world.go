package world

type World struct {
	currentId EntityId
	entities  map[EntityId]Entity
}

func (w *World) NextId() EntityId {
	w.currentId++
	return w.currentId
}

func (w *World) InsertEntity(e Entity) {
	w.entities[e.GetId()] = e
}

func (w *World) RemoveEntity(e Entity) {
	delete(w.entities, e.GetId())
}

func New() *World {
	w := &World{}
	r1 := NewRoom(w.NextId())
	r1.Description = "Du stehst auf einem Dorfplatz"
	r1.Name = "Dorfplatz"

	w.InsertEntity(r1)
	return w
}
