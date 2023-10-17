package vec3

type Vec3 struct {
	X int
	Y int
	Z int
}

func NewVec3 (X, Y, Z int) Vec3 {
    return Vec3{X, Y, Z}
}
