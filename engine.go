package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Engine struct {
	canvas      Canvas
	roomsCanvas Canvas
	uiBack      Canvas
	uiFront     Canvas
	player      Player
	room        Room
}

func NewEngine() *Engine {
	return &Engine{
		canvas:      Canvas{},
		roomsCanvas: Canvas{},
		uiBack:      Canvas{},
		uiFront:     Canvas{},
		player:      Player{},
		room:        Room{},
	}
}

var firstInput bool = true
var uiWidth int = 25

// This is the main loop
func (e *Engine) MainLoop(channel <-chan []byte) {
	e.Init()

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

// Initialize logic
func (e *Engine) Init() {
	// Room canvas
	e.roomsCanvas.DrawSquare(1, 1, terminal.width-uiWidth-2, terminal.height)
	// We generate the room/s
	e.room = *NewRoom(5, 5, 13, 7, true, true, true, true, true)
	e.room.Draw(&e.roomsCanvas, 5, 3)
	e.room.visible = false
	for j := range 5 {
		for i := range 9 {
			if j == 0 && i == 0 {
				continue
			}
			e.room.Draw(&e.roomsCanvas, 5+i*14, 3+j*8)
		}
	}
	// Generate the player
	e.player = Player{
		x: 2, y: 2, char: '⍤', health: 3,
	}
	// UI Static stuff
	e.uiBack.ClearBuffer()
	e.uiBack.DrawSquare(terminal.width-uiWidth, 1, uiWidth, terminal.height)
	// This is the "Ready to play screen"
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(terminal.pos(1, 1))
	fmt.Printf("Press any key...")
}

// Update logic
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
	e.uiFront.ClearBuffer()
	// UI info
	e.uiFront.buffer.WriteString(terminal.pos(terminal.width-23, 2))
	e.uiFront.buffer.WriteString("Health: " + strings.Repeat("♥", e.player.health))
	// e.uiFront.buffer.WriteString(terminal.pos(terminal.width-23, 3))
	// e.uiFront.buffer.WriteString(fmt.Sprintf("Size: %dx%d", terminal.width, terminal.height))
	// e.uiFront.buffer.WriteString(terminal.pos(terminal.width-23, 4))
	// e.uiFront.buffer.WriteString(fmt.Sprintf("UI: %dx%d", 25, terminal.height))
	// e.uiFront.buffer.WriteString(terminal.pos(terminal.width-23, 5))
	// e.uiFront.buffer.WriteString(fmt.Sprintf("World: %dx%d", terminal.width-25, terminal.height))
	// e.uiFront.buffer.WriteString(terminal.pos(terminal.width-23, 6))
	// e.uiFront.buffer.WriteString(fmt.Sprintf("Player: %dx%d", e.player.x, e.player.y))
}

// Rendering logic
func (e *Engine) Render() {
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(e.canvas.ToString())
	fmt.Print(e.roomsCanvas.ToString())
	fmt.Print(e.uiBack.ToString())
	fmt.Print(e.uiFront.ToString())
	e.player.Render()
}
