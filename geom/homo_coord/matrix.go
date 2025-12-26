package homocoord

type Mat4 [16]float32   // Row-major: [0,1,2,3 | 4,5,6,7 | 8,9,10,11 | 12,13,14,15]
type Mat3x4 [12]float32 // [0,1,2,3 | 4,5,6,7 | 8,9,10,11]

func IdentityMat4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Mat4Mul(A, B Mat4) Mat4 {
	var C Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				C[i*4+j] += A[i*4+k] * B[k*4+j]
			}
		}
	}
	return C
}

func Mat4MulVec4(m Mat4, v Vec4) Vec4 {
	return Vec4{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W,
		m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W,
	}
}

func Mat3x4MulVec4(m Mat3x4, v Vec4) Vec3 {
	return Vec3{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,   // x' = x*d*cw/vw
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,   // y' = y*d*ch/vh
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W, // z' = z (w for divide)
	}
}
