// Package frvector - A collection of Vector routines.
package frvector

import "math"

// Vec3 - It must be encapsulated in a struct to hide from the caller.
type Vec3 struct {
	c [3]float64
}

// Set - Sets a Vec3 given 3 floats.
func (v *Vec3) Set(x, y, z float64) {
	v.c[0] = x
	v.c[1] = y
	v.c[2] = z
}

// Get - Returns the Vec3.
func (v *Vec3) Get() Vec3 {
	return *v
}

// GetComponent - Returns the i'th component of the Vec3 origin.
func (v *Vec3) GetComponent(i int) float64 {
	return v.c[i]
}

// Normalize - Normalize the i'th component of the Vec3 direction.
func (v *Vec3) Normalize() float64 {
	var d = math.Sqrt(v.c[0]*v.c[0] + v.c[1]*v.c[1] + v.c[2]*v.c[2])
	v.c[0] /= d
	v.c[1] /= d
	v.c[2] /= d
	return d
}

// Scale - Scale the origin by 's' along the direction.
func (v *Vec3) Scale(s float64) {
	v.c[0] = v.c[0] * s
	v.c[1] = v.c[1] * s
	v.c[2] = v.c[2] * s
}
