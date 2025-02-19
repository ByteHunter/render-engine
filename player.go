package main

import "fmt"

type Player struct {
	position Vector2d
	char     rune
	health   int
}

func (p *Player) Render() {
    fmt.Print(terminal.pos2d(p.position))
	fmt.Printf("%c", p.char)
}
