package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type direction int

const (
	upLeft direction = iota + 1
	upRight
	downLeft
	downRight
)

const (
	enterKey      = 257
	leftArrowKey  = 263
	rightArrowKey = 262
)

const (
	width        = 1000
	height       = 600
	ballSpeed    = 8
	ballRadius   = 25
	hitBarSpeed  = 14
	hitBarLength = 100
	hitBarHeight = 25
	gameName     = "Nithins's Pong Game!"
)

var (
	score           int
	gameOver        bool
	hitBarLeft      int
	screenSize      [2]int
	ballCentreY     int
	ballCentreX     int
	ballDirection   direction
	accelerateLeft  bool
	accelerateRight bool
)

func resetGame() {
	rand.Seed(time.Now().Unix())
	screenSize = [2]int{width, height}
	hitBarLeft = screenSize[0]/2 - hitBarLength/2
	ballCentreY = 150
	ballCentreX = rand.Intn(screenSize[0]-200) + 100
	ballDirection = upLeft
	accelerateLeft = false
	accelerateRight = false
	gameOver = false
	score = 0
}

func drawBall() {
	rl.DrawCircle(int32(ballCentreX), int32(ballCentreY), float32(ballRadius), rl.Red)
}

func drawHitBar() {
	rl.DrawRectangle(int32(hitBarLeft), int32(screenSize[1]-hitBarHeight), int32(hitBarLength), int32(hitBarHeight), rl.Blue)
}

func litsenKeyboardEvents() {
	if rl.IsKeyDown(enterKey) {
		if gameOver {
			resetGame()
		}
	}
	if rl.IsKeyDown(leftArrowKey) {
		accelerateLeft = true
	}
	if rl.IsKeyDown(rightArrowKey) {
		accelerateRight = true
	}
	if rl.IsKeyUp(leftArrowKey) {
		accelerateLeft = false
	}
	if rl.IsKeyUp(rightArrowKey) {
		accelerateRight = false
	}
}

func moveBall() {
	if ballDirection == upLeft {
		if ((ballCentreX - ballSpeed) > ballRadius) && ((ballCentreY - ballSpeed) > ballRadius) {
			ballCentreX -= ballSpeed
			ballCentreY -= ballSpeed
		} else if (ballCentreY - ballSpeed) > ballRadius {
			ballDirection = upRight
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = downLeft
		} else {
			ballDirection = downRight
		}
	} else if ballDirection == upRight {
		if ((ballCentreX + ballSpeed) < (screenSize[0] - ballRadius)) && ((ballCentreY - ballSpeed) > ballRadius) {
			ballCentreX += ballSpeed
			ballCentreY -= ballSpeed
		} else if (ballCentreX + ballSpeed) < (screenSize[0] - ballRadius) {
			ballDirection = downRight
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = upLeft
		} else {
			ballDirection = downLeft
		}
	} else if ballDirection == downLeft {
		if (ballCentreX - ballSpeed) > ballRadius {
			if (ballCentreX+ballRadius) >= hitBarLeft && (ballCentreX-ballRadius) <= (hitBarLeft+hitBarLength) {
				if (ballCentreY + ballSpeed) < (screenSize[1] - (ballRadius + hitBarHeight)) {
					ballCentreX -= ballSpeed
					ballCentreY += ballSpeed
				} else {
					score++
					ballDirection = upLeft
				}
			} else if ballCentreY+ballSpeed >= screenSize[1] {
				gameOver = true
			} else {
				ballCentreX -= ballSpeed
				ballCentreY += ballSpeed
			}
		} else if (ballCentreY + ballSpeed) < (screenSize[1] - ballRadius) {
			ballDirection = downRight
		} else {
			ballDirection = upRight
		}
	} else if ballDirection == downRight {
		if (ballCentreX + ballSpeed) < (screenSize[0] - ballRadius) {
			if (ballCentreX+ballRadius) >= hitBarLeft && (ballCentreX-ballRadius) <= (hitBarLeft+hitBarLength) {
				if (ballCentreY + ballSpeed) < (screenSize[1] - (ballRadius + hitBarHeight)) {
					ballCentreX += ballSpeed
					ballCentreY += ballSpeed
				} else {
					score++
					ballDirection = upRight
				}
			} else if ballCentreY+ballSpeed >= screenSize[1] {
				gameOver = true
			} else {
				ballCentreX += ballSpeed
				ballCentreY += ballSpeed
			}
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = downLeft
		} else {
			ballDirection = upLeft
		}
	}
}

func moveHitBar() {
	if accelerateLeft {
		if (hitBarLeft - hitBarSpeed) >= 0 {
			hitBarLeft -= hitBarSpeed
		}
	} else if accelerateRight {
		if (hitBarLeft + hitBarSpeed + hitBarLength) <= screenSize[0] {
			hitBarLeft += hitBarSpeed
		}
	}
}

func main() {
	resetGame()
	rl.InitWindow(int32(screenSize[0]), int32(screenSize[1]), gameName)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if !gameOver {
			moveBall()
			moveHitBar()
		}
		litsenKeyboardEvents()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawBall()
		drawHitBar()
		rl.DrawText(strconv.Itoa(score), 840, 70, 80, rl.White)
		if gameOver {
			rl.DrawText("Game Over!", 220, 200, 110, rl.White)
			rl.DrawText("Your Score : "+strconv.Itoa(score), 350, 390, 40, rl.Gray)
			rl.DrawText("Press enter key to continue...", 650, 520, 18, rl.LightGray)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
