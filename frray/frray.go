package ray

import "github.com/bkoz/golang-sandbox/frvector"

// Ray - It must be encapsulated in a struct to hide from the caller.
type Ray struct {
	o frvector.Vec3
	d frvector.Vec3
}

// SetOrigin - Sets a Ray's origin given 3 floats.
func (r *Ray) SetOrigin(x, y, z float64) {
	r.o.Set(x, y, z)
}

// GetOrigin - Returns the origin of the ray.
func (r *Ray) GetOrigin() frvector.Vec3 {
	return r.o
}

// SetDirection - Sets a Ray's Direction given 3 floats.
func (r *Ray) SetDirection(x, y, z float64) {
	r.o.Set(x, y, z)
}

// GetDirection - Returns the Direction of the ray.
func (r *Ray) GetDirection() frvector.Vec3 {
	return r.o
}
