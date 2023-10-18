package main

import (
    "math"
    "github.com/AxterDoesCode/rtnoobgo/pkg/vec3"
)

type Scene struct {
	Spheres []Sphere
}

type Sphere struct {
	center vec3.Vec3
	radius int
	color Color
}

func CanvasToViewport(x, y, vw, vh, cw, ch int) vec3.Vec3 {
	d := 1
	return vec3.Vec3{X: x * (vw / cw), Y: y * (vh / ch), Z: d}
}

func IntersectRaySphere (O, D vec3.Vec3, sphere Sphere) (int, int) {
    r := sphere.radius
    CO := Add(O, D.Neg())

    a := Dot(D, D)
    b := 2 * Dot(CO, D)
    c := Dot(CO, CO) - r*r

    discriminant := b*b - 4*a*c
    if discriminant < 0 {
        return int(math.Inf(1)), int(math.Inf(1))
    }

    t1 := ()
}

func TraceRay(Origin, D vec3.Vec3, t_min, t_max int, scene Scene) Color {
	closest_t := 999999
    var closest_sphere Sphere

	for _, sphere := range scene.Spheres {
		t1, t2 := IntersectRaySphere(Origin, D, sphere)
	}
}
