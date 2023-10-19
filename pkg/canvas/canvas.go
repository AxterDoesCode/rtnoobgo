package canvas

import "fmt"

type Color struct {
	R int
	G int
	B int
}

type Canvas struct {
	Width  int
	Height int
	canvas [][]Color
}

func NewColor(R, G, B int) Color {
	return Color{R, G, B}
}

func NewCanvas(w, h int) Canvas {
	tempCanvas := make([][]Color, h+1)
	for i := range tempCanvas {
		tempCanvas[i] = make([]Color, w+1)
	}
	return Canvas{
		Width:  w,
		Height: h,
		canvas: tempCanvas,
	}
}

func (c *Canvas) PutPixel(x, y int, color Color) {
	// Convert our coordinate system to system coords
	sx := (c.Width / 2) + x
	sy := (c.Height / 2) - y
	c.canvas[sy][sx] = color
}

func (c *Canvas) PrintCanvas() {
	fmt.Printf("P3\n%v %v\n255\n", c.Width, c.Height)
	for y := 0; y < c.Height; y++ {
        for x := 0; x < c.Width; x++ {
            temp := c.canvas[y][x]
            fmt.Printf("%v %v %v\n", temp.R, temp.G, temp.B)
        }
	}
}
