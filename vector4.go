package xy

import "math"

/*
Vector4 A 4-element structure that can be used to represent 4D coordinates or any other quadruplet of numeric values.

It uses floating-point coordinates. By default, these floating-point values use 32-bit precision, unlike float which
is always 64-bit. If double precision is needed, compile the engine with the option precision=double.

See [Vector4i] for its integer counterpart.

Note: In a boolean context, a [Vector4] will evaluate to false if it's equal to Vector4{0, 0, 0, 0}. Otherwise, a
[Vector4] will always evaluate to true.
*/
type Vector4 [4]float

// NewVector4 constructs a new Vector4 from the given x, y, z, and w.
func NewVector4(x, y, z, w float64) Vector4 { //Vector4(float,float,float,float)
	return Vector4{float(x), float(y), float(z), float(w)}
}

func (v Vector4) X() float64 { return float64(v[X]) }
func (v Vector4) Y() float64 { return float64(v[Y]) }
func (v Vector4) Z() float64 { return float64(v[Z]) }
func (v Vector4) W() float64 { return float64(v[W]) }

func (v *Vector4) SetX(x float64) { v[X] = float(x) }
func (v *Vector4) SetY(y float64) { v[Y] = float(y) }
func (v *Vector4) SetZ(z float64) { v[Z] = float(z) }
func (v *Vector4) SetW(w float64) { v[W] = float(w) }

// "Constants"

func (Vector4) ZERO() Vector4 { return Vector4{} }
func (Vector4) ONE() Vector4  { return Vector4{1, 1, 1, 1} }
func (Vector4) INF() Vector4 {
	return Vector4{float(math.Inf(0)), float(math.Inf(0)), float(math.Inf(0)), float(math.Inf(0))}
}

func (v Vector4) Vector4i() Vector4i { //Vector4i(Vector4)
	return Vector4i{int32(v[X]), int32(v[Y]), int32(v[Z]), int32(v[W])}
}

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector4) Abs() Vector4 { return v.abs() } //Vector4.abs

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func (v Vector4) Ceil() Vector4 { return v.ceil() } //Vector4.ceil

// Clamp returns a new vector with all components clamped between the components of min and max, by running
// [Clampf] on each component.
func (v Vector4) Clamp(min, max Vector4) Vector4 { //Vector4.clamp
	return Vector4{
		Clampf(v[X], min[X], max[X]),
		Clampf(v[Y], min[Y], max[Y]),
		Clampf(v[Z], min[Z], max[Z]),
		Clampf(v[W], min[W], max[W]),
	}
}

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of
// interpolation.
func (v Vector4) CubicInterpolate(b, pre_a, post_b Vector4, weight float64) Vector4 { //Vector4.cubic_interpolate
	return Vector4{
		CubicInterpolate(v[X], b[X], pre_a[X], post_b[X], float(weight)),
		CubicInterpolate(v[Y], b[Y], pre_a[Y], post_b[Y], float(weight)),
		CubicInterpolate(v[Z], b[Z], pre_a[Z], post_b[Z], float(weight)),
		CubicInterpolate(v[W], b[W], pre_a[W], post_b[W], float(weight)),
	}
}

// CubicInyerpolateInTime performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of
// interpolation.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func (v Vector4) CubicInterpolateInTime(b, pre_a, post_b Vector4, weight, b_t, pre_a_t, post_b_t float64) Vector4 { //Vector4.cubic_interpolate_in_time
	return Vector4{
		CubicInterpolateInTime(v[X], b[X], pre_a[X], post_b[X], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
		CubicInterpolateInTime(v[Y], b[Y], pre_a[Y], post_b[Y], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
		CubicInterpolateInTime(v[Z], b[Z], pre_a[Z], post_b[Z], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
		CubicInterpolateInTime(v[W], b[W], pre_a[W], post_b[W], float(weight), float(b_t), float(pre_a_t), float(post_b_t)),
	}
}

// DirectionTo returns the normalized vector pointing from this vector to to. This is equivalent to using
// (b - a).Normalized().
func (v Vector4) DirectionTo(to Vector4) Vector4 { //Vector4.direction_to
	return to.Sub(v).Normalized()
}

// DistanceSquaredTo returns the squared distance between this vector and to.
//
// This method runs faster than DistanceTo, so prefer it if you need to compare vectors or need the squared
// distance for some formula.
func (v Vector4) DistanceSquaredTo(to Vector4) float64 { //Vector4.distance_squared_to
	return v.Sub(to).LengthSquared()
}

// DistanceTo returns the distance between this vector and to.
func (v Vector4) DistanceTo(to Vector4) float64 { //Vector4.distance_to
	return v.Sub(to).Length()
}

// Dot returns the dot product of this vector and b.
func (v Vector4) Dot(b Vector4) float64 { //Vector4.dot
	return float64(v[X])*float64(b[X]) + float64(v[Y])*float64(b[Y]) + float64(v[Z])*float64(b[Z]) + float64(v[W])*float64(b[W])
}

// Floor returns a new vector with all components rounded down (towards negative infinity).
func (v Vector4) Floor() Vector4 { return v.floor() } //Vector4.floor

// Inverse returns the inverse of the vector. This is the same as
//
//	Vector4{1.0 / v[X], 1.0 / v[Y], 1.0 / v[Z], 1.0 / v[W]}.
func (v Vector4) Inverse() Vector4 { //Vector4.inverse
	return Vector4{1.0 / v[X], 1.0 / v[Y], 1.0 / v[Z], 1.0 / v[W]}
}

// IsApproximatelyEqual returns true if this vector and to are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func (v Vector4) IsApproximatelyEqual(to Vector4) bool { //Vector4.is_equal_approx
	return IsApproximatelyEqual(v[X], to[X]) && IsApproximatelyEqual(v[Y], to[Y]) && IsApproximatelyEqual(v[Z], to[Z]) && IsApproximatelyEqual(v[W], to[W])
}

// IsFinite returns true if this vector is finite, by calling [IsFinite] on each component.
func (v Vector4) IsFinite() bool { //Vector4.is_finite
	return IsFinite(v[X]) && IsFinite(v[Y]) && IsFinite(v[Z]) && IsFinite(v[W])
}

// IsNormalized returns true if this vector is normalized, by checking if its length is approximately 1.0.
func (v Vector4) IsNormalized() bool { return IsApproximatelyEqual(v.LengthSquared(), 1.0) } //Vector4.is_normalized

// IsApproximatelyZero returns true if this vector's values are approximately zero, by running
// [IsApproximatelyZero] on each component.
//
// This method is faster than using is_equal_approx with one value as a zero vector.
func (v Vector4) IsApproximatelyZero() bool { //Vector4.is_zero_approx
	return IsApproximatelyZero(v[X]) && IsApproximatelyZero(v[Y]) && IsApproximatelyZero(v[Z]) && IsApproximatelyZero(v[W])
}

// Length returns the length (magnitude) of this vector.
func (v Vector4) Length() float64 { return Sqrt(v.LengthSquared()) } //Vector4.length

// LengthSquared returns the squared length (squared magnitude) of this vector.
func (v Vector4) LengthSquared() float64 { //Vector4.length_squared
	return float64(v[X])*float64(v[X]) + float64(v[Y])*float64(v[Y]) + float64(v[Z])*float64(v[Z]) + float64(v[W])*float64(v[W])
}

// Lerp returns the result of the linear interpolation between this vector and to by amount weight. weight
// is on the range of 0.0 to 1.0, representing the amount of interpolation.
func (v Vector4) Lerp(to Vector4, weight float64) Vector4 { //Vector4.lerp
	return v.lerp(to, weight)
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func (v Vector4) MaxAxis() Axis { //Vector4.max_axis_index
	var max_index = 0
	var max_value = v[X]
	for i := 1; i < len(v); i++ {
		if v[i] > max_value {
			max_index = i
			max_value = v[i]
		}
	}
	return Axis(max_index)
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [W].
func (v Vector4) MinAxis() Axis { //Vector4.min_axis_index
	var min_index = 0
	var min_value = v[X]
	for i := 1; i < len(v); i++ {
		if v[i] <= min_value {
			min_index = i
			min_value = v[i]
		}
	}
	return Axis(min_index)
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.length().
// See also [Vector4.IsNormalized].
//
// Note: This function may return incorrect values if the input vector length is near zero.
func (v Vector4) Normalized() Vector4 { //Vector4.normalized
	return v.Divf(v.Length())
}

// Posmod returns a vector composed of the [Fposmod] of this vector's components and mod.
func (v Vector4) Posmodf(mod float64) Vector4 { //Vector4.posmod
	return Vector4{
		Fposmod(v[X], float(mod)),
		Fposmod(v[Y], float(mod)),
		Fposmod(v[Z], float(mod)),
		Fposmod(v[W], float(mod)),
	}
}

// Posmod returns a vector composed of the [Fposmod] of this vector's components and modv's components.
func (v Vector4) Posmodv(mod Vector4) Vector4 { //Vector4.posmodv
	return Vector4{
		Fposmod(v[X], mod[X]),
		Fposmod(v[Y], mod[Y]),
		Fposmod(v[Z], mod[Z]),
		Fposmod(v[W], mod[W]),
	}
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func (v Vector4) Round() Vector4 { return v.round() } //Vector4.round

// Sign returns a new vector with each component set to 1.0 if it's positive, -1.0 if it's negative,
// and 0.0 if it's zero. The result is identical to calling [Signf] on each component.
func (v Vector4) Sign() Vector4 { //Vector4.sign
	return Vector4{
		Signf(v[X]),
		Signf(v[Y]),
		Signf(v[Z]),
		Signf(v[W]),
	}
}

// Snapped returns a new vector with all components snapped to the nearest multiple of step.
// This can also be used to round the components to an arbitrary number of decimals.
func (v Vector4) Snapped(step Vector4) Vector4 { //Vector4.snapped
	return Vector4{
		Snappedf(v[X], step[X]),
		Snappedf(v[Y], step[Y]),
		Snappedf(v[Z], step[Z]),
		Snappedf(v[W], step[W]),
	}
}

func (v Vector4) Add(other Vector4) Vector4 { //Vector4 + Vector4
	return Vector4{v[X] + other[X], v[Y] + other[Y], v[Z] + other[Z], v[W] + other[W]}
}
func (v Vector4) Sub(other Vector4) Vector4 { //Vector4 - Vector4
	return Vector4{v[X] - other[X], v[Y] - other[Y], v[Z] - other[Z], v[W] - other[W]}
}
func (v Vector4) Mul(other Vector4) Vector4 { //Vector4 * Vector4
	return Vector4{v[X] * other[X], v[Y] * other[Y], v[Z] * other[Z], v[W] * other[W]}
}
func (v Vector4) Div(other Vector4) Vector4 { //Vector4 / Vector4
	return Vector4{v[X] / other[X], v[Y] / other[Y], v[Z] / other[Z], v[W] / other[W]}
}
func (v Vector4) Addf(other float64) Vector4 { //Vector4 + float
	return Vector4{v[X] + float(other), v[Y] + float(other), v[Z] + float(other), v[W] + float(other)}
}
func (v Vector4) Subf(other float64) Vector4 { //Vector4 - float
	return Vector4{v[X] - float(other), v[Y] - float(other), v[Z] - float(other), v[W] - float(other)}
}
func (v Vector4) Mulf(other float64) Vector4 { //Vector4 * float
	return Vector4{v[X] * float(other), v[Y] * float(other), v[Z] * float(other), v[W] * float(other)}
}
func (v Vector4) Divf(other float64) Vector4 { //Vector4 / float
	return Vector4{v[X] / float(other), v[Y] / float(other), v[Z] / float(other), v[W] / float(other)}
}
func (v Vector4) Neg() Vector4 { return Vector4{-v[X], -v[Y], -v[Z], -v[W]} }

// Transform transform transforms (multiplies) the [Vector4] by the given [Projection]'s transformation matrix.
func (v Vector4) Transform(p Projection) Vector4 { //Projection * Vector4
	return Vector4{
		p[0][0]*v[X] + p[1][0]*v[Y] + p[2][0]*v[Z] + p[3][0]*v[W],
		p[0][1]*v[X] + p[1][1]*v[Y] + p[2][1]*v[Z] + p[3][1]*v[W],
		p[0][2]*v[X] + p[1][2]*v[Y] + p[2][2]*v[Z] + p[3][2]*v[W],
		p[0][3]*v[X] + p[1][3]*v[Y] + p[2][3]*v[Z] + p[3][3]*v[W]}
}
