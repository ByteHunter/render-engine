package main

import "strings"

type Canvas struct {
	buffer     strings.Builder
	brush      rune
	needRedraw bool
}

func (c *Canvas) ClearBuffer() {
	c.buffer.Reset()
}

func (c *Canvas) AddString(s string) {
	c.buffer.WriteString(s)
}

func (c *Canvas) AddRune(r rune) {
	c.buffer.WriteRune(r)
}

func (c *Canvas) ToString() string {
	return c.buffer.String()
}

func (c *Canvas) SetBrush(b rune) {
	c.brush = b
}

func (c *Canvas) DrawBox(position, size Vector2d) {
	c.buffer.WriteString(terminal.background(0, 0, 0))
	c.buffer.WriteString(terminal.foreground(255, 255, 255))
	c.buffer.WriteString(terminal.pos2d(position))
	for range size.y {
		c.buffer.WriteString(CURSOR_SAVE)
		c.buffer.WriteString(strings.Repeat(string(c.brush), size.x))
		c.buffer.WriteString(CURSOR_LOAD)
		c.buffer.WriteString(terminal.cursorDown(1))
	}
}

func (c *Canvas) VerticalLine(position Vector2d, h int, r rune) {
	c.buffer.WriteString(terminal.pos2d(position))
	for range h {
		c.buffer.WriteString(CURSOR_SAVE)
		c.buffer.WriteRune(r)
		c.buffer.WriteString(CURSOR_LOAD)
		c.buffer.WriteString(terminal.cursorDown(1))
	}
}

func (c *Canvas) HorizontalLine(position Vector2d, w int, r rune) {
	c.buffer.WriteString(terminal.pos2d(position))
	for range w {
		c.buffer.WriteRune(r)
	}
}

func (c *Canvas) DrawSquare(position, size Vector2d) {
	var tl, tr, bl, br rune = '╔', '╗', '╚', '╝'
	var hr, vr rune = '═', '║'
	// Draw the lines
	c.HorizontalLine(position, size.x, hr)
	c.HorizontalLine(Vector2d{position.x, position.y + size.y}, size.x, hr)
	c.VerticalLine(position, size.y, vr)
	c.VerticalLine(Vector2d{position.x + size.x, position.y}, size.y, vr)
	// Draw the corners
	c.buffer.WriteString(terminal.pos2d(position))
	c.buffer.WriteRune(tl)
	c.buffer.WriteString(terminal.pos(position.x+size.x, position.y))
	c.buffer.WriteRune(tr)
	c.buffer.WriteString(terminal.pos(position.x, position.y+size.y))
	c.buffer.WriteRune(bl)
	c.buffer.WriteString(terminal.pos(position.x+size.x, position.y+size.y))
	c.buffer.WriteRune(br)
}
