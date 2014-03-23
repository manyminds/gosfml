/*
#############################################
#	GOSFML2
#	SFML Examples:	 Pong
#	Ported from C++ to Go
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	const (
		paddleSpeed = float32(400)
		ballSpeed   = float32(400)
	)

	var (
		gameWidth  uint        = 800
		gameHeight uint        = 600
		paddleSize sf.Vector2f = sf.Vector2f{25, 100}
		ballRadius float32     = 10
	)

	ticker := time.NewTicker(time.Second / 60)
	AITicker := time.NewTicker(time.Second / 10)
	rand.Seed(time.Now().UnixNano())

	renderWindow := sf.NewRenderWindow(sf.VideoMode{gameWidth, gameHeight, 32}, "Pong (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())

	// Load the sounds used in the game
	buffer, _ := sf.NewSoundBufferFromFile("resources/ball.wav")
	ballSound := sf.NewSound(buffer)

	// Create the left paddle
	leftPaddle, _ := sf.NewRectangleShape()
	leftPaddle.SetSize(sf.Vector2f{paddleSize.X - 3, paddleSize.Y - 3})
	leftPaddle.SetOutlineThickness(3)
	leftPaddle.SetOutlineColor(sf.ColorBlack())
	leftPaddle.SetFillColor(sf.Color{100, 100, 200, 255})
	leftPaddle.SetOrigin(sf.Vector2f{paddleSize.X / 2, paddleSize.Y / 2})

	// Create the right paddle
	rightPaddle, _ := sf.NewRectangleShape()
	rightPaddle.SetSize(sf.Vector2f{paddleSize.X - 3, paddleSize.Y - 3})
	rightPaddle.SetOutlineThickness(3)
	rightPaddle.SetOutlineColor(sf.ColorBlack())
	rightPaddle.SetFillColor(sf.Color{200, 100, 100, 255})
	rightPaddle.SetOrigin(sf.Vector2f{paddleSize.X / 2, paddleSize.Y / 2})

	// Create the ball
	ball, _ := sf.NewCircleShape()
	ball.SetRadius(ballRadius - 3)
	ball.SetOutlineThickness(3)
	ball.SetOutlineColor(sf.ColorBlack())
	ball.SetFillColor(sf.ColorWhite())
	ball.SetOrigin(sf.Vector2f{ballRadius / 2, ballRadius / 2})

	// Load the text font
	font, _ := sf.NewFontFromFile("resources/sansation.ttf")

	// Initialize the pause message
	pauseMessage, _ := sf.NewText(font)
	pauseMessage.SetCharacterSize(40)
	pauseMessage.SetPosition(sf.Vector2f{170, 150})
	pauseMessage.SetColor(sf.ColorWhite())
	pauseMessage.SetString("Welcome to SFML pong!\nPress space to start the game")

	var (
		rightPaddleSpeed float32 = 0
		ballAngle        float32 = 0
		isPlaying        bool    = false
	)

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyReleased:
					switch ev.Code {
					case sf.KeyEscape:
						renderWindow.Close()
					case sf.KeySpace:
						if !isPlaying {
							// (re)start the game
							isPlaying = true

							// reset position of the paddles and ball
							leftPaddle.SetPosition(sf.Vector2f{10 + paddleSize.X/2, float32(gameHeight) / 2})
							rightPaddle.SetPosition(sf.Vector2f{float32(gameWidth) - 10 - paddleSize.X/2, float32(gameHeight) / 2})
							ball.SetPosition(sf.Vector2f{float32(gameWidth) / 2, float32(gameHeight) / 2})

							// reset the ball angle
							for {
								// Make sure the ball initial angle is not too much vertical
								ballAngle = rand.Float32() * math.Pi * 2
								if math.Abs(math.Cos(float64(ballAngle))) > 0.7 {
									break
								}
							}
						}
					}
				case sf.EventClosed:
					renderWindow.Close()
				}
			}

			//playing
			if isPlaying {
				deltaTime := time.Second / 60

				// Move the player's paddle
				if sf.KeyboardIsKeyPressed(sf.KeyUp) && leftPaddle.GetPosition().Y-paddleSize.Y/2 > 5 {
					leftPaddle.Move(sf.Vector2f{0, -paddleSpeed * float32(deltaTime.Seconds())})
				}

				if sf.KeyboardIsKeyPressed(sf.KeyDown) && leftPaddle.GetPosition().Y+paddleSize.Y/2 < float32(gameHeight)-5 {
					leftPaddle.Move(sf.Vector2f{0, paddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the computer's paddle
				if (rightPaddleSpeed < 0 && rightPaddle.GetPosition().Y-paddleSize.Y/2 > 5) || (rightPaddleSpeed > 0 && rightPaddle.GetPosition().Y+paddleSize.Y/2 < float32(gameHeight)-5) {
					rightPaddle.Move(sf.Vector2f{0, rightPaddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the ball
				factor := ballSpeed * float32(deltaTime.Seconds())
				ball.Move(sf.Vector2f{float32(math.Cos(float64(ballAngle))) * factor, float32(math.Sin(float64(ballAngle))) * factor})

				// Check collisions between the ball and the screen
				if ball.GetPosition().X-ballRadius < 0 {
					isPlaying = false
					pauseMessage.SetString("You lost !\nPress space to restart or\nescape to exit")
				}

				if ball.GetPosition().X+ballRadius > float32(gameWidth) {
					isPlaying = false
					pauseMessage.SetString("You won !\nPress space to restart or\nescape to exit")
				}

				if ball.GetPosition().Y-ballRadius < 0 {
					ballAngle = -ballAngle
					ball.SetPosition(sf.Vector2f{ball.GetPosition().X, ballRadius + 0.1})
					ballSound.Play()
				}

				if ball.GetPosition().Y+ballRadius > float32(gameHeight) {
					ballAngle = -ballAngle
					ball.SetPosition(sf.Vector2f{ball.GetPosition().X, float32(gameHeight) - ballRadius - 0.1})
					ballSound.Play()
				}

				// Check the collisions between the ball and the paddles
				// Left Paddle
				if ball.GetPosition().X-ballRadius < leftPaddle.GetPosition().X+paddleSize.X/2 &&
					ball.GetPosition().X-ballRadius > leftPaddle.GetPosition().X &&
					ball.GetPosition().Y+ballRadius >= leftPaddle.GetPosition().Y-paddleSize.Y/2 &&
					ball.GetPosition().Y-ballRadius <= leftPaddle.GetPosition().Y+paddleSize.Y/2 {

					if ball.GetPosition().Y > leftPaddle.GetPosition().Y {
						ballAngle = math.Pi - ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						ballAngle = math.Pi - ballAngle - rand.Float32()*math.Pi*0.2
					}

					ball.SetPosition(sf.Vector2f{leftPaddle.GetPosition().X + ballRadius + paddleSize.X/2 + 0.1, ball.GetPosition().Y})
					ballSound.Play()
				}

				// Right Paddle
				if ball.GetPosition().X+ballRadius > rightPaddle.GetPosition().X-paddleSize.X/2 &&
					ball.GetPosition().X+ballRadius < rightPaddle.GetPosition().X &&
					ball.GetPosition().Y+ballRadius >= rightPaddle.GetPosition().Y-paddleSize.Y/2 &&
					ball.GetPosition().Y-ballRadius <= rightPaddle.GetPosition().Y+paddleSize.Y/2 {

					if ball.GetPosition().Y > rightPaddle.GetPosition().Y {
						ballAngle = math.Pi - ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						ballAngle = math.Pi - ballAngle - rand.Float32()*math.Pi*0.2
					}

					ball.SetPosition(sf.Vector2f{rightPaddle.GetPosition().X - ballRadius - paddleSize.X/2 - 0.1, ball.GetPosition().Y})
					ballSound.Play()
				}
			}

			// Clear the window
			renderWindow.Clear(sf.Color{50, 200, 50, 0})

			if isPlaying {
				renderWindow.Draw(leftPaddle, sf.DefaultRenderStates())
				renderWindow.Draw(rightPaddle, sf.DefaultRenderStates())
				renderWindow.Draw(ball, sf.DefaultRenderStates())
			} else {
				renderWindow.Draw(pauseMessage, sf.DefaultRenderStates())
			}

			// Display things on screen
			renderWindow.Display()
		case <-AITicker.C:
			if ball.GetPosition().Y+ballRadius > rightPaddle.GetPosition().Y+paddleSize.Y/2 {
				rightPaddleSpeed = paddleSpeed
			} else if ball.GetPosition().Y-ballRadius < rightPaddle.GetPosition().Y-paddleSize.Y/2 {
				rightPaddleSpeed = -paddleSpeed
			} else {
				rightPaddleSpeed = 0
			}
		}
	}
}
