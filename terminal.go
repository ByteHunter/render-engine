package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var TCONF_RAW = []string{
	"-F", "/dev/tty", "-icanon", "-echo",
	"min", "0", "time", "0",
	"-isig", "-ixon",
}

var TCONF_NORMAL = []string{
	"-F", "/dev/tty", "icanon", "echo",
}

type Terminal struct {
	width, height int
	size          Vector2d
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

// GLOBAL METHODS

func CommandOutput(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func GetTerminalSize() (int, int) {
	output := CommandOutput("stty", "size")
	output = strings.TrimSpace(output)
	size := strings.Split(output, " ")
	rows, _ := strconv.Atoi(size[0])
	cols, _ := strconv.Atoi(size[1])

	return cols, rows
}

func ReadRaw() []byte {
	var b []byte = make([]byte, 4096)
	os.Stdin.Sync()
	n, err := os.Stdin.Read(b)
	if err != nil {
		return []byte{}
	}
	return b[:n]
}

// SETTINGS

func (t *Terminal) CursorVisibility(b bool) {
	if b {
		fmt.Print(CURSOR_SHOW)
	} else {
		fmt.Print(CURSOR_HIDE)
	}
}

func (t *Terminal) LineWrap(b bool) {
	if b {
		fmt.Print(LINE_WRAP)
	} else {
		fmt.Print(NO_LINE_WRAP)
	}
}

// CONFIGURE

func (t *Terminal) init() {
	width, height := GetTerminalSize()
	t.width, t.height = width, height
	t.size = Vector2d{width, height}
}

func (t *Terminal) configure() {
	exec.Command("stty", TCONF_RAW...).Run()
	t.CursorVisibility(false)
}

func (t *Terminal) restore() {
	exec.Command("stty", TCONF_NORMAL...).Run()
	t.CursorVisibility(true)
}

// PRIMITIVES

func (t *Terminal) color(r, g, b int) string {
	return NUMS[r] + SEP + NUMS[g] + SEP + NUMS[b]
}

func (t *Terminal) foreground(r, g, b int) string {
	return FOREGROUND + t.color(r, g, b) + SEP + "m"
}

func (t *Terminal) background(r, g, b int) string {
	return BACKGROUND + t.color(r, g, b) + SEP + "m"
}

func (t *Terminal) reset() string {
	return RESET
}

func (t *Terminal) pos(x, y int) string {
	return CSI + NUMS[y] + SEP + NUMS[x] + "H"
}

func (t *Terminal) pos2d(p Vector2d) string {
	return CSI + NUMS[p.y] + SEP + NUMS[p.x] + "H"
}

func (t *Terminal) cursorUp(n int) string {
	return CSI + NUMS[n] + "A"
}

func (t *Terminal) cursorDown(n int) string {
	return CSI + NUMS[n] + "B"
}

func (t *Terminal) cursorForward(n int) string {
	return CSI + NUMS[n] + "C"
}

func (t *Terminal) cursorBackward(n int) string {
	return CSI + NUMS[n] + "D"
}
