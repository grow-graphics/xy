package xy

import "math"

const (
	x = iota
	y
	z
	w
)

type Axis int

const (
	X Axis = iota
	Y
	Z
	W
)

/*
Vector2 is a 2-element structure that can be used to represent 2D coordinates or any other pair of numeric values.

It uses floating-point coordinates. By default, these floating-point values use 32-bit precision, unlike [Float]
which is always 64-bit. If double precision is needed, compile the engine with the option precision=double.

See [Vector2i] for its integer counterpart.

Note: In a boolean context, a [Vector2] will evaluate to false if it's equal to Vector2{0, 0}. Otherwise, a [Vector2]
will always evaluate to true.
*/
type Vector2 [2]float

// NewVector2 constructs a new Vector2 from the given x and y.
func NewVector2(x, y float64) Vector2 { return Vector2{float(x), float(y)} } //Vector2(float,float)

func (v Vector2) X() float64 { return float64(v[x]) }
func (v Vector2) Y() float64 { return float64(v[y]) }

func (v *Vector2) SetY(y float64) { v[Y] = float(y) }
func (v *Vector2) SetX(x float64) { v[X] = float(x) }

func (v Vector2) ZERO() Vector2  { v[x], v[y] = 0, 0; return v }
func (v Vector2) ONE() Vector2   { v[x], v[y] = 1, 1; return v }
func (v Vector2) INF() Vector2   { v[x], v[y] = float(math.Inf(1)), float(math.Inf(1)); return v }
func (v Vector2) LEFT() Vector2  { v[x], v[y] = -1, 0; return v }
func (v Vector2) RIGHT() Vector2 { v[x], v[y] = 1, 0; return v }
func (v Vector2) UP() Vector2    { v[x], v[y] = 0, -1; return v }
func (v Vector2) DOWN() Vector2  { v[x], v[y] = 0, 1; return v }

// Vector2i constructs a new Vector2i from the given Vector2 by truncating components' fractional parts
// (rounding towards zero). For a different behavior consider passing the result of [Vector2.Ceil],
// [Vector2.Floor] or [Vector2.Round] to this constructor instead.
func (v Vector2) Vector2i() Vector2i { return Vector2i{int32(v[x]), int32(v[y])} } //Vector2i(Vector2)

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector2) Abs() Vector2 { return v.abs() } //Vector2.abs

// Angle Returns this vector's angle with respect to the positive X axis, or (1, 0) vector, in radians.
//
// For example, Const(Vector2.Right).Angle() will return zero, Const(Vector2.Down).Angle() will return
// PI / 2 (a quarter turn, or 90 degrees), and Vector2(1, -1).Angle() will return -Pi / 4
// (a negative eighth turn, or -45 degrees).
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle.png
//
// Equivalent to the result of [Atan2] when called with the vector's y and x as parameters:
//
//	Atan2(y, x).
func (v Vector2) Angle() Radians { return Radians(Atan2(float64(v[y]), float64(v[x]))) } //Vector2.angle

// AngleTo Returns the angle to the given vector, in radians.
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle_to.png
func (v Vector2) AngleTo(to Vector2) Radians { return Radians(Atan2(v.Cross(to), v.Dot(to))) } //Vector2.angle_to

// AngleToPoint returns the angle between the line connecting the two points and the X axis, in radians.
// a.AngleToPoint(b) is equivalent of doing (b - a).Angle().
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle_to_point.png
func (v Vector2) AngleToPoint(to Vector2) Radians { return Radians(v.Sub(to).Angle()) } //Vector2.angle_to_point

// Aspect returns the aspect ratio of this vector, the ratio of x to y.
func (v Vector2) Aspect() float64 { return float64(v[x] / v[y]) } //Vector2.aspect

// BezierDerivative returns the derivative at the given t on the Bézier curve defined by this
// vector and the given control_1, control_2, and end points.
func (v Vector2) BezierDerivative(control1, control2, end Vector2, t float64) Vector2 { //Vector2.bezier_derivative
	var res Vector2
	res[x] = BezierDerivative(res[x], control1[x], control2[x], end[x], float(t))
	res[y] = BezierDerivative(res[y], control1[y], control2[y], end[y], float(t))
	return res
}

// BezierInterpolate returns the point at the given t on the Bézier curve defined by this vector
// and the given control_1, control_2, and end points.
func (v Vector2) BezierInterpolate(control1, control2, end Vector2, t float64) Vector2 { //Vector2.bezier_interpolate
	var res Vector2
	res[x] = BezierInterpolate(res[x], control1[x], control2[x], end[x], float(t))
	res[y] = BezierInterpolate(res[y], control1[y], control2[y], end[y], float(t))
	return res
}

// Bounce returns a new vector "bounced off" from a plane defined by the given normal.
func (v Vector2) Bounce(n Vector2) Vector2 { //Vector2.bounce
	return Vector2.Sub(Vector2{}, v.Reflect(n))
}

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func (v Vector2) Ceil() Vector2 { return v.ceil() } //Vector2.ceil

// Clamp returns a new vector with all components clamped to the given min and max.
func (v Vector2) Clamp(from, to Vector2) Vector2 { //Vector2.clamp
	return Vector2{Clampf(v[x], from[x], to[x]), Clampf(v[y], from[y], to[y])}
}

// Cross returns the 2D analog of the cross product for this vector and with.
//
// This is the signed area of the parallelogram formed by the two vectors. If the second vector
// is clockwise from the first vector, then the cross product is the positive area. If
// counter-clockwise, the cross product is the negative area.
//
// Note: Cross product is not defined in 2D mathematically. This method embeds the 2D
// vectors in the XY plane of 3D space and uses their cross product's Z component as the analog.
func (v Vector2) Cross(other Vector2) float64 { return float64(v[x]*other[y] - v[y]*other[x]) } //Vector2.cross

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and
// post_b as handles, and returns the result at position weight. weight is on the range of
// 0.0 to 1.0, representing the amount of interpolation.
func (v Vector2) CubicInterpolate(b, preA, postB Vector2, weight float64) Vector2 { //Vector2.cubic_interpolate
	return Vector2{
		CubicInterpolate(v[x], b[x], preA[x], postB[x], float(weight)),
		CubicInterpolate(v[y], b[y], preA[y], postB[y], float(weight)),
	}
}

// CubicInterpolateInTime performs a cubic interpolation between this vector and b using pre_a
// and post_b as handles, and returns the result at position weight. weight is on the range of
// 0.0 to 1.0, representing the amount of interpolation.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func (v Vector2) CubicInterpolateInTime(b, preA, postB Vector2, weight, b_t, pre_a_t, post_b_t float64) Vector2 { //Vector2.cubic_interpolate_in_time
	return Vector2{
		CubicInterpolateInTime(v[x], b[x], preA[x], postB[x], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
		CubicInterpolateInTime(v[y], b[y], preA[y], postB[y], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
	}
}

// DirectionTo returns the normalized vector pointing from this vector to to. This is equivalent
// to using (b - a).Normalized().
func (v Vector2) DirectionTo(to Vector2) Vector2 { return v.Sub(to).Normalized() } //Vector2.direction_to

// DistanceSquaredTo returns the squared distance between this vector and to.
//
// This method runs faster than distance_to, so prefer it if you need to compare vectors or
// need the squared distance for some formula.
func (v Vector2) DistanceSquaredTo(to Vector2) float64 { //Vector2.distance_squared_to
	return float64((v[x]-to[x])*(v[x]-to[x]) + (v[y]-to[y])*(v[y]-to[y]))
}

// DistanceTo returns the distance between this vector and to.
func (v Vector2) DistanceTo(to Vector2) float64 { //Vector2.distance_to
	return float64(math.Sqrt(float64(v.DistanceSquaredTo(to))))
}

// Dot returns the dot product of this vector and with. This can be used to compare the angle
// between two vectors. For example, this can be used to determine whether an enemy is facing
// the player.
//
// The dot product will be 0 for a straight angle (90 degrees), greater than 0 for angles narrower
// than 90 degrees and lower than 0 for angles wider than 90 degrees.
//
// When using unit (normalized) vectors, the result will always be between -1.0 (180 degree angle)
// when the vectors are facing opposite directions, and 1.0 (0 degree angle) when the vectors are aligned.
//
// Note: Vector2.Dot(a,b) is equivalent to Vector2.Dot(b,a)
func (v Vector2) Dot(other Vector2) float64 { return float64(v[x]*other[x] + v[y]*other[y]) } //Vector2.dot

// Floor returns a new vector with all components rounded down (towards negative infinity).
func (v Vector2) Floor() Vector2 { return v.floor() } //Vector2.floor

// IsAproximatelyEqual returns true if this vector and to are approximately equal, by running
// [IsAproximatelyEqual] on each component.
func (v Vector2) IsApproximatelyEqual(to Vector2) bool { //Vector2.is_equal_approx
	return IsApproximatelyEqual(v[x], to[x]) && IsApproximatelyEqual(v[y], to[y])
}

// IsFinite returns true if this vector is finite, by calling [IsFinite] on each component.
func (v Vector2) IsFinite() bool { return IsFinite(v[x]) && IsFinite(v[y]) } //Vector2.is_finite

// IsNormalized returns true if the vector is normalized, i.e. its length is approximately equal to 1.
func (v Vector2) IsNormalized() bool { return IsApproximatelyEqual(v.LengthSquared(), 1) } //Vector2.is_normalized

// IsAproximatelyZero returns true if this vector is approximately equal to Vector2.Zero.
func (v Vector2) IsApproximatelyZero() bool { return v.IsApproximatelyEqual(Vector2{}) } //Vector2.is_zero_approx

// Length returns the length (magnitude) of this vector.
func (v Vector2) Length() float64 { //Vector2.length
	return float64(math.Sqrt(float64(v.LengthSquared())))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
func (v Vector2) LengthSquared() float64 { return float64(v[x]*v[x] + v[y]*v[y]) } //Vector2.length_squared

// Lerp returns the result of the linear interpolation between this vector and to by amount weight. weight
// is on the range of 0.0 to 1.0, representing the amount of interpolation.
func (v Vector2) Lerp(to Vector2, weight float64) Vector2 { return v.lerp(to, weight) } //Vector2.lerp

// LimitLength returns the vector with a maximum length by limiting its length to length.
func (v Vector2) LimitLength(length float64) Vector2 { //Vector2.limit_length
	var l = v.Length()
	if l > 0 && length < l {
		v = v.Mulf(1 / l)
		v = v.Mulf(length)
	}
	return v
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all
// components are equal, this method returns [X].
func (v Vector2) MaxAxis() Axis { //Vector2.max_axis_index
	if v[y] > v[x] {
		return Y
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all
// components are equal, this method returns [Y].
func (v Vector2) MinAxis() Axis { //Vector2.min_axis_index
	if v[x] < v[y] {
		return X
	}
	return Y
}

// MoveToward returns a new vector moved toward to by the fixed delta amount. Will not go past
// the final value.
func (v Vector2) MoveToward(to Vector2, delta float64) Vector2 { //Vector2.move_toward
	var vd = to.Sub(v)
	var len = vd.Length()
	if len <= delta || len < cmpEpsilon {
		return to
	}
	return v.Add(vd.Mulf(len * delta))
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.Length().
// See also is_normalized.
//
// Note: This function may return incorrect values if the input vector length is near zero.
func (v Vector2) Normalized() Vector2 { //Vector2.normalized
	length := v.Length()
	if length == 0 {
		return Vector2{}
	}
	return Vector2{float(float64(v[x]) / length), float(float64(v[y]) / length)}
}

// Orthogonal returns a perpendicular vector rotated 90 degrees counter-clockwise compared to the original,
// with the same length.
func (v Vector2) Orthogonal() Vector2 { return Vector2{v[y], -v[x]} } //Vector2.orthogonal

// Posmode returns a vector composed of the [Fposmod] of this vector's components and [Mod].
func (v Vector2) Posmod(mod float64) Vector2 { //Vector2.posmod
	return Vector2{Fposmod(v[x], float(mod)), Fposmod(v[y], float(mod))}
}

// Posmod returns a vector composed of the [Fposmod] of this vector's components and [Mod].
func (v Vector2) PosmodVector(mod Vector2) Vector2 { //Vector2.posmodv
	return Vector2{Fposmod(v[x], mod[x]), Fposmod(v[y], mod[y])}
}

// Project returns the result of projecting the vector onto the given vector b.
func (v Vector2) Project(b Vector2) Vector2 { //Vector2.project
	return b.Mulf(v.Dot(b) / b.LengthSquared())
}

// Reflect returns the result of reflecting the vector from a line defined by the given direction vector n.
func (v Vector2) Reflect(n Vector2) Vector2 { //Vector2.reflect
	return n.Mulf(2 * v.Dot(n)).Sub(v)
}

// Rotated returns the result of rotating this vector by angle (in radians).
func (v Vector2) Rotated(by Radians) Vector2 { //Vector2.rotated
	var cs = Cos(float(by))
	var sn = Sin(float(by))
	return Vector2{v[x]*cs - v[y]*sn, v[x]*sn + v[y]*cs}
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func (v Vector2) Round() Vector2 { return v.round() } //Vector2.round

// Sign returns a new vector with each component set to 1.0 if it's positive, -1.0 if it's negative,
// and 0.0 if it's zero. The result is identical to calling [Signf] on each component.
func (v Vector2) Sign() Vector2 { return v.sign() } //Vector2.sign

// Slerp returns the result of spherical linear interpolation between this vector and to, by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// This method also handles interpolating the lengths if the input vectors have different lengths. For the
// special case of one or both input vectors having zero length, this method behaves like lerp.
func (v Vector2) Slerp(to Vector2, weight Radians) Vector2 { //Vector2.slerp
	var start_length_sq = v.LengthSquared()
	var end_length_sq = to.LengthSquared()
	if start_length_sq == 0.0 || end_length_sq == 0.0 {
		// Zero length vectors have no angle, so the best we can do is either lerp or throw an error.
		return v.Lerp(to, float64(weight))
	}
	var start_length = Sqrt(start_length_sq)
	var result_length = Lerpf(start_length, Sqrt(end_length_sq), float64(weight))
	var angle = v.AngleTo(to)
	return v.Rotated(angle * weight).Mulf(result_length / start_length)
}

// Slide returns the result of sliding the vector along a plane defined by the given normal.
func (v Vector2) Slide(n Vector2) Vector2 { //Vector2.slide
	return v.Sub(n.Mulf(v.Dot(n)))
}

// Snapped returns a new vector with all components snapped to the nearest multiple of step.
func (v Vector2) Snapped(step Vector2) Vector2 { //Vector2.snapped
	return Vector2{
		Snappedf(v[x], step[x]),
		Snappedf(v[y], step[y]),
	}
}

func (v Vector2) Add(other Vector2) Vector2 { return Vector2{v[x] + other[x], v[y] + other[y]} } //Vector2 + Vector2
func (v Vector2) Sub(other Vector2) Vector2 { return Vector2{v[x] - other[x], v[y] - other[y]} } //Vector2 - Vector2
func (v Vector2) Mul(other Vector2) Vector2 { return Vector2{v[x] * other[x], v[y] * other[y]} } //Vector2 * Vector2
func (v Vector2) Div(other Vector2) Vector2 { return Vector2{v[x] / other[x], v[y] / other[y]} } //Vector2 / Vector2
func (v Vector2) Addf(other float64) Vector2 { // Vector2 + float
	return Vector2{v[x] + float(other), v[y] + float(other)}
}
func (v Vector2) Subf(other float64) Vector2 { // Vector2 - float
	return Vector2{v[x] - float(other), v[y] - float(other)}
}
func (v Vector2) Mulf(other float64) Vector2 { // Vector2 * float
	return Vector2{v[x] * float(other), v[y] * float(other)}
}
func (v Vector2) Divf(other float64) Vector2 { // Vector2 / float
	return Vector2{v[x] / float(other), v[y] / float(other)}
}
func (v Vector2) Neg() Vector2 { return Vector2{-v[x], -v[y]} } // -Vector2

func (v Vector2) Transform(t Transform2D) Vector2 { return Vector2{t.tdotx(v), t.tdoty(v)}.Add(t[2]) } //Transform2D * Vector2D
