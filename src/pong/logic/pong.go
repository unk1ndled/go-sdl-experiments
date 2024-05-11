package pong

import (
	"fmt"
	"math"
	"time"

	"github.com/unk1ndled/nier/src/ds"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ballRadius              = 5
	maxVelocity     float64 = 15
	collisionFactor         = 2
	minXVelocity            = maxVelocity / 3

	verticalboost = 1

	stickSpeed  = 7
	stickHeight = 150
	stickWidth  = 10
	//extra stick length to help users
	invisStick = 10

	distanceToBorder float64 = 50
)

var (
	screenWidth  int32
	screenHeight int32
)

type Ball struct {
	//ball center
	position     *ds.Vector2D
	velocity     *ds.Vector2D
	acceleration *ds.Vector2D
}

func (b *Ball) update() {
	b.velocity.Add(b.acceleration)
	b.velocity.ClampMagnitude(maxVelocity)
	if math.Abs(b.velocity[0]) < minXVelocity {
		b.velocity[0] = minXVelocity * (b.velocity[0] / math.Abs(b.velocity[0]))
	}
	b.acceleration.MultiplyByScalar(0)

	b.position.Add(b.velocity)
}

func (b *Ball) handleHorizontalStickCollision(extra int) {
	b.acceleration[0] = -b.velocity[0]
	b.acceleration[1] = math.Abs(float64(extra)) + math.Abs(b.velocity[1])
	if extra > 0 {
		b.acceleration[1] *= -1
	}
}

func (b *Ball) reset() {
	b.position = ds.NewVec2D(float64(screenWidth/2), float64(screenHeight/2))
	b.velocity = ds.RandomVec2D(10, 3).Normalized()
	b.acceleration = ds.RandomVec2D(0, 0)
	b.velocity.MultiplyByScalar(maxVelocity / 2)
}

func drawCircle(renderer *sdl.Renderer, centerX, centerY int32) {
	x := int32(ballRadius - 1)
	y := int32(0)
	dx := 2
	dy := 2
	err := dx - (ballRadius << 1)

	for x >= y {
		renderer.DrawPoint(centerX+x, centerY+y)
		renderer.DrawPoint(centerX+y, centerY+x)
		renderer.DrawPoint(centerX-y, centerY+x)
		renderer.DrawPoint(centerX-x, centerY+y)
		renderer.DrawPoint(centerX-x, centerY-y)
		renderer.DrawPoint(centerX-y, centerY-x)
		renderer.DrawPoint(centerX+y, centerY-x)
		renderer.DrawPoint(centerX+x, centerY-y)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (ballRadius << 1)
		}
	}
}

type Pong struct {
	renderer *sdl.Renderer

	leftStick  *ds.Vector2D
	rightStick *ds.Vector2D

	sticks []sdl.Rect

	ball Ball

	score [2]int8
}

func InitPong(sw, sh int32, renderer *sdl.Renderer) *Pong {
	screenWidth = sw
	screenHeight = sh
	pong := &Pong{

		renderer: renderer,

		leftStick:  ds.NewVec2D(distanceToBorder, float64(screenHeight)/2),
		rightStick: ds.NewVec2D(float64(screenWidth)-(distanceToBorder+stickWidth), float64(screenHeight)/2),

		ball: Ball{
			position:     ds.NewVec2D(float64(screenWidth/2), float64(screenHeight/2)),
			velocity:     ds.RandomVec2D(5, 1).Normalized(),
			acceleration: ds.RandomVec2D(0, 0),
		},

		score: [2]int8{0, 0},
	}
	pong.ball.velocity.MultiplyByScalar(maxVelocity / 2)

	pong.sticks = []sdl.Rect{
		{X: int32(pong.leftStick[0]), Y: int32(pong.leftStick[1]), W: stickWidth, H: stickHeight},
		{X: int32(pong.rightStick[0]), Y: int32(pong.rightStick[1]), W: stickWidth, H: stickHeight},
	}
	return pong
}

func (p *Pong) moveStick(stick *ds.Vector2D, isUp bool) {
	var increment float64
	sign := 1

	// 10 for extra space
	if isUp {
		if stick[1] > 10 {
			increment = stick[1]
		}
		sign *= -1
	} else {
		diff := float64(screenHeight) - (stick[1] + stickHeight)
		if diff > 10 {
			increment = diff
		}
	}
	if math.Abs(increment) > stickSpeed {
		increment = stickSpeed
	}
	stick.Add(ds.NewVec2D(0, float64(sign)*increment))
}

func (p *Pong) HandleInput() [2]int {

	// left index for left stick and right for right this array will contain a bonus y coordinate boost after collision
	extra := [2]int{0, 0}
	// Check for key presses
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_W] != 0 {
		p.moveStick(p.leftStick, true)
		extra[0] = verticalboost
	} else if keys[sdl.SCANCODE_S] != 0 {
		p.moveStick(p.leftStick, false)
		extra[0] = -verticalboost

	}
	if keys[sdl.SCANCODE_UP] != 0 {
		p.moveStick(p.rightStick, true)
		extra[1] = verticalboost
	} else if keys[sdl.SCANCODE_DOWN] != 0 {
		p.moveStick(p.rightStick, false)
		extra[1] = -verticalboost
	}
	//update sticks for drawing
	p.sticks[0].Y = int32(p.leftStick[1])
	p.sticks[1].Y = int32(p.rightStick[1])
	return extra

}

func (p *Pong) handleWallCollisions() {
	x, y := p.ball.position[0], p.ball.position[1]

	if y+ballRadius >= float64(screenHeight) || y-ballRadius <= 0 {
		p.ball.acceleration[1] = -collisionFactor * p.ball.velocity[1]
	}

	
	// Check collision with right wall
	if x+ballRadius >= float64(screenWidth) {
		p.handleWallCollisionReset(0)
	} else if x-ballRadius <= 0 { // Check collision with left wall
		p.handleWallCollisionReset(1)
	}
}

func (p *Pong) handleWallCollisionReset(player int) {
	p.ball.reset()
	p.score[player]++
	fmt.Printf("Score is:\nRight: %d\nLeft: %d\n", p.score[0], p.score[1])
	time.Sleep(500 * time.Millisecond)
}

func (p *Pong) handleCollisions(extra [2]int) {
	x, y := p.ball.position[0], p.ball.position[1]

	// Function to check hortizontal collision with a stick
	checkStickCollision := func(stickPos *ds.Vector2D) bool {
		return x+ballRadius >= stickPos[0] && x-ballRadius <= stickPos[0]+stickWidth &&
			(y+ballRadius >= stickPos[1]-invisStick && y-ballRadius <= stickPos[1]+stickHeight+invisStick)
	}

	// Check collision with right stick
	if checkStickCollision(p.rightStick) {
		p.ball.handleHorizontalStickCollision(extra[1])
		// Check collision with left stick
	} else if checkStickCollision(p.leftStick) {
		p.ball.handleHorizontalStickCollision(extra[0])
	}
	p.ball.acceleration.MultiplyByScalar(collisionFactor)
}

func (p *Pong) Update() {
	p.renderer.SetDrawColor(255, 255, 255, 255)
	p.renderer.DrawRects(p.sticks)
	drawCircle(p.renderer, int32(p.ball.position[0]), int32(p.ball.position[1]))
}

func (p *Pong) Play() {
	quit := false

	for !quit {
		p.renderer.SetDrawColor(0, 0, 0, 255)
		p.renderer.Clear()
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if e.GetType() == sdl.QUIT {
				quit = true
			}
		}

		//physics update
		extra := p.HandleInput()
		p.handleCollisions(extra)
		p.handleWallCollisions()
		p.ball.update()

		//vis update
		p.Update()

		p.renderer.SetDrawColor(255, 0, 0, 255)
		p.renderer.Present()
		time.Sleep(10 * time.Millisecond)
	}
}
