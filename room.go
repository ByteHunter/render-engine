package main

type Room struct {
	id         int
	pos        Vector2d // This will be the position
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
    return true
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
