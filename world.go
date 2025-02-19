package main

type World struct {
	rooms       []Room
	position    Vector2d
	currentRoom int
}

func NewWorld() *World {
	return &World{
		rooms:       []Room{},
		position:    Vector2d{45, 15},
		currentRoom: 0,
	}
}

func (w *World) Generate() {
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{0, 0}, Vector2d{30, 10}, true, true, true, true, true,
	))
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{1, 0}, Vector2d{30, 10}, true, false, false, false, true,
	))
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{0, 1}, Vector2d{30, 10}, true, true, false, false, false,
	))
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{-1, 0}, Vector2d{30, 10}, true, false, true, false, false,
	))
	w.rooms = append(w.rooms, *NewRoom(
		Vector2d{0, -1}, Vector2d{30, 10}, true, false, false, true, false,
	))
}

func (w *World) GetRoom(pos Vector2d) (Room, bool) {
	for _, room := range w.rooms {
		if V.Equal(room.pos, pos) {
			return room, true
		}
	}

	return Room{}, false
}

func (w *World) GetStartingRoomIndex() int {
	return 0
}

func (w *World) GetStartingPosition() Vector2d {
	if len(w.rooms) > 0 {
		roomIndex := 1
		r := w.rooms[roomIndex]
		pad := V.Sum(r.size, V.Identity)
		topLeft := V.Sum(w.position, Vector2d{pad.x * r.pos.x, pad.y * r.pos.y})
		return V.Sum(
			topLeft,
			Vector2d{
				r.size.x/2 - 0,
				r.size.y/2 - 0,
			},
		)
	}
	return V.Sum(w.position, V.Identity)
}

func (w *World) GetRoomWoldPosition(pos Vector2d) Vector2d {
	r, _ := w.GetRoom(pos)
	pad := V.Sum(r.size, V.Identity)
	return Vector2d{
		x: pad.x * r.pos.x,
		y: pad.y * r.pos.y,
	}
}

func (w *World) GetRoomInnerBounds(index int) (Vector2d, Vector2d) {
	if len(w.rooms) < (index + 1) {
		return w.position, w.position
	}
	r := w.rooms[index]
	pad := V.Sum(r.size, V.Identity)
	topLeft := V.Sum(w.position, Vector2d{pad.x * r.pos.x, pad.y * r.pos.y})
	return topLeft, V.Sum(
		topLeft,
		Vector2d{
			r.size.x - 1,
			r.size.y - 1,
		},
	)
}

func (w *World) Draw(c *Canvas) {
	for _, r := range w.rooms {
		drawPos := V.Sum(w.position, Vector2d{r.pos.x * (r.size.x + 1), r.pos.y * (r.size.y + 1)})
		r.Draw(c, drawPos)
	}
}
