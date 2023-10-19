package main

import (
	"math"

	"github.com/AxterDoesCode/rtnoobgo/pkg/canvas"
	"github.com/AxterDoesCode/rtnoobgo/pkg/vec3"
)

func main() {
	Origin := vec3.New(0, 0, 0)
	scene := Scene{
		Spheres: []*Sphere{},
	}

	s2 := &Sphere{
		center: vec3.New(2, 0, 4),
		radius: 1,
		color:  canvas.NewColor(0, 0, 255),
	}

	s3 := &Sphere{
		center: vec3.New(-2, 0, 4),
		radius: 1,
		color:  canvas.NewColor(0, 255, 0),
	}

	s1 := &Sphere{
		center: vec3.New(0, 1, 5),
		radius: 1,
		color:  canvas.NewColor(255, 0, 0),
	}

    scene.AddSphere(s1)
	scene.AddSphere(s2)
	scene.AddSphere(s3)

	canvas := canvas.NewCanvas(256, 256)
	for x := -canvas.Width / 2; x <= canvas.Width/2; x++ {
		for y := -canvas.Height / 2; y <= canvas.Height/2; y++ {
			D := CanvasToViewport(x, y, canvas.Width, canvas.Height, 1, 1)
			color := TraceRay(Origin, D, 1, math.Inf(1), scene)
			canvas.PutPixel(x, y, color)
		}
	}
    canvas.PrintCanvas()
}
