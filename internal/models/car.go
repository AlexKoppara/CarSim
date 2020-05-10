package models

// Car struct
type Car struct {
	Make  string
	Model string
	Lat   float64
	Long  float64
	Vx    float64
	Vy    float64
	Acc   float64
}

// Drive => position = velocty * time
func (c *Car) Drive(seconds float64) {
	c.Lat += c.Vy * seconds
	c.Long += c.Vx * seconds
	return
}

// Turn => new vx & vy
func (c *Car) Turn(vx float64, vy float64) {
	c.Vx = vx
	c.Vy = vy
	return
}
