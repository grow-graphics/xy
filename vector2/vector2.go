package vector2

import "grow.graphics/xy"

func New(x, y float64) xy.Vector2                         { return xy.NewVector2(x, y) }
func Abs(v xy.Vector2) xy.Vector2                         { return v.Abs() }
func Angle(v xy.Vector2) xy.Radians                       { return v.Angle() }
func Aspect(v xy.Vector2) float64                         { return v.Aspect() }
func Ceil(v xy.Vector2) xy.Vector2                        { return v.Ceil() }
func Clamp(v xy.Vector2, min, max xy.Vector2) xy.Vector2  { return v.Clamp(min, max) }
func Cross(v, with xy.Vector2) float64                    { return v.Cross(with) }
func Distance(v, to xy.Vector2) float64                   { return v.DistanceTo(to) }
func DistanceSquared(v, to xy.Vector2) float64            { return v.DistanceSquaredTo(to) }
func Dot(v, with xy.Vector2) float64                      { return v.Dot(with) }
func Floor(v xy.Vector2) xy.Vector2                       { return v.Floor() }
func IsFinite(v xy.Vector2) bool                          { return v.IsFinite() }
func IsNormalized(v xy.Vector2) bool                      { return v.IsNormalized() }
func IsApproximatelyZero(v xy.Vector2) bool               { return v.IsApproximatelyZero() }
func Length(v xy.Vector2) float64                         { return v.Length() }
func LengthSquared(v xy.Vector2) float64                  { return v.LengthSquared() }
func Lerp(from, to xy.Vector2, weight float64) xy.Vector2 { return from.Lerp(to, weight) }
func Normalize(v xy.Vector2) xy.Vector2                   { return v.Normalized() }
func Round(v xy.Vector2) xy.Vector2                       { return v.Round() }
func Sign(v xy.Vector2) xy.Vector2                        { return v.Sign() }

func Add(v, with xy.Vector2) xy.Vector2 { return v.Add(with) }
func Sub(v, with xy.Vector2) xy.Vector2 { return v.Sub(with) }
func Mul(v, with xy.Vector2) xy.Vector2 { return v.Mul(with) }
func Div(v, with xy.Vector2) xy.Vector2 { return v.Div(with) }

func Addf(v xy.Vector2, with float64) xy.Vector2 { return v.Addf(with) }
func Subf(v xy.Vector2, with float64) xy.Vector2 { return v.Subf(with) }
func Mulf(v xy.Vector2, with float64) xy.Vector2 { return v.Mulf(with) }
func Divf(v xy.Vector2, with float64) xy.Vector2 { return v.Divf(with) }

func Neg(v xy.Vector2) xy.Vector2 { return v.Neg() }

func Transform(v xy.Vector2, by xy.Transform2D) xy.Vector2 { return v.Transform(by) }
