package main

import (
    "math"
    "github.com/AxterDoesCode/rtnoobgo/pkg/vec3"
    "github.com/AxterDoesCode/rtnoobgo/pkg/canvas"
)

type Scene struct {
	Spheres []Sphere
}

type Sphere struct {
	center vec3.Vec3
	radius int
	color canvas.Color
}

func CanvasToViewport(x, y, vw, vh, cw, ch int) vec3.Vec3 {
    var d float64 = 1
	return vec3.Vec3{X: x * (vw / cw), Y: y * (vh / ch), Z: d}
}

func IntersectRaySphere (O, D vec3.Vec3, sphere Sphere) (float64, float64) {
    r := sphere.radius
    CO := vec3.Add(O, D.Neg())

    a := vec3.Dot(D, D)
    b := 2 * vec3.Dot(CO, D)
    c := vec3.Dot(CO, CO) - float64(r*r)

    discriminant := b*b - 4*a*c
    if discriminant < 0 {
        return int(math.Inf(1)), int(math.Inf(1))
    }

    t1 := (-b + math.Sqrt(discriminant)) / (2*a)
    t2 := (-b - math.Sqrt(discriminant)) / (2*a)
    return t1, t2
}

func TraceRay(Origin, D vec3.Vec3, t_min, t_max int, scene Scene) Color {
	closest_t := math.Inf(1)
    var closest_sphere Sphere

	for _, sphere := range scene.Spheres {
		t1, t2 := IntersectRaySphere(Origin, D, sphere)
	}
}
