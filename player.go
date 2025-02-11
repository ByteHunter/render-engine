package main

import "fmt"

type Player struct {
	x, y int
	char rune
    health int
}

func (p *Player) Render() {
	fmt.Print(terminal.pos(p.x, p.y))
	fmt.Printf("%c", p.char)
}
