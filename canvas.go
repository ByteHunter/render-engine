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
