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

type pongBall struct {
	x int       // CentreX
	y int       // CentreY
	d direction // Direction of the ball
	s int       // Speed of the ball
	r int       // Radius of the ball
}

type hitBar struct {
	length     int  // Length of hitbar
	height     int  // Height of hitbar
	speed      int  // Speed of hitbar
	leftX      int  // left X cord of hitbar
	accelLeft  bool // Accelerate to left
	accelRight bool // Accelerate to Right
}

type pongGame struct {
	name     string   // Name of the Game
	score    int      // Game Score
	size     [2]int   // Game screen size
	gameOver bool     // Game over flag
	ball     pongBall // Properties of ball
	bar      hitBar   // Properties of hitbar
}

func resetGame(game *pongGame) {
	rand.Seed(time.Now().Unix())

	// Init game specific values
	game.gameOver = false
	game.name = gameName
	game.score = 0
	game.size = [2]int{width, height}

	// Init ball specific values
	game.ball.x = rand.Intn(game.size[0]-200) + 100
	game.ball.y = 150
	game.ball.d = upLeft
	game.ball.s = ballSpeed
	game.ball.r = ballRadius

	// Init hit bar specific values
	game.bar.accelLeft = false
	game.bar.accelRight = false
	game.bar.height = hitBarHeight
	game.bar.length = hitBarLength
	game.bar.leftX = game.size[0]/2 - game.bar.length/2
	game.bar.speed = hitBarSpeed
}

func drawBall(ball *pongBall) {
	rl.DrawCircle(int32(ball.x), int32(ball.y), float32(ball.r), rl.Red)
}

func drawHitBar(bar *hitBar) {
	rl.DrawRectangle(int32(bar.leftX), int32(height-bar.height), int32(bar.length), int32(bar.height), rl.Blue)
}

func litsenKeyboardEvents(game *pongGame) {
	bar := &game.bar
	if rl.IsKeyDown(enterKey) {
		if game.gameOver { // Reset game when pressing enter key
			resetGame(game)
		}
	}
	if rl.IsKeyDown(leftArrowKey) {
		bar.accelLeft = true
	}
	if rl.IsKeyDown(rightArrowKey) {
		bar.accelRight = true
	}
	if rl.IsKeyUp(leftArrowKey) {
		bar.accelLeft = false
	}
	if rl.IsKeyUp(rightArrowKey) {
		bar.accelRight = false
	}
}

func moveBall(game *pongGame) {
	ball := &game.ball
	bar := &game.bar
	if ball.d == upLeft { // When ball's direction is UP_LEFT
		if ((ball.x - ball.s) > ball.r) && ((ball.y - ball.s) > ball.r) {
			ball.x -= ball.s
			ball.y -= ball.s
		} else if (ball.y - ball.s) > ball.r {
			ball.d = upRight
		} else if (ball.x - ball.s) > ball.r {
			ball.d = downLeft
		} else {
			ball.d = downRight
		}
	} else if ball.d == upRight { // When ball's direction is UP_RIGHT
		if ((ball.x + ball.s) < (game.size[0] - ball.r)) && ((ball.y - ball.s) > ball.r) {
			ball.x += ball.s
			ball.y -= ball.s
		} else if (ball.x + ball.s) < (game.size[0] - ball.r) {
			ball.d = downRight
		} else if (ball.x - ball.s) > ball.r {
			ball.d = upLeft
		} else {
			ball.d = downLeft
		}
	} else if ball.d == downLeft { // When ball's direction is DOWN_LEFT
		if (ball.x - ball.s) > ball.r {
			// If ball is in between hit bar cordinates
			if (ball.x+ball.r) >= bar.leftX && (ball.x-ball.r) <= (bar.leftX+bar.length) {
				if (ball.y + ball.s) < (game.size[1] - (ball.r + bar.height)) {
					ball.x -= ball.s
					ball.y += ball.s
				} else { // Ball touches hit bar
					game.score++
					ball.d = upLeft
				}
			} else if ball.y+ball.s >= game.size[1] { // Ball touches floor; Game over
				game.gameOver = true
			} else {
				ball.x -= ball.s
				ball.y += ball.s
			}
		} else if (ball.y + ball.s) < (game.size[1] - ball.r) {
			ball.d = downRight
		} else {
			ball.d = upRight
		}
	} else if ball.d == downRight { // When ball's direction is DOWN_RIGHT
		if (ball.x + ball.s) < (game.size[0] - ball.r) {
			// If ball is in between hit bar cordinates
			if (ball.x+ball.r) >= bar.leftX && (ball.x-ball.r) <= (bar.leftX+bar.length) {
				if (ball.y + ball.s) < (game.size[1] - (ball.r + bar.height)) {
					ball.x += ball.s
					ball.y += ball.s
				} else { // Ball touches hit bar
					game.score++
					ball.d = upRight
				}
			} else if ball.y+ball.s >= game.size[1] { // Ball touches floor; Game over
				game.gameOver = true
			} else {
				ball.x += ball.s
				ball.y += ball.s
			}
		} else if (ball.x - ball.s) > ball.r {
			ball.d = downLeft
		} else {
			ball.d = upLeft
		}
	}
}

func moveHitBar(game *pongGame) {
	bar := &game.bar
	if bar.accelLeft { // Move hit bar to left
		if (bar.leftX - hitBarSpeed) >= 0 {
			bar.leftX -= hitBarSpeed
		}
	} else if bar.accelRight { // Move hit bar to right
		if (bar.leftX + hitBarSpeed + bar.length) <= game.size[0] {
			bar.leftX += hitBarSpeed
		}
	}
}

func main() {
	game := new(pongGame) // Original Game object
	resetGame(game)       // init game object for the first time
	rl.InitWindow(int32(game.size[0]), int32(game.size[1]), game.name)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if !game.gameOver {
			moveBall(game)   // Calculate next position of ball
			moveHitBar(game) // Calculate next position of hit bar
		}
		litsenKeyboardEvents(game)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawBall(&game.ball)  // Draw ball
		drawHitBar(&game.bar) // Draw bar
		rl.DrawText(strconv.Itoa(game.score), 840, 70, 80, rl.White)
		if game.gameOver {
			rl.DrawText("Game Over!", 220, 200, 110, rl.White)
			rl.DrawText("Your Score : "+strconv.Itoa(game.score), 350, 390, 40, rl.Gray)
			rl.DrawText("Press enter key to continue...", 650, 520, 18, rl.LightGray)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
