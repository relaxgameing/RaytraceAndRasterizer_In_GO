package homocoord

import "math"

type Vec2 struct {
	X, Y float32
}

type Vec3 struct {
	X, Y, Z float32
}

type Vec4 struct {
	X, Y, Z, W float32
}

func Vec3ToHomogeneous(v Vec3) Vec4 { return Vec4{v.X, v.Y, v.Z, 1.0} }

func HomogeneousToVec3(v Vec4) Vec3 {
	if v.W == 0 {
		return Vec3{v.X, v.Y, v.Z}
	} // vector
	return Vec3{v.X / v.W, v.Y / v.W, v.Z / v.W} // point
}

func (v *Vec3) Dot(p Vec3) float32 { return v.X*p.X + v.Y*p.Y + v.Z*p.Z }

func (v *Vec3) Add(p Vec3) Vec3      { return Vec3{v.X + p.X, v.Y + p.Y, v.Z + p.Z} }
func (v *Vec3) Subtract(p Vec3) Vec3 { return Vec3{v.X - p.X, v.Y - p.Y, v.Z - p.Z} }

func (v *Vec3) ScalarPrd(factor float32) Vec3 {
	return Vec3{
		v.X * factor,
		v.Y * factor,
		v.Z * factor,
	}
}

// Translation matrix
func Translation(tx, ty, tz float32) Mat4 {
	return Mat4{1, 0, 0, tx, 0, 1, 0, ty, 0, 0, 1, tz, 0, 0, 0, 1}
}

// Scale matrix
func Scale(sx, sy, sz float32) Mat4 {
	return Mat4{sx, 0, 0, 0, 0, sy, 0, 0, 0, 0, sz, 0, 0, 0, 0, 1}
}

// Rotation Y-axis (radians)
func RotateY(angle float32) Mat4 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))
	return Mat4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}
