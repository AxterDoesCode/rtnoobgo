package vec3

type Vec3 struct {
	X int
	Y int
	Z int
}

func New (X, Y, Z int) Vec3 {
    return Vec3{X, Y, Z}
}

func Dot(u, v Vec3) int {
    return u.X * v.X + u.Y * v.Y + u.Z * v.Z 
}

func (v Vec3) Neg () Vec3 {
    return Vec3{
        X: -v.X,
        Y: -v.Y,
        Z: -v.Z,
    }
}

func Add (u, v Vec3) Vec3 {
    return Vec3{
        X: u.X + v.X,
        Y: u.Y + v.Y,
        Z: u.Z + v.Z,
    }
}
