package main

import (
	"fmt"
	"time"
	"golang.org/x/term"
	"strings"
)

func draw_line(char string, width int) {
	fmt.Print(strings.Repeat(char, width))
	fmt.Print("\n")
}

func draw_fill(char string, width int, height int) {
	fmt.Print(strings.Repeat(char, width * height), "\n")
}

func move_cursor(line int, col int) {
	fmt.Printf("\033[%d;%dH", line, col)
	time.Sleep(1 * time.Millisecond)
}

func player_position() {
	
}

func game_loop(width int, height int) {
	// game loop
	for range time.Tick(60 * time.Millisecond) {
		fmt.Print("\033[2J") // clear screen
		move_cursor(400,0) // move cursor home
		draw_fill("%", width, height)

		move_cursor(15, 13)
	}
}

func main() {
	// get terminal size
	width, height, err := term.GetSize(0)
	if err != nil {
		return
	}
	fmt.Print(width, height)
	
	// enter alternate buffer
	fmt.Print("\033[?1049h")
	
	game_loop(width, height)
}
