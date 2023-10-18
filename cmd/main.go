package main

func main() {
	O := NewVec3(0, 0, 0)
	canvas := NewCanvas(256, 256)
	for x := -canvas.Width / 2; x < canvas.Width/2; x++ {
		for y := -canvas.Height / 2; y < canvas.Height/2; y++ {
            D := CanvasToViewport(x, y, 1, 1, canvas.Width, canvas.Height)
            color := TraceRay
		}
	}
}
