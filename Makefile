BINARY_NAME=render-engine

build:
	go build -o ${BINARY_NAME} main.go canvas.go constants.go engine.go player.go room.go runes.go terminal.go utils.go world.go

run:
	go run main.go canvas.go constants.go engine.go player.go room.go runes.go terminal.go utils.go world.go
