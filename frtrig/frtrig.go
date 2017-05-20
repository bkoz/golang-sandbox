// Package frtrig - Misc trig functions
package frtrig

import (
	"math"
)

// Pi - Define Pi
type Pi struct {
	p         float64
	precision uint
}

// Init - Set defaults.
func (pi *Pi) Init() {
	pi.SetPrecision(6)
}

// Get - Returns the value of pi.p.
func (pi *Pi) Get() float64 {
	return pi.p
}

// SetPrecision - Accuracy and compute time increases with x. 6 is a reasonable value.
func (pi *Pi) SetPrecision(x uint) {
	var exponent float64
	exponent = float64(x)
	var n float64
	n = math.Trunc(math.Pow(10, exponent))
	pi.precision = uint(n)
}

// GetPrecision - Get the precision.
func (pi *Pi) GetPrecision() uint {
	return pi.precision
}

// Calc - Calculate Pi, the slow way. This was written to busy up a cpu for performance testing.
func (pi *Pi) Calc() {

	pi.p = 0.0

	for i := pi.GetPrecision(); i > 0; i-- {
		var x float64
		var y float64
		x = float64(i + 1.0)
		y = float64(2.0*i - 1.0)
		// Calculate series in parenthesis
		pi.p = pi.p + math.Pow(-1.0, x)/y
		// When at last number in series, multiply by 4
		if i == 1 {
			pi.p = pi.p * 4.0
		}
	}
}
