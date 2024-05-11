package pong

import (
	"math"
	"time"

	"github.com/unk1ndled/nier/src/ds"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ballRadius           = 5
	balldiameter         = 2 * ballRadius
	maxVelocity  float64 = 15

	stickSpeed  = 7
	stickHeight = 150
	stickWidth  = 10

	distanceToBorder float64 = 35
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

func (b *Ball) HandleWallCollisions() {
	x, y := b.position[0], b.position[1]
	if x+ballRadius >= float64(screenWidth) || x-ballRadius <= 0 {
		b.velocity[0] = -b.velocity[0]
	}
	if y+ballRadius >= float64(screenHeight) || y-ballRadius <= 0 {
		b.velocity[1] = -b.velocity[1]
	}

}

func (b *Ball) Update() {
	b.position.Add(b.velocity)

	//these walls
	b.velocity.Add(b.acceleration)
	b.velocity.ClampMagnitude(maxVelocity)
	b.HandleWallCollisions()
	b.acceleration.MultiplyByScalar(0)
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

func InitPong(screenwidth, screenheight int32, renderer *sdl.Renderer) *Pong {
	screenWidth = screenwidth
	screenHeight = screenheight
	pong := &Pong{

		renderer: renderer,

		leftStick: ds.NewVec2D(distanceToBorder, float64(screenheight)/2),
		//  should be float64(screenwidth)-*distanceToBorder
		rightStick: ds.NewVec2D(float64(screenwidth)-(distanceToBorder+stickWidth), float64(screenheight)/2),

		ball: Ball{
			position:     ds.NewVec2D(float64(screenwidth/2), float64(screenheight/2)),
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

func (p *Pong) HandleInput() {

	// Check for key presses
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_W] != 0 {
		p.moveStick(p.leftStick, true)
	}
	if keys[sdl.SCANCODE_S] != 0 {
		p.moveStick(p.leftStick, false)
	}
	if keys[sdl.SCANCODE_UP] != 0 {
		p.moveStick(p.rightStick, true)
	}
	if keys[sdl.SCANCODE_DOWN] != 0 {
		p.moveStick(p.rightStick, false)
	}
	p.sticks[0].Y = int32(p.leftStick[1])
	p.sticks[1].Y = int32(p.rightStick[1])
}

func (p *Pong) Update() {
	p.renderer.SetDrawColor(255, 255, 255, 255)
	p.renderer.DrawRects(p.sticks)

	p.ball.Update()
	drawCircle(p.renderer, int32(p.ball.position[0]), int32(p.ball.position[1]))
	p.handleCollisions()
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

		p.HandleInput()
		p.Update()

		p.renderer.SetDrawColor(255, 0, 0, 255)
		p.renderer.Present()
		time.Sleep(10 * time.Millisecond)
	}
}

func (p *Pong) handleCollisions() {
	x, y := p.ball.position[0], p.ball.position[1]
	if x+ballRadius >= p.rightStick[0] && (y+ballRadius >= p.rightStick[1] && y-ballRadius <= p.rightStick[1]+stickHeight) {
		p.ball.velocity[0] = -p.ball.velocity[0]
	} else if x-ballRadius <= p.leftStick[0] && (y+ballRadius >= p.leftStick[1] && y-ballRadius <= p.leftStick[1]+stickHeight) {
		p.ball.velocity[0] = -p.ball.velocity[0]
	}

}
