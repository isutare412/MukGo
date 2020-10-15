package common

import "math"

// Coordinate defines GPS coordinate in degree unit.
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Distance returns distance between 2 coordinates in meters.
func (c *Coordinate) Distance(o *Coordinate) float64 {
	return Distance(c.Latitude, c.Longitude, o.Latitude, o.Longitude)
}

// RangeSquare returns 2 points that are dist meter away from c.
func (c *Coordinate) RangeSquare(dist float64) (
	northWest, southEast *Coordinate,
) {
	northWest, southEast = &Coordinate{}, &Coordinate{}
	theta := (dist / earthRadius) * 180 / math.Pi

	northWest.Latitude = c.Latitude + theta
	northWest.Longitude = c.Longitude - theta
	southEast.Latitude = c.Latitude - theta
	southEast.Longitude = c.Longitude + theta
	return
}
