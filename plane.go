package xy

/*
Plane represents a normalized plane equation. normal is the normal of the plane (a, b, c normalized), and d is the
distance from the origin to the plane (in the direction of "normal"). "Over" or "Above" the plane is considered the
side of the plane towards where the normal is pointing.
*/
type Plane struct {
	/*
		Normal of the plane, typically a unit vector. Shouldn't be a zero vector as Plane with such normal does not
		represent a valid plane.

		In the scalar equation of the plane ax + by + cz = d, this is the vector (a, b, c), where d is the d property.
	*/
	Normal Vector3
	/*
		D is the distance from the origin to the plane, expressed in terms of normal (according to its direction and magnitude).
		Actual absolute distance from the origin to the plane can be calculated as Abs(d) / normal.Length() (if normal has zero
		length then this Plane does not represent a valid plane).

		In the scalar equation of the plane ax + by + cz = d, this is d, while the (a, b, c) coordinates are represented by the
		normal property.
	*/
	D float
}

// NewPlane creates a plane from the three points, given in clockwise order.
func NewPlane(a, b, c Vector3) Plane {
	normal := a.Sub(c).Cross(a.Sub(b)).Normalized()
	return Plane{
		Normal: normal,
		D:      float(normal.Dot(a)),
	}
}

// "Fields"

func (p Plane) X() float64 { return float64(p.Normal[0]) }
func (p Plane) Y() float64 { return float64(p.Normal[1]) }
func (p Plane) Z() float64 { return float64(p.Normal[2]) }

func (p *Plane) SetX(x float64) { p.Normal[0] = float(x) }
func (p *Plane) SetY(y float64) { p.Normal[1] = float(y) }
func (p *Plane) SetZ(z float64) { p.Normal[2] = float(z) }

// "Constants"

func (Plane) YZ() Plane { return Plane{Normal: Vector3{1, 0, 0}} }
func (Plane) XZ() Plane { return Plane{Normal: Vector3{0, 1, 0}} }
func (Plane) XY() Plane { return Plane{Normal: Vector3{0, 0, 1}} }

// "Methods"

// DistanceTo returns the shortest distance from the plane to the position point.
// If the point is above the plane, the distance will be positive. If below,
// the distance will be negative.
func (p Plane) DistanceTo(point Vector3) float64 { //Plane.distance_to
	return p.Normal.Dot(point) - float64(p.D)
}

// GetCenter returns the center of the plane.
func (p Plane) Center() Vector3 { //Plane.get_center
	return p.Normal.Mulf(float64(p.D))
}

// HasPoint returns true if point is inside the plane. Comparison uses a custom minimum
// tolerance threshold.
func (p Plane) HasPoint(point Vector3, tolerance float64) bool { //Plane.has_point
	if tolerance < cmpEpsilon {
		tolerance = cmpEpsilon
	}
	return Absf(p.Normal.Dot(point)-float64(p.D)) <= tolerance
}

// Intersect3 returns the intersection point of the three planes b, c and this plane. If
// no intersection is found, false is returned.
func (p Plane) Intersect3(b, c Plane) (Vector3, bool) { //Plane.intersect_3
	var normal0 = p.Normal
	var normal1 = b.Normal
	var normal2 = c.Normal
	var (
		denom = normal0.Cross(normal1).Dot(normal2)
	)
	if IsApproximatelyZero(denom) {
		return Vector3{}, false
	}
	return normal1.Cross(normal2).Mulf(float64(p.D)).Add(normal2.Cross(normal0).Mulf(float64(b.D))).Add(normal0.Cross(normal1).Mulf(float64(c.D))).Divf(denom), true
}

// IntersectsRay returns the intersection point of a ray consisting of the position from and the
// direction normal dir with this plane. If no intersection is found, false is returned.
func (p Plane) IntersectsRay(from, dir Vector3) (Vector3, bool) { //Plane.intersects_ray
	var segment = dir
	var den = p.Normal.Dot(segment)
	if IsApproximatelyZero(den) {
		return Vector3{}, false
	}
	var dist = (p.Normal.Dot(from) - float64(p.D)) / den
	if dist > cmpEpsilon { //this is a ray, before the emitting pos (p_from) doesn't exist
		return Vector3{}, false
	}
	dist = -dist
	return from.Add(segment.Mulf(dist)), true
}

// IntersectsSegment returns the intersection point of a segment from position from to position to
// with this plane. If no intersection is found, false is returned.
func (p Plane) IntersectsSegment(from, to Vector3) (Vector3, bool) { //Plane.intersects_segment
	var segment = to.Sub(from)
	var den = p.Normal.Dot(segment)
	if IsApproximatelyZero(den) {
		return Vector3{}, false
	}
	var dist = (p.Normal.Dot(from) - float64(p.D)) / den
	if dist < -cmpEpsilon || dist > 1+cmpEpsilon {
		return Vector3{}, false
	}
	return from.Add(segment.Mulf(dist)), true
}

// IsApproximatelyEqual returns true if this plane and other are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func (p Plane) IsApproximatelyEqual(other Plane) bool { //Plane.is_equal_approx
	return p.Normal.IsApproximatelyEqual(other.Normal) && IsApproximatelyEqual(p.D, other.D)
}

// IsFinite returns true if this plane is finite, by calling [IsFinite] on each component.
func (p Plane) IsFinite() bool { return p.Normal.IsFinite() && IsFinite(p.D) } //Plane.is_finite

// IsPointOver returns true if point is located above the plane.
func (p Plane) IsPointOver(point Vector3) bool { //Plane.is_point_over
	return p.Normal.Dot(point) > float64(p.D)
}

// Normalized returns a copy of the plane, with normalized normal (so it's a unit vector).
// Returns Plane{0, 0, 0, 0} if normal can't be normalized (it has zero length).
func (p Plane) Normalized() Plane { //Plane.normalized
	var l = p.Normal.Length()
	if l == 0 {
		return Plane{}
	}
	return Plane{
		Normal: p.Normal.Divf(l),
		D:      p.D / float(l),
	}
}

// Project returns the orthogonal projection of point into a point in the plane.
func (p Plane) Project(point Vector3) Vector3 { //Plane.project
	return point.Sub(p.Normal.Mulf(p.DistanceTo(point)))
}

// Transform transforms (multiplies) the [Plane] by the given [Transform3D] transformation matrix.
func (p Plane) Transform(t Transform3D) Plane { //Transform3D * Plane
	b := t.Basis.Inverse().Transposed()

	// Transform a single point on the plane.
	var point = p.Normal.Mulf(float64(p.D))
	point.Transform(t)

	// Use inverse transpose for correct normals with non-uniform scaling.
	var normal = b.Transform(p.Normal).Normalized()

	var d = p.Normal.Dot(point)
	return Plane{normal, float(d)}
}
