package xy

import "math"

/*
Vector4i is a 4-element structure that can be used to represent 4D grid coordinates or any other quadruplet of integers.

It uses integer coordinates and is therefore preferable to [Vector4] when exact precision is required. Note that the
values are limited to 32 bits, and unlike Vector4 this cannot be configured with an engine build option. Use [Int] or
[PackedInt64Array] if 64-bit values are needed.

Note: In a boolean context, a [Vector4i] will evaluate to false if it's equal to Vector4i{0, 0, 0, 0}. Otherwise, a
[Vector3i] will always evaluate to true.
*/
type Vector4i [4]int32

// "Fields"

func (v Vector4i) X() int64 { return int64(v[X]) }
func (v Vector4i) Y() int64 { return int64(v[Y]) }
func (v Vector4i) Z() int64 { return int64(v[Z]) }
func (v Vector4i) W() int64 { return int64(v[W]) }

func (v *Vector4i) SetX(x int64) { v[X] = int32(x) }
func (v *Vector4i) SetY(y int64) { v[Y] = int32(y) }
func (v *Vector4i) SetZ(z int64) { v[Z] = int32(z) }
func (v *Vector4i) SetW(w int64) { v[W] = int32(w) }

// "Constants"

func (Vector4i) ZERO() Vector4i { return Vector4i{} }
func (Vector4i) ONE() Vector4i  { return Vector4i{1, 1, 1, 1} }
func (Vector4i) MIN() Vector4i {
	return Vector4i{math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32}
}
func (Vector4i) MAX() Vector4i {
	return Vector4i{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
}

func (v Vector4i) Vector4() Vector4 { //Vector4(Vector4i)
	return Vector4{float(v[X]), float(v[Y]), float(v[Z]), float(v[W])}
}

// "Methods"

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector4i) Abs() Vector4i { return v.abs() } //Vector4i.abs

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Clampi] on each component.
func (v Vector4i) Clamp(min, max Vector4i) Vector4i { //Vector4i.clamp
	return Vector4i{
		Clampi(v[X], min[X], max[X]),
		Clampi(v[Y], min[Y], max[Y]),
		Clampi(v[Z], min[Z], max[Z]),
		Clampi(v[W], min[W], max[W]),
	}
}

// Length returns the length (magnitude) of this vector.
func (v Vector4i) Length() float64 { return Sqrt(float64(v.LengthSquared())) } //Vector4i.length

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than [Vector4i.Length], so prefer it if you need to
// compare vectors or need the squared distance for some formula.
func (v Vector4i) LengthSquared() int64 { //Vector4i.length_squared
	return int64(v[X])*int64(v[X]) + int64(v[Y])*int64(v[Y]) + int64(v[Z])*int64(v[Z]) + int64(v[W])*int64(v[W])
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func (v Vector4i) MaxAxis() Axis { //Vector4i.max_axis_index
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
func (v Vector4i) MinAxis() Axis { //Vector4i.min_axis_index
	var min_index = 0
	var min_value = v[W]
	for i := 1; i < len(v); i++ {
		if v[i] < min_value {
			min_index = i
			min_value = v[i]
		}
	}
	return Axis(min_index)
}

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's negative, and 0 if
// it's 0. The result is identical to running [Signi] on each component.
func (v Vector4i) Sign() Vector4i { return v.sign() } //Vector4i.sign

// Snapped returns a new vector with each component snapped to the nearest multiple of step.
func (v Vector4i) Snapped(step Vector4i) Vector4i { //Vector4i.snapped
	return Vector4i{
		Snappedi(v[X], step[X]),
		Snappedi(v[Y], step[Y]),
		Snappedi(v[Z], step[Z]),
		Snappedi(v[W], step[W]),
	}
}

// "Operators"

func (v Vector4i) Add(other Vector4i) Vector4i { //Vector4i + Vector4i
	return Vector4i{v[X] + other[X], v[Y] + other[Y], v[Z] + other[Z], v[W] + other[W]}
}
func (v Vector4i) Sub(other Vector4i) Vector4i { //Vector4i - Vector4i
	return Vector4i{v[X] - other[X], v[Y] - other[Y], v[Z] - other[Z], v[W] - other[W]}
}
func (v Vector4i) Mul(other Vector4i) Vector4i { //Vector4i * Vector4i
	return Vector4i{v[X] * other[X], v[Y] * other[Y], v[Z] * other[Z], v[W] * other[W]}
}
func (v Vector4i) Div(other Vector4i) Vector4i { //Vector4i / Vector4i
	return Vector4i{v[X] / other[X], v[Y] / other[Y], v[Z] / other[Z], v[W] / other[W]}
}
func (v Vector4i) Mod(other Vector4i) Vector4i { //Vector4i % Vector4i
	return Vector4i{v[X] % other[X], v[Y] % other[Y], v[Z] % other[Z], v[W] % other[W]}
}
func (v Vector4i) Addi(other int64) Vector4i { //Vector4i + int64
	return Vector4i{v[X] + int32(other), v[Y] + int32(other), v[Z] + int32(other), v[W] + int32(other)}
}
func (v Vector4i) Subi(other int64) Vector4i { //Vector4i - int64
	return Vector4i{v[X] - int32(other), v[Y] - int32(other), v[Z] - int32(other), v[W] - int32(other)}
}
func (v Vector4i) Muli(other int64) Vector4i { //Vector4i * int64
	return Vector4i{v[X] * int32(other), v[Y] * int32(other), v[Z] * int32(other), v[W] * int32(other)}
}
func (v Vector4i) Divi(other int64) Vector4i { //Vector4i / int64
	return Vector4i{v[X] / int32(other), v[Y] / int32(other), v[Z] / int32(other), v[W] / int32(other)}
}
func (v Vector4i) Modi(other int64) Vector4i { //Vector4i % int64
	return Vector4i{v[X] % int32(other), v[Y] % int32(other), v[Z] % int32(other), v[W] % int32(other)}
}
func (v Vector4i) Neg() Vector4i { return Vector4i{-v[X], -v[Y], -v[Z], -v[W]} } //-Vector4i
