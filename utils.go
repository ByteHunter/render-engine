package main

func Clamp(a, b, v int) int {
	if v < a {
		return a
	}
	if v > b {
		return b
	}
	return v
}
