package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Engine struct {
	canvas      Canvas
	roomsCanvas Canvas
	ui          Canvas
	player      Player
	room        Room
}

func NewEngine() *Engine {
	return &Engine{
		canvas: Canvas{},
		ui:     Canvas{},
		player: Player{},
		room:   Room{},
	}
}

var firstInput bool = true
var uiWidth int = 25

// This is the main loop
func (e *Engine) MainLoop(channel <-chan []byte) {
	// We generate the room/s
	e.room = *NewRoom(1, 1, 13, 7, true, true, true, true, true)
	e.room.Draw(&e.roomsCanvas, 1, 1)
	// Generate the player
	e.player = Player{
		x: 2, y: 2, char: '⍤', health: 3,
	}
	// This is the "Ready to play screen"
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(terminal.pos(1, 1))
	fmt.Printf("Press any key...")

	for {
		r := e.ReadInput(channel)
		if firstInput {
			r = []byte(K_EMPTY)
			firstInput = false
		}

		if bytes.Equal(r, []byte(K_ESC)) {
			break
		}

		e.Update(r)
		e.Render()
	}
}

// This goroutine reads the input
func (e *Engine) InputLoop(buffer chan<- []byte) {
	for {
		r := ReadRaw()
		if len(r) > 0 {
			buffer <- r
		}
	}
}

func (e *Engine) ReadInput(channel <-chan []byte) []byte {
	// For now it will be a pure input based tick rate
	return <-channel

	// TODO: non-blocking input in the future
	// select {
	// case r := <-channel:
	// 	return r
	// default:
	// 	return []byte(K_EMPTY)
	// }
}

func (e *Engine) Update(r []byte) {
	key := string(r)
	if key == K_ARROW_UP || key == KEY_k {
		e.player.y--
	}
	if key == K_ARROW_DOWN || key == KEY_j {
		e.player.y++
	}
	if key == K_ARROW_LEFT || key == KEY_h {
		e.player.x--
	}
	if key == K_ARROW_RIGHT || key == KEY_l {
		e.player.x++
	}

	e.player.y = Clamp(1, 45, e.player.y)
	e.player.x = Clamp(1, 153, e.player.x)

	// Some UI updates
	e.ui.ClearBuffer()
	// UI borders
	e.ui.HorizontalLine(terminal.width-uiWidth, 1, terminal.width-1, '═')
	e.ui.HorizontalLine(terminal.width-uiWidth, terminal.height, terminal.width-1, '═')
	e.ui.VerticalLine(terminal.width-uiWidth, 1, terminal.height, '║')
	e.ui.VerticalLine(terminal.width, 1, terminal.height, '║')
	e.ui.buffer.WriteString(terminal.pos(terminal.width-uiWidth, 1) + "╔")
	e.ui.buffer.WriteString(terminal.pos(terminal.width, 1) + "╗")
	e.ui.buffer.WriteString(terminal.pos(terminal.width-uiWidth, terminal.height) + "╚")
	e.ui.buffer.WriteString(terminal.pos(terminal.width, terminal.height) + "╝")
	// UI info
	e.ui.buffer.WriteString(terminal.pos(terminal.width-23, 2))
	e.ui.buffer.WriteString("Health: " + strings.Repeat("♥", e.player.health))
	e.ui.buffer.WriteString(terminal.pos(terminal.width-23, 3))
	e.ui.buffer.WriteString(fmt.Sprintf("Size: %dx%d", terminal.width, terminal.height))
	e.ui.buffer.WriteString(terminal.pos(terminal.width-23, 4))
	e.ui.buffer.WriteString(fmt.Sprintf("UI: %dx%d", 25, terminal.height))
	e.ui.buffer.WriteString(terminal.pos(terminal.width-23, 5))
	e.ui.buffer.WriteString(fmt.Sprintf("World: %dx%d", terminal.width-25, terminal.height))
	e.ui.buffer.WriteString(terminal.pos(terminal.width-23, 6))
	e.ui.buffer.WriteString(fmt.Sprintf("Player: %dx%d", e.player.x, e.player.y))
}

func (e *Engine) Render() {
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(e.canvas.ToString())
	fmt.Print(e.roomsCanvas.ToString())
	fmt.Print(e.ui.ToString())
	e.player.Render()
}
