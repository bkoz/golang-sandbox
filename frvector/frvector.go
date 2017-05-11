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

// Get - Returns the 3D array of floats.
func (v *Vec3) Get() [3]float64 {
	return v.c
}

// Normalize - Normalize the i'th component of the Vec3 direction.
func (v *Vec3) Normalize() float64 {
	var d = math.Sqrt(v.c[0]*v.c[0] + v.c[1]*v.c[1] + v.c[2]*v.c[2])
	v.c[0] /= d
	v.c[1] /= d
	v.c[2] /= d
	return d
}

// Scale - Scale this vector by s in the direction of v2.
func (v *Vec3) Scale(s float64, v2 Vec3) {
	v.c[0] += v2.c[0] * s
	v.c[1] += v2.c[1] * s
	v.c[2] += v2.c[2] * s
}

// Add - Add v2 to this vector.
func (v *Vec3) Add(v2 Vec3) {
	v.c[0] += v2.c[0]
	v.c[1] += v2.c[1]
	v.c[2] += v2.c[2]
}

// Subtract - Subtract v2 from this vector.
func (v *Vec3) Subtract(v2 Vec3) {
	v.c[0] -= v2.c[0]
	v.c[1] -= v2.c[1]
	v.c[2] -= v2.c[2]
}

// Dot - Calculate and return the dot product.
func (v *Vec3) Dot(v2 Vec3) float64 {
	return (v.c[0] * v2.c[0]) + (v.c[1] * v2.c[1]) + (v.c[2] * v2.c[2])
}

// Cross - Calculate and return the cross product.
func (v *Vec3) Cross(v2 Vec3) Vec3 {
	var c Vec3
	c.c[0] = v.c[1]*v2.c[2] - v.c[2]*v2.c[1]
	c.c[1] = v.c[2]*v2.c[0] - v.c[0]*v2.c[2]
	c.c[2] = v.c[0]*v2.c[1] - v.c[1]*v2.c[0]
	return c
}
