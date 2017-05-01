// Package frvector
package frvector

import (
	"testing"
)

// TestGetC
func TestGetC(t *testing.T) {
	var v Vec3
	v.Set(1, 1, 1)
	if v.GetC()[0] == 1 && v.GetC()[1] == 1 && v.GetC()[2] == 1 {

	} else {
		t.Error(`GetC() = false`)
	}
}

// TestNormalize
func TestNormalize(t *testing.T) {
	var v Vec3
	v.Set(1, 1, 1)
	v.Normalize()
	if v.GetComponent(0) == 0.5773502691896258 &&
		v.GetComponent(1) == 0.5773502691896258 &&
		v.GetComponent(2) == 0.5773502691896258 {
	} else {
		t.Error(`Normalize() = false`)
	}
}
