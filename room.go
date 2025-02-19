package main

type Room struct {
	id         int
	pos        Vector2d // This is the position on the world grid
	size       Vector2d
	visible    bool
	a, b, c, d bool // These are the available doors, clockwise from the top (a)
}

func NewRoom(pos, size Vector2d, visible, a, b, c, d bool) *Room {
	return &Room{
		pos:     pos,
		size:    size,
		visible: visible,
		a:       a,
		b:       b,
		c:       c,
		d:       d,
	}
}

func (r *Room) IsValidPosition(pos Vector2d) bool {
	if pos.x < 1 || pos.x > 28 || pos.y < 1 || pos.y > 8 {
		return false
	}
	return true
}

func (r *Room) GetSpawnLocation() Vector2d {
	return Vector2d{r.size.x / 2, r.size.y / 2}
}

func (r *Room) IsDoor(pos Vector2d) bool {
	if (V.Equal(pos, Vector2d{15, 0}) && r.a) {
		return true
	}
	if (V.Equal(pos, Vector2d{29, 5}) && r.b) {
		return true
	}
	if (V.Equal(pos, Vector2d{15, 9}) && r.c) {
		return true
	}
	if (V.Equal(pos, Vector2d{0, 5}) && r.d) {
		return true
	}

	return false
}

func (r *Room) GetDoorDirection(pos Vector2d) Vector2d {
	if (V.Equal(pos, Vector2d{15, 0})) {
		return V.Up
	}
	if (V.Equal(pos, Vector2d{29, 5})) {
		return V.Right
	}
	if (V.Equal(pos, Vector2d{15, 9})) {
		return V.Down
	}
	if (V.Equal(pos, Vector2d{0, 5})) {
		return V.Left
	}

	return V.Zero
}

func (r *Room) GetNextRoomEnterPosition(pos Vector2d) Vector2d {
	if (V.Equal(pos, Vector2d{15, 0})) {
		return Vector2d{15, 8}
	}
	if (V.Equal(pos, Vector2d{29, 5})) {
		return Vector2d{1, 5}
	}
	if (V.Equal(pos, Vector2d{15, 9})) {
		return Vector2d{15, 1}
	}
	if (V.Equal(pos, Vector2d{0, 5})) {
		return Vector2d{28, 5}
	}

	return V.Identity
}

// Draw a room at a given position in a Canvas
func (r *Room) Draw(c *Canvas, p Vector2d) {
	c.SetBrush(BLOCK_2593)
	c.DrawBox(p, r.size)
	if r.visible {
		c.SetBrush(' ')
	} else {
		c.SetBrush(BLOCK_2571)
	}
	c.DrawBox(Vector2d{p.x + 1, p.y + 1}, Vector2d{r.size.x - 2, r.size.y - 2})
	if r.a {
		c.AddString(terminal.pos(p.x+r.size.x/2, p.y) + " ")
	}
	if r.b {
		c.AddString(terminal.pos(p.x+r.size.x-1, p.y+r.size.y/2) + " ")
	}
	if r.c {
		c.AddString(terminal.pos(p.x+r.size.x/2, p.y+r.size.y-1) + " ")
	}
	if r.d {
		c.AddString(terminal.pos(p.x, p.y+r.size.y/2) + " ")
	}
}
