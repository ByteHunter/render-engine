package main

import (
	"fmt"
)

var terminal Terminal

func main() {
	terminal = *NewTerminal()
	terminal.init()
	terminal.LineWrap(false)
	terminal.configure()
	defer terminal.restore()


    engine := NewEngine()

	fmt.Printf(CLEAR_SCREEN)
	fmt.Printf(HOME)

	channel := make(chan []byte)
	go engine.InputLoop(channel)
	engine.MainLoop(channel)

	terminal.LineWrap(true)
	fmt.Printf(RESET)
	fmt.Printf(terminal.pos(1, terminal.height))
}

