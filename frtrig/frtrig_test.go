// Package frtrig
package frtrig

import "testing"
import "fmt"

// TestCalc
func TestCalc(t *testing.T) {
	p := new(Pi)
	p.SetPrecision(6)
	p.Calc()
	fmt.Println("pi = ", p.Get())
	if p.Get() == 3.1415916535897934 {
	} else {
		t.Error()
	}
}

// TestSet
func TestSetPrecision(t *testing.T) {
	p := new(Pi)
	p.SetPrecision(7)
	fmt.Println("pi.precision = ", p.GetPrecision())
	if p.GetPrecision() == 10000000 {
	} else {
		t.Error()
	}
}
