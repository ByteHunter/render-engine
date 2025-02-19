package main

import "fmt"

type Player struct {
	position     Vector2d // This should be the world position instead
	roomPosition Vector2d
	char         rune
	health       int
}

func (p *Player) Render() {
	fmt.Print(terminal.pos2d(p.position))
	fmt.Printf("%c", p.char)
}

func (p *Player) RenderAt(pos Vector2d) {
	fmt.Print(terminal.pos2d(pos))
	fmt.Printf("%c", p.char)
}
