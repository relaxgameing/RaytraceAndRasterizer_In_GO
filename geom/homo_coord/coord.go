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

func (v *Vec3) Add(p Vec3) Vec3      { return Vec3{v.X + p.X, v.Y + p.Y, v.Z + p.Z} }
func (v *Vec3) Subtract(p Vec3) Vec3 { return Vec3{v.X - p.X, v.Y - p.Y, v.Z - p.Z} }
func (v *Vec3) Dot(p Vec3) float32   { return v.X*p.X + v.Y*p.Y + v.Z*p.Z }
func (v *Vec3) Cross(p Vec3) Vec3 {
	return Vec3{
		X: v.Y*p.Z - v.Z*p.Y,
		Y: v.Z*p.X - v.X*p.Z,
		Z: v.X*p.Y - v.Y*p.X,
	}
}

func (v *Vec3) ScalarPrd(factor float32) Vec3 {
	return Vec3{
		v.X * factor,
		v.Y * factor,
		v.Z * factor,
	}
}

// * returns the point more towards the left
func LeftPoint(p1, p2 Vec3) (left, right Vec3) {
	if p1.X < p2.X {
		return p1, p2
	}
	return p2, p1
}

// *returns the point more towards the top
func UpperPoint(p1, p2 Vec3) (upper, lower Vec3) {
	if p1.Y >= p2.Y {
		return p1, p2
	}
	return p2, p1
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

func RotateX(angle float32) Mat4 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))
	return Mat4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

func RotateZ(angle float32) Mat4 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))
	return Mat4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}
