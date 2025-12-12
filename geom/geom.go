package geom

import "math"

// pi radian = 180 deg
func DegreeToRadian(degree float32) float32 {
	return degree * (math.Pi / 180)
}

func RadianToDegree(radian float32) float32 {
	return radian * (180 / math.Pi)
}
