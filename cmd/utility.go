package main

type Scene struct {
	Spheres []Sphere
}

type Sphere struct {
	center Vec3
	radius int
	color Color
}

func CanvasToViewport(x, y, vw, vh, cw, ch int) vec3.Vec3 {
	d := 1
	return vec3.Vec3{X: x * (vw / cw), Y: y * (vh / ch), Z: d}
}

func IntersectRaySphere (O, D vec3.Vec3, sphere Sphere) {
    r := sphere.radius
    CO := //Continue from here
    //implement vector addition, negation and multiplication here
}

func TraceRay(Origin, D vec3.Vec3, t_min, t_max int, scene Scene) canvas.Color {
	closest_t := 999999
    var closest_sphere Sphere

	for _, sphere := range scene.Spheres {
		t1, t2 := IntersectRaySphere(Origin, D, sphere)
	}
}
