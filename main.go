package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	cellSize     = 10
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "cells")
	rl.SetTargetFPS(60)

	gridWidth := screenWidth / cellSize
	gridHeight := screenHeight / cellSize

	grid := make([][]bool, gridWidth)
	for i := range grid {
		grid[i] = make([]bool, gridHeight)
	}

	colors := make([][]rl.Color, gridWidth)
	for i := range colors {
		colors[i] = make([]rl.Color, gridHeight)
		for j := range colors[i] {
			colors[i][j] = rl.White
		}
	}

	rainbowTriggered := false

	for !rl.WindowShouldClose() {

		if rl.IsActivated(rl.KeyR) {
			rainbowTriggered = !rainbowTriggered

			if rainbowTriggered {
				for i := range grid {
					for j := range grid[i] {
						if grid[i][j] {
							colors[i][j] = rl.color{
								R: unit8(rand.Int(256)),
								G: unit8(rand.Int(256)),
								B: unit8(rand.Int(256)),
								A: 255,
							}
						}
					}
				}
			} else {
				// if rainbowtrigger is off, all colors are set to white
				for i := range colors {
					for j := range colors[i] {
						colors[i][j] = rl.White
					}
				}
			}
		}

		//Check for mouse-click
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mouseX := rl.GetMouseX()
			MouseY := rl.GetMouseY()

			gridX := mouseX / cellSize
			gridY := mouseY / cellSize
		}

	}
}
