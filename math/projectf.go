package mathgl

import (
	"math"
)

func Ortho(left, right, bottom, top, near, far float64) Mat4f {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)

	return Mat4f{float32(2. / rml), 0, 0, 0, 0, float32(2. / tmb), 0, 0, 0, 0, float32(-2. / fmn), 0, float32(-(right + left) / rml), float32(-(top + bottom) / tmb), float32(-(far + near) / fmn), 1}
}

func Ortho2D(left, right, top, bottom float64) Mat4f {
	return Ortho(left, right, top, bottom, -1, 1)
}

func Perspective(fovy, aspect, near, far float64) Mat4f {
	fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	nmf, f := near-far, 1./math.Tan(fovy/2)

	return Mat4f{float32(f / aspect), 0, 0, 0, 0, float32(f), 0, 0, 0, 0, float32((near + far) / nmf), -1, 0, 0, float32((2. * far * near) / nmf), 0}
}

func Frustum(left, right, bottom, top, near, far float64) Mat4f {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	A, B, C, D := (right+left)/rml, (top+bottom)/tmb, -(far+near)/fmn, (2*far*near)/fmn

	return Mat4f{float32((2. * near) / rml), 0, 0, 0, 0, float32((2. * near) / tmb), 0, 0, float32(A), float32(B), float32(C), -1, 0, 0, float32(D), 0}
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) Mat4f {
	F := Vec3f{
		float32(centerX - eyeX),
		float32(centerY - eyeY),
		float32(centerZ - eyeZ)}

	f := F.Normalize()

	Up := Vec3f{
		float32(upX),
		float32(upY),
		float32(upZ)}

	Upp := Up.Normalize()

	s := f.Cross(Upp)
	u := s.Cross(f)

	M := Mat4f{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(-eyeX, -eyeY, -eyeZ))
}

func LookAtV(eye, center, up Vec3f) Mat4f {
	F := center.Sub(eye)

	f := F.Normalize()

	Upp := up.Normalize()

	s := f.Cross(Upp)
	u := s.Cross(f)

	M := Mat4f{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(float64(-eye[0]), float64(-eye[1]), float64(-eye[2])))
}
