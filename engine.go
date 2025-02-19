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
	world       World
	logs        []string
}

func NewEngine() *Engine {
	return &Engine{
		canvas:      Canvas{},
		roomsCanvas: Canvas{},
		uiBack:      Canvas{},
		uiFront:     Canvas{},
		player:      Player{},
		logs:        []string{},
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
	e.roomsCanvas.DrawSquare(Vector2d{1, 1}, Vector2d{term.width - uiWidth - 2, term.height})
	// We generate the world
	e.world = *NewWorld()
	e.world.Generate()
	e.world.Draw(&e.roomsCanvas)
	// Generate the player
	e.player = Player{
		worldPosition: V.Zero,
		roomPosition:  V.Identity,
		currentRoom:   V.Zero,
		char:          '⍤',
		health:        3,
	}
	startingRoom, _ := e.world.GetRoom(e.player.currentRoom)
	e.player.roomPosition = startingRoom.GetSpawnLocation()
	// UI Static stuff
	e.uiBack.ClearBuffer()
	e.uiBack.DrawSquare(Vector2d{terminal.width - uiWidth, 1}, Vector2d{uiWidth, term.height})
	// This is the "Ready to play screen"
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(terminal.pos(1, 1))
	fmt.Printf("Press any key...")
}

// Update logic
func (e *Engine) Update(r []byte) {
	key := string(r)

	e.UpdatePlayer(key)
	e.UpdateUi()
}

func (e *Engine) UpdatePlayer(key string) {
	// Player movement
	direction := V.Zero
	if key == K_ARROW_UP || key == KEY_k {
		direction = V.Up
	}
	if key == K_ARROW_DOWN || key == KEY_j {
		direction = V.Down
	}
	if key == K_ARROW_LEFT || key == KEY_h {
		direction = V.Left
	}
	if key == K_ARROW_RIGHT || key == KEY_l {
		direction = V.Right
	}
	// Updating the room position
	nextPosition := V.Sum(e.player.roomPosition, direction)
	r, _ := e.world.GetRoom(e.player.currentRoom)

	if r.IsDoor(nextPosition) {
		doorDirection := r.GetDoorDirection(nextPosition)
		e.player.currentRoom = V.Sum(e.player.currentRoom, doorDirection)
		e.player.roomPosition = r.GetNextRoomEnterPosition(nextPosition)

		// Reveal next room
		nextRoom, _ := e.world.GetRoom(e.player.currentRoom)
		if !nextRoom.visible {
			nextRoom.visible = true
			e.roomsCanvas.ClearBuffer()
			e.world.Draw(&e.roomsCanvas)
		}

		return
	}
	if r.IsValidPosition(nextPosition) {
		e.player.roomPosition = nextPosition
	}
}

func (e *Engine) UpdateUi() {
	// Some UI updates
	e.uiFront.ClearBuffer()
	// UI info
	uiStrings := []string{
		"Health: " + strings.Repeat("♥", e.player.health),
		"------- DEBUG: -------",
		fmt.Sprintf("Terminal:  %3dx%d", terminal.width, terminal.height),
		fmt.Sprintf("UI:        %3dx%d", uiWidth, terminal.height),
		fmt.Sprintf("World:     %3dx%d", terminal.width-uiWidth, terminal.height),
		fmt.Sprintf("Pos. Room: %3dx%d", e.player.roomPosition.x, e.player.roomPosition.y),
		fmt.Sprintf("Cur. Room: %3dx%d", e.player.currentRoom.x, e.player.currentRoom.y),
	}
	for i, s := range uiStrings {
		e.uiFront.AddString(terminal.pos(terminal.width-23, 2+i))
		e.uiFront.AddString(s)
	}
}

// Rendering logic
func (e *Engine) Render() {
	fmt.Print(CLEAR_SCREEN)

	fmt.Print(e.canvas.ToString())
	fmt.Print(e.roomsCanvas.ToString())
	fmt.Print(e.uiBack.ToString())
	fmt.Print(e.uiFront.ToString())

	renderPos := V.Sum(e.player.roomPosition, e.world.position)
	renderPos = V.Sum(renderPos, e.world.GetRoomWoldPosition(e.player.currentRoom))
	e.player.RenderAt(renderPos)
}
