package main

type World struct {
	rooms []Room
}

func NewWorld() *World {
	return &World{
		rooms: []Room{},
	}
}

func (w *World) Generate() {
	w.rooms = append(w.rooms, *NewRoom(
		0, 0, 30, 10, true, true, true, true, true,
	))
	w.rooms = append(w.rooms, *NewRoom(
		1, 0, 30, 10, true, true, true, true, true,
	))
}

func (w *World) Draw(c *Canvas) {
	for _, r := range w.rooms {
		r.Draw(c, 5+r.x*(r.w+1), 5)
	}
}
