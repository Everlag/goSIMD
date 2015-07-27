package simd

type Vec4 [4]float32

func (v1 Vec4) Mul(v2 Vec4) Vec4 {
	return Vec4{
		v1[0] * v2[0],
		v1[1] * v2[1],
		v1[2] * v2[2],
		v1[3] * v2[3],
	}
}

func mul(v1, v2 Vec4) Vec4

func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
		v1[3] + v2[3],
	}
}

func add(v1, v2 Vec4) Vec4
