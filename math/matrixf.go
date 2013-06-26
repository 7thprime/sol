package mathgl

import (
//"math"
)

type Mat2f [4]float32
type Mat2x3f [6]float32
type Mat2x4f [8]float32
type Mat3x2f [6]float32
type Mat3f [9]float32
type Mat3x4f [12]float32
type Mat4x2f [8]float32
type Mat4x3f [12]float32
type Mat4f [16]float32

func Ident2f() Mat2f {
	return Mat2f{1, 0, 0, 1}
}

func Ident3f() Mat3f {
	return Mat3f{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

func Ident4f() Mat4f {
	return Mat4f{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

func dimensionalAdd (r1 []float32, r2 []float32, width int, height int) ([]float32) {
    size := width * height
    var r3 [size]float32
    for i := 0; i < size ; i++ {
        r3[i] = r1[i] + r2[i]
    }
}

func (m1 Mat2f) Add(m2 Mat2f) Mat2f {
	return Mat2f{
        m1[0] + m2[0], m1[1] + m2[1],
        m1[2] + m2[2], m1[3] + m2[3]}
}

func (m1 Mat2x3f) Add(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f{
        m1[0] + m2[0], m1[1] + m2[1],
        m1[2] + m2[2], m1[3] + m2[3],
        m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Mat2x4f) Add(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f{
        m1[0] + m2[0], m1[1] + m2[1],
        m1[2] + m2[2], m1[3] + m2[3],
        m1[4] + m2[4], m1[5] + m2[5],
        m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Mat3x2f) Add(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f{
        m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2],
        m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Mat3f) Add(m2 Mat3f) Mat3f {
	return Mat3f{
        m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2],
        m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5],
        m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

func (m1 Mat3x4f) Add(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f{
        m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2],
        m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5],
        m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8],
        m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Mat4x2f) Add(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Mat4x3f) Add(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Mat4f) Add(m2 Mat4f) Mat4f {
	return Mat4f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

func (m1 Mat2f) Sub(m2 Mat2f) Mat2f {
	return Mat2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

func (m1 Mat2x3f) Sub(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Mat2x4f) Sub(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Mat3x2f) Sub(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Mat3f) Sub(m2 Mat3f) Mat3f {
	return Mat3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

func (m1 Mat3x4f) Sub(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Mat4x2f) Sub(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Mat4x3f) Sub(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Mat4f) Sub(m2 Mat4f) Mat4f {
	return Mat4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

func (m1 Mat2f) Mul(c float32) Mat2f {
	return Mat2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

func (m1 Mat2x3f) Mul(c float32) Mat2x3f {
	return Mat2x3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Mat2x4f) Mul(c float32) Mat2x4f {
	return Mat2x4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Mat3x2f) Mul(c float32) Mat3x2f {
	return Mat3x2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Mat3f) Mul(c float32) Mat3f {
	return Mat3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

func (m1 Mat3x4f) Mul(c float32) Mat3x4f {
	return Mat3x4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Mat4x2f) Mul(c float32) Mat4x2f {
	return Mat4x2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Mat4x3f) Mul(c float32) Mat4x3f {
	return Mat4x3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Mat4f) Mul(c float32) Mat4f {
	return Mat4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

func (m1 Mat2f) Mul2x1(m2 Vec2f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

func (m1 Mat2f) Mul2(m2 Mat2f) Mat2f {
	return Mat2f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

func (m1 Mat2f) Mul2x3(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

func (m1 Mat2f) Mul2x4(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

func (m1 Mat2x3f) Mul3x1(m2 Vec3f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

func (m1 Mat2x3f) Mul3x2(m2 Mat3x2f) Mat2f {
	return Mat2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

func (m1 Mat2x3f) Mul3(m2 Mat3f) Mat2x3f {
	return Mat2x3f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

func (m1 Mat2x3f) Mul3x4(m2 Mat3x4f) Mat2x4f {
	return Mat2x4f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

func (m1 Mat2x4f) Mul4x1(m2 Vec4f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

func (m1 Mat2x4f) Mul4x2(m2 Mat4x2f) Mat2f {
	return Mat2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

func (m1 Mat2x4f) Mul4x3(m2 Mat4x3f) Mat2x3f {
	return Mat2x3f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

func (m1 Mat2x4f) Mul4(m2 Mat4f) Mat2x4f {
	return Mat2x4f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

func (m1 Mat3x2f) Mul2x1(m2 Vec2f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

func (m1 Mat3x2f) Mul2(m2 Mat2f) Mat3x2f {
	return Mat3x2f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

func (m1 Mat3x2f) Mul2x3(m2 Mat2x3f) Mat3f {
	return Mat3f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

func (m1 Mat3x2f) Mul2x4(m2 Mat2x4f) Mat3x4f {
	return Mat3x4f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

func (m1 Mat3f) Mul3x1(m2 Vec3f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

func (m1 Mat3f) Mul3x2(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

func (m1 Mat3f) Mul3(m2 Mat3f) Mat3f {
	return Mat3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

func (m1 Mat3f) Mul3x4(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

func (m1 Mat3x4f) Mul4x1(m2 Vec4f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

func (m1 Mat3x4f) Mul4x2(m2 Mat4x2f) Mat3x2f {
	return Mat3x2f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

func (m1 Mat3x4f) Mul4x3(m2 Mat4x3f) Mat3f {
	return Mat3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

func (m1 Mat3x4f) Mul4(m2 Mat4f) Mat3x4f {
	return Mat3x4f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

func (m1 Mat4x2f) Mul2x1(m2 Vec2f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

func (m1 Mat4x2f) Mul2(m2 Mat2f) Mat4x2f {
	return Mat4x2f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

func (m1 Mat4x2f) Mul2x3(m2 Mat2x3f) Mat4x3f {
	return Mat4x3f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

func (m1 Mat4x2f) Mul2x4(m2 Mat2x4f) Mat4f {
	return Mat4f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

func (m1 Mat4x3f) Mul3x1(m2 Vec3f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

func (m1 Mat4x3f) Mul3x2(m2 Mat3x2f) Mat4x2f {
	return Mat4x2f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

func (m1 Mat4x3f) Mul3(m2 Mat3f) Mat4x3f {
	return Mat4x3f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

func (m1 Mat4x3f) Mul3x4(m2 Mat3x4f) Mat4f {
	return Mat4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

func (m1 Mat4f) Mul4x1(m2 Vec4f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

func (m1 Mat4f) Mul4x2(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

func (m1 Mat4f) Mul4x3(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

func (m1 Mat4f) Mul4(m2 Mat4f) Mat4f {
	return Mat4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}

func (m1 Mat2f) Transpose() Mat2f {
	return Mat2f{m1[0], m1[2], m1[1], m1[3]}
}

func (m1 Mat2x3f) Transpose() Mat3x2f {
	return Mat3x2f{m1[0], m1[3], m1[1], m1[4], m1[2], m1[5]}
}

func (m1 Mat2x4f) Transpose() Mat4x2f {
	return Mat4x2f{m1[0], m1[4], m1[1], m1[5], m1[2], m1[6], m1[3], m1[7]}
}

func (m1 Mat3x2f) Transpose() Mat2x3f {
	return Mat2x3f{m1[0], m1[2], m1[4], m1[1], m1[3], m1[5]}
}

func (m1 Mat3f) Transpose() Mat3f {
	return Mat3f{m1[0], m1[3], m1[6], m1[1], m1[4], m1[7], m1[2], m1[5], m1[8]}
}

func (m1 Mat3x4f) Transpose() Mat4x3f {
	return Mat4x3f{m1[0], m1[4], m1[8], m1[1], m1[5], m1[9], m1[2], m1[6], m1[10], m1[3], m1[7], m1[11]}
}

func (m1 Mat4x2f) Transpose() Mat2x4f {
	return Mat2x4f{m1[0], m1[2], m1[4], m1[6], m1[1], m1[3], m1[5], m1[7]}
}

func (m1 Mat4x3f) Transpose() Mat3x4f {
	return Mat3x4f{m1[0], m1[3], m1[6], m1[9], m1[1], m1[4], m1[7], m1[10], m1[2], m1[5], m1[8], m1[11]}
}

func (m1 Mat4f) Transpose() Mat4f {
	return Mat4f{m1[0], m1[4], m1[8], m1[12], m1[1], m1[5], m1[9], m1[13], m1[2], m1[6], m1[10], m1[14], m1[3], m1[7], m1[11], m1[15]}
}