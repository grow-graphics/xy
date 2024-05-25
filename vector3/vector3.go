package vector3

import "grow.graphics/xy"

func New(x, y, z float64) xy.Vector3                      { return xy.NewVector3(x, y, z) }
func Abs(v xy.Vector3) xy.Vector3                         { return v.Abs() }
func Ceil(v xy.Vector3) xy.Vector3                        { return v.Ceil() }
func Clamp(v xy.Vector3, min, max xy.Vector3) xy.Vector3  { return v.Clamp(min, max) }
func Cross(v, with xy.Vector3) xy.Vector3                 { return v.Cross(with) }
func Distance(v, to xy.Vector3) float64                   { return v.DistanceTo(to) }
func DistanceSquared(v, to xy.Vector3) float64            { return v.DistanceSquaredTo(to) }
func Dot(v, with xy.Vector3) float64                      { return v.Dot(with) }
func Floor(v xy.Vector3) xy.Vector3                       { return v.Floor() }
func IsFinite(v xy.Vector3) bool                          { return v.IsFinite() }
func IsNormalized(v xy.Vector3) bool                      { return v.IsNormalized() }
func IsApproximatelyZero(v xy.Vector3) bool               { return v.IsApproximatelyZero() }
func Length(v xy.Vector3) float64                         { return v.Length() }
func LengthSquared(v xy.Vector3) float64                  { return v.LengthSquared() }
func Lerp(from, to xy.Vector3, weight float64) xy.Vector3 { return from.Lerp(to, weight) }
func Normalize(v xy.Vector3) xy.Vector3                   { return v.Normalized() }
func Round(v xy.Vector3) xy.Vector3                       { return v.Round() }

func Add(v, with xy.Vector3) xy.Vector3 { return v.Add(with) }
func Sub(v, with xy.Vector3) xy.Vector3 { return v.Sub(with) }
func Mul(v, with xy.Vector3) xy.Vector3 { return v.Mul(with) }
func Div(v, with xy.Vector3) xy.Vector3 { return v.Div(with) }

func Addf(v xy.Vector3, with float64) xy.Vector3 { return v.Addf(with) }
func Subf(v xy.Vector3, with float64) xy.Vector3 { return v.Subf(with) }
func Mulf(v xy.Vector3, with float64) xy.Vector3 { return v.Mulf(with) }
func Divf(v xy.Vector3, with float64) xy.Vector3 { return v.Divf(with) }

func Neg(v xy.Vector3) xy.Vector3 { return v.Neg() }

func Transform(v xy.Vector3, by xy.Transform3D) xy.Vector3 { return v.Transform(by) }
