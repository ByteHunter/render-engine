package main

type Room struct {
	x, y       int // Coordinates of the room, not in the canvas
	w, h       int
	visible    bool
	a, b, c, d bool // These are the available doors, clockwise from the top (a)
}

func NewRoom(x, y, w, h int, visible, a, b, c, d bool) *Room {
	return &Room{
		x:       x,
		y:       y,
		w:       w,
		h:       h,
		visible: visible,
		a:       a,
		b:       b,
		c:       c,
		d:       d,
	}
}

// Draw a room at a given position
func (r *Room) Draw(c *Canvas, x, y int) {
	c.SetBrush(BLOCK_2593)
	c.DrawBox(x, y, r.w, r.h)
	if r.visible {
		c.SetBrush(' ')
	} else {
		c.SetBrush(BLOCK_2571)
	}
	c.DrawBox(x+1, y+1, r.w-2, r.h-2)
	if r.a {
		c.buffer.WriteString(terminal.pos(x+r.w/2, y) + " ")
	}
	if r.b {
		c.buffer.WriteString(terminal.pos(x+r.w-1, y+r.h/2) + " ")
	}
	if r.c {
		c.buffer.WriteString(terminal.pos(x+r.w/2, y+r.h-1) + " ")
	}
	if r.d {
		c.buffer.WriteString(terminal.pos(x, y+r.h/2) + " ")
	}
}
