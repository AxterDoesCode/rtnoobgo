package vec3

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func New (X, Y, Z float64) Vec3 {
    return Vec3{X, Y, Z}
}

func Dot(u, v Vec3) float64 {
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
