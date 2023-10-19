package main

import (
    "math"
    "github.com/AxterDoesCode/rtnoobgo/pkg/vec3"
    "github.com/AxterDoesCode/rtnoobgo/pkg/canvas"
)

type Scene struct {
	Spheres []*Sphere
}

type Sphere struct {
	center vec3.Vec3
	radius int
	color canvas.Color
}

func (s *Scene) AddSphere (sphere *Sphere) {
    s.Spheres = append(s.Spheres, sphere)
}

func CanvasToViewport(x, y, cw, ch int, vw, vh float64) vec3.Vec3 {
    var d float64 = 1
	return vec3.Vec3{X: float64(x) * (vw / float64(cw)), Y: float64(y) * (vh / float64(ch)), Z: d}
}

func IntersectRaySphere (O, D vec3.Vec3, sphere Sphere) (float64, float64) {
    r := sphere.radius
    CO := vec3.Add(O, D.Neg())

    a := vec3.Dot(D, D)
    b := 2 * vec3.Dot(CO, D)
    c := vec3.Dot(CO, CO) - float64(r*r)

    discriminant := (b*b) - (4*a*c)
    if discriminant < 0 {
        return math.Inf(1), math.Inf(1)
    }

    t1 := (-b + math.Sqrt(discriminant)) / (2*a)
    t2 := (-b - math.Sqrt(discriminant)) / (2*a)
    return t1, t2
}

func TraceRay(Origin, D vec3.Vec3, t_min, t_max float64, scene Scene) canvas.Color {
	closest_t := math.Inf(1)
    var closest_sphere *Sphere

	for _, sphere := range scene.Spheres {
		t1, t2 := IntersectRaySphere(Origin, D, *sphere)
        if t1 >= t_min && t1 <= t_max && t1 < closest_t {
            closest_t = t1
            closest_sphere = sphere
        }

        if t2 >= t_min && t2 <= t_max && t2 < closest_t {
            closest_t = t2
            closest_sphere = sphere
        }
	}
    if closest_sphere == nil {
        return canvas.NewColor(0, 0, 0)
    }
    return closest_sphere.color
}
