package canvas

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

func NewCanvas(w, h int) Canvas {
	tempCanvas := make([][]Color, h)
	for i := range tempCanvas {
		tempCanvas[i] = make([]Color, w)
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
    sy := (c.Height / 2) - x
    c.canvas[sy][sx] = color
}
