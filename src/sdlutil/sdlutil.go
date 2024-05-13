package sdlutil

import "github.com/unk1ndled/nier/src/sdlutil/digits"

type Color struct {
	R, G, B byte
}

type SdlContext struct {
	pixels       []byte
	screenWidth  int
	screenHeight int
}

func NewSdlContext(pixels []byte, screenWidth int, screenHeight int) *SdlContext {
	return &SdlContext{pixels: pixels, screenWidth: screenWidth, screenHeight: screenHeight}
}

func (sc *SdlContext) SetPixel(x, y int, c *Color) {
	index := (x + (sc.screenWidth * y)) * 4
	if index+4 <= len(sc.pixels) && index >= 0 {
		sc.pixels[index] = c.R
		sc.pixels[index+1] = c.G
		sc.pixels[index+2] = c.B
	}
}

func (sc *SdlContext) DrawRect(x, y, width, height int, color *Color) {

	x -= width / 2
	y -= height / 2
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			sc.SetPixel(x+i, y+j, color)
		}
	}

}

func (sc *SdlContext) DrawDigit(digit byte, x, y int, cellSize int, color *Color) {

	//cell border
	border := 1
	//a cell as i had forseen is a building block, a pixel with a defined width and height
	m := digits.Dictionary[digit]

	//centering x and y are the coordinates of the center of the digits thus the next transformation
	// 3 cuz a digit is made out of 3 cells width and 5 height
	digitx := x - (3 * cellSize / 2)
	digity := y - (5 * cellSize / 2)

	for index, cell := range m {
		//im dumb so ill explain it to future me
		// checking if the cell on the array has a value of 1
		if cell == 1 {
			//from array to grid index transformation
			cellx := index % 3
			celly := index / 3
			//multiplying by the cellsize
			xoffset := cellx * cellSize
			yoffset := celly * cellSize
			//drawing a cell
			for i := border; i < cellSize-border; i++ {
				for j := border; j < cellSize-border; j++ {
					sc.SetPixel(digitx+xoffset+i, digity+yoffset+j, color)
				}
			}
		}

	}
}

func (sc *SdlContext) DrawLine(startX, startY, endX, endY int, color *Color) {

	dx := endX - startX
	dy := endY - startY

	D := 2*dy - dx
	y := startY
	for x := startX; x < endX; x++ {
		sc.SetPixel(x, y, color)
		if D > 0 {
			y = y + 1
			D = D - 2*dx
		}
		D = D + 2*dy
	}

}

func (sc *SdlContext) DrawCircle(centerX, centerY, radius int, color *Color) {
	x := radius - 1
	y := int(0)
	dx := 2
	dy := 2
	err := dx - (radius << 1)

	for x >= y {
		sc.SetPixel(centerX+x, centerY+y, color)
		sc.SetPixel(centerX+y, centerY+x, color)
		sc.SetPixel(centerX-y, centerY+x, color)
		sc.SetPixel(centerX-x, centerY+y, color)
		sc.SetPixel(centerX-x, centerY-y, color)
		sc.SetPixel(centerX-y, centerY-x, color)
		sc.SetPixel(centerX+y, centerY-x, color)
		sc.SetPixel(centerX+x, centerY-y, color)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (int(radius) << 1)
		}
	}
}
