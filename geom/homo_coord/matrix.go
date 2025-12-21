package homocoord

type Mat4 [4][4]float32   // Row-major: m[0][0], m[0][1], m[0][2], m[0][3] | m[1][0]...
type Mat3x4 [3][4]float32 // m[0][0], m[0][1], m[0][2], m[0][3] | m[1][0]...

func Mat4Mul(A, B Mat4) Mat4 {
	var C Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func Mat4MulVec4(m Mat4, v Vec4) Vec4 {
	return Vec4{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W,
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W,
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W,
		m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W,
	}
}

func Mat3x4MulVec4(m Mat3x4, v Vec4) Vec3 {
	return Vec3{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W, // x' = x*d*cw/vw
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W, // y' = y*d*ch/vh
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W, // z' = z (w for divide)
	}
}
