// Package frvector - A collection of vector routines.
package frvector

import "math"

// Vector - It must be encapsulated in a struct to hide from the caller.
type Vector struct {
	o [3]float64
	d [3]float64
}

// Set - Sets a vector given 3 floats.
func (v *Vector) Set(ox, oy, oz, dx, dy, dz float64) {
	v.o[0] = ox
	v.o[1] = oy
	v.o[2] = oz
	v.d[0] = dx
	v.d[1] = dy
	v.d[2] = dz
}

// Get - Returns the vector.
func (v *Vector) Get() Vector {
	return *v
}

// GetOrigin - Returns the origin of the vector.
func (v *Vector) GetOrigin(i int) float64 {
	return v.o[i]
}

// GetDirection - Returns the origin of the vector.
func (v *Vector) GetDirection(i int) float64 {
	return v.d[i]
}

// Normalize - Normalize the direction of the vector.
func (v *Vector) Normalize() float64 {
	var d = math.Sqrt(v.d[0]*v.d[0] + v.d[1]*v.d[1] + v.d[2]*v.d[2])
	v.d[0] = v.d[0] / d
	v.d[1] = v.d[1] / d
	v.d[2] = v.d[2] / d
	return d
}
