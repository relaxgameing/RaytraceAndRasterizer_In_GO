package geom

type Ray struct {
	Point           WorldPoint
	Lambda          float32
	DirectionVector Vector
}

func (r *Ray) GetPointOnRay() *WorldPoint {
	return &WorldPoint{
		r.Point.X + r.Lambda*r.DirectionVector.X,
		r.Point.Y + r.Lambda*r.DirectionVector.Y,
		r.Point.Z + r.Lambda*r.DirectionVector.Z,
	}
}

func (r *Ray) GetPointOnRayWithLambda(lambda float32) *WorldPoint {
	return &WorldPoint{
		r.Point.X + lambda*r.DirectionVector.X,
		r.Point.Y + lambda*r.DirectionVector.Y,
		r.Point.Z + lambda*r.DirectionVector.Z,
	}
}
