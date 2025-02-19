package main

type World struct {
	rooms       []Room
	position    Vector2d
	currentRoom int
}

func NewWorld() *World {
	return &World{
		rooms:       []Room{},
		position:    Vector2d{2, 2},
		currentRoom: 0,
	}
}

func (w *World) Generate() {
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{0, 0}, Vector2d{30, 10}, true, true, true, true, true,
	))
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{1, 0}, Vector2d{30, 10}, true, true, true, true, true,
	))
}

func (w *World) GetStartingPosition() Vector2d {
	if len(w.rooms) > 0 {
		topLeft := V.Sum(w.position, V.Zero)
		return V.Sum(
			topLeft,
			Vector2d{
				w.rooms[0].size.x/2 - 0,
				w.rooms[0].size.y/2 - 0,
			},
		)
	}
	return V.Sum(w.position, V.Identity)
}

func (w *World) GetRoomInnerBounds(index int) (Vector2d, Vector2d) {
	if len(w.rooms) < (index + 1) {
		return w.position, w.position
	}
	topLeft := V.Sum(w.position, V.Zero)
	return topLeft, V.Sum(
		topLeft,
		Vector2d{
			w.rooms[0].size.x - 1,
			w.rooms[0].size.y - 1,
		},
	)
}

func (w *World) Draw(c *Canvas) {
	for _, r := range w.rooms {
		drawPos := V.Sum(w.position, Vector2d{r.pos.x * (r.size.x + 1), 0})
		r.Draw(c, drawPos)
	}
}
