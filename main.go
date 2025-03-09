package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth      = 800
	screenHeight     = 600
	cellSize         = 10
	accelerationRate = 0.2
)

type cellState struct {
	active  bool
	color   rl.Color
	falling bool
	yPos    float64
	ySpeed  float64
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "cells")
	rl.SetTargetFPS(60)

	gridWidth := screenWidth / cellSize
	gridHeight := screenHeight / cellSize

	cells := make([][]cellState, gridWidth)
	for i := range cells {
		cells[i] = make([]cellState, gridHeight)
		for j := range cells[i] {
			cells[i][j] = cellState{
				active:  false,
				color:   rl.White,
				falling: false,
				yPos:    float64(j),
				ySpeed:  0,
			}
		}
	}

	rainbowTriggered := false

	for !rl.WindowShouldClose() {

		//Get DeltaTime for smoother physics
		deltaTime := rl.GetFrameTime()

		if rl.IsKeyPressed(rl.KeyR) {
			rainbowTriggered = !rainbowTriggered

			if rainbowTriggered {
				for i := range cells {
					for j := range cells[i] {
						if cells[i][j].active {
							cells[i][j].color = rl.Color{
								R: uint8(rand.Intn(256)),
								G: uint8(rand.Intn(256)),
								B: uint8(rand.Intn(256)),
								A: 255,
							}
						}
					}
				}
			} else {
				// if rainbowtrigger is off, all colors are set to white
				for i := range cells {
					for j := range cells[i] {
						cells[i][j].color = rl.White
					}
				}
			}
		}

		//Check for mouse-click
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			gridX := mouseX / cellSize
			gridY := mouseY / cellSize

			//change state on/off
			if gridX >= 0 && gridX < int32(gridWidth) && gridY >= 0 && gridY < int32(gridHeight) {
				cells[gridX][gridY].active = !cells[gridX][gridY].active

				if cells[gridX][gridY].active {
					cells[gridX][gridY].falling = true
					cells[gridX][gridY].ySpeed = 0
				}

				//If cell become active and rainbowTrigger is on, give it a random color
				if cells[gridX][gridY].active && rainbowTriggered {
					cells[gridX][gridY].color = rl.Color{
						R: uint8(rand.Intn(256)),
						G: uint8(rand.Intn(256)),
						B: uint8(rand.Intn(256)),
						A: 255,
					}
				}
			}
		}

		for x := 0; x < gridWidth; x++ {
			for y := gridHeight - 1; y >= 0; y-- {
				if cells[x][y].active && cells[x][y].falling {
					//update the speed and position acceleration
					cells[x][y].ySpeed += accelerationRate * float64(deltaTime) * 60
					cells[x][y].yPos += cells[x][y].ySpeed

					//check if at bottom or for a cell
					gridYPos := int(cells[x][y].yPos)

					if gridYPos >= gridHeight-1 || (gridYPos+1 < gridHeight && gridYPos != y && cells[x][gridYPos+1].active) {
						if gridYPos >= gridHeight-1 {
							gridYPos = gridHeight - 1
						}

						if gridYPos != y {
							cells[x][gridYPos] = cells[x][y]
							cells[x][gridYPos].falling = false
							cells[x][gridYPos].yPos = float64(gridYPos)
							cells[x][gridYPos].ySpeed = 0

							//remove old position
							cells[x][y] = cellState{
								active:  false,
								color:   rl.White,
								falling: false,
								yPos:    float64(y),
								ySpeed:  0,
							}
						} else {
							// cell stop falling
							cells[x][y].falling = false
						}
					}
				}
			}
		}

		//begin draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		//draw the grid
		for x := 0; x < screenWidth; x += cellSize {
			rl.DrawLine(int32(x), 0, int32(x), int32(screenHeight), rl.DarkGray)
		}

		for y := 0; y < screenHeight; y += cellSize {
			rl.DrawLine(0, int32(y), int32(screenWidth), int32(y), rl.DarkGray)
		}

		//create active cells
		for i := range cells {
			for j := range cells[i] {
				if cells[i][j].active {
					var yPosition int32
					if cells[i][j].falling {
						yPosition = int32(cells[i][j].yPos * float64(cellSize))
					} else {
						yPosition = int32(j * cellSize)
					}

					rl.DrawRectangle(
						int32(i*cellSize),
						yPosition,
						int32(cellSize),
						int32(cellSize),
						cells[i][j].color,
					)
				}
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
