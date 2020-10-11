package loc

// Coordinate defines GPS coordinate in degree unit.
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *Coordinate) distance(o *Coordinate) float64 {
	return Distance(c.Latitude, c.Longitude, o.Latitude, o.Longitude)
}
