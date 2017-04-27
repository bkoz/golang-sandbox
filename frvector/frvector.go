// Package frvector - A collection of vector routines.
package frvector

// Vector - A private 3D array of float64.
type Vector [3]float64

// VecStruct - It should probably be called Vector, not VecStruct. It
// must be encapsulated in a struct to hide from the caller.
type VecStruct struct {
	vec Vector
}

// Set - Sets a vector given 3 floats.
func (v *VecStruct) Set(x, y, z float64) {
	v.vec[0] = x
	v.vec[1] = y
	v.vec[2] = z
}

// Get - returns the vector.
func (v *VecStruct) Get() Vector {
	return v.vec
}
