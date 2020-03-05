package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const gameName string = "Nithins's Pong Game!"

const (
	width        int32 = 1000
	height       int32 = 600
	ballSpeed    int32 = 8
	ballRadius   int32 = 25
	hitBarSpeed  int   = 14
	hitBarLength int32 = 100
	hitBarHeight int32 = 25
)

var (
	screenSize    [2]int32
	hitBarLeft    int32
	ballCentreY   int32
	ballCentreX   int32
	ballDirection string
)

func main() {
	initGame()

	rl.InitWindow(screenSize[0], screenSize[1], gameName)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		moveBall()
		rl.ClearBackground(rl.Black)
		drawBall()
		drawHitBar()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func initGame() {
	rand.Seed(time.Now().Unix())
	screenSize = [2]int32{width, height}
	hitBarLeft = screenSize[0]/2 - int32(hitBarLength)/2
	ballCentreY = 150
	ballCentreX = rand.Int31n(screenSize[0]-200) + 100
	ballDirection = "UP_LEFT"
}

func drawBall() {
	rl.DrawCircle(ballCentreX, ballCentreY, float32(ballRadius), rl.Red)
}

func drawHitBar() {
	rl.DrawRectangle(hitBarLeft, (screenSize[1] - hitBarHeight), hitBarLength, hitBarHeight, rl.Blue)
}

func moveBall() {
	if ballDirection == "UP_LEFT" {
		if ((ballCentreX - ballSpeed) > ballRadius) && ((ballCentreY - ballSpeed) > ballRadius) {
			ballCentreX -= ballSpeed
			ballCentreY -= ballSpeed
		} else if (ballCentreY - ballSpeed) > ballRadius {
			ballDirection = "UP_RIGHT"
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = "DOWN_LEFT"
		} else {
			ballDirection = "DOWN_RIGHT"
		}
	} else if ballDirection == "DOWN_LEFT" {
		if ((ballCentreX - ballSpeed) > ballRadius) && ((ballCentreY + ballSpeed) < (screenSize[1] - ballRadius)) {
			ballCentreX -= ballSpeed
			ballCentreY += ballSpeed
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = "UP_LEFT"
		} else if (ballCentreY + ballSpeed) < (screenSize[1] - ballRadius) {
			ballDirection = "DOWN_RIGHT"
		} else {
			ballDirection = "UP_RIGHT"
		}
	} else if ballDirection == "DOWN_RIGHT" {
		if ((ballCentreX + ballSpeed) < (screenSize[0] - ballRadius)) && ((ballCentreY + ballSpeed) < (screenSize[1] - ballRadius)) {
			ballCentreX += ballSpeed
			ballCentreY += ballSpeed
		} else if (ballCentreX + ballSpeed) < (screenSize[0] - ballRadius) {
			ballDirection = "UP_RIGHT"
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = "DOWN_LEFT"
		} else {
			ballDirection = "UP_LEFT"
		}
	} else if ballDirection == "UP_RIGHT" {
		if ((ballCentreX + ballSpeed) < (screenSize[0] - ballRadius)) && ((ballCentreY - ballSpeed) > ballRadius) {
			ballCentreX += ballSpeed
			ballCentreY -= ballSpeed
		} else if (ballCentreX + ballSpeed) < (screenSize[0] - ballRadius) {
			ballDirection = "DOWN_RIGHT"
		} else if (ballCentreX - ballSpeed) > ballRadius {
			ballDirection = "UP_LEFT"
		} else {
			ballDirection = "DOWN_LEFT"
		}
	}
}
