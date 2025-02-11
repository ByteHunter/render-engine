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

func (c *Canvas) ToString() string {
	return c.buffer.String()
}

func (c *Canvas) SetBrush(b rune) {
	c.brush = b
}

func (c *Canvas) DrawBox(x, y, w, h int) {
	c.buffer.WriteString(terminal.background(0, 0, 0))
	c.buffer.WriteString(terminal.foreground(255, 255, 255))
	c.buffer.WriteString(terminal.pos(x, y))
	for range h {
		c.buffer.WriteString(CURSOR_SAVE)
		c.buffer.WriteString(strings.Repeat(string(c.brush), w))
		c.buffer.WriteString(CURSOR_LOAD)
		c.buffer.WriteString(terminal.cursorDown(1))
	}
}

func (c *Canvas) VerticalLine(x, y, h int, r rune) {
	c.buffer.WriteString(terminal.pos(x, y))
	for range h {
		c.buffer.WriteString(CURSOR_SAVE)
		c.buffer.WriteRune(r)
		c.buffer.WriteString(CURSOR_LOAD)
		c.buffer.WriteString(terminal.cursorDown(1))
	}
}

func (c *Canvas) HorizontalLine(x, y, w int, r rune) {
	c.buffer.WriteString(terminal.pos(x, y))
	for range w {
		c.buffer.WriteRune(r)
	}
}

func (c *Canvas) DrawSquare(x, y, w, h int) {
	var tl, tr, bl, br rune = '╔', '╗', '╚', '╝'
	var hr, vr rune = '═', '║'
    // Draw the lines
	c.HorizontalLine(x, y, w, hr)
	c.HorizontalLine(x, y+h, w, hr)
	c.VerticalLine(x, y, h, vr)
	c.VerticalLine(x+w, y, h, vr)
    // Draw the corners
	c.buffer.WriteString(terminal.pos(x, y))
	c.buffer.WriteRune(tl)
	c.buffer.WriteString(terminal.pos(x+w, y))
	c.buffer.WriteRune(tr)
	c.buffer.WriteString(terminal.pos(x, y+h))
	c.buffer.WriteRune(bl)
	c.buffer.WriteString(terminal.pos(x+w, y+h))
	c.buffer.WriteRune(br)
}
