# Render Engine

Proof of concept for a 2D renderer in the Terminal.

The idea of this project is to create a renderer engine while building a simple game.

# How to build and run

The only requirement is to have Golang installed in your system.

To build the app use the command:

```sh
go build main.go canvas.go constants.go engine.go player.go room.go runes.go terminal.go utils.go world.go
```

If you want to simply run use the `go run` command:

```
go run main.go canvas.go constants.go engine.go player.go room.go runes.go terminal.go utils.go world.go
```

If you have the `make` installed you can use it to build and run with ease.

Build: `make build`

Run: `make run`