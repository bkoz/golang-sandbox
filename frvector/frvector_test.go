// Package frvector
package frvector

import (
	"testing"
)

// TestGetC
func TestGet(t *testing.T) {
	var v Vec3
	v.Set(1, 1, 1)
	if v.Get()[0] == 1 && v.Get()[1] == 1 && v.Get()[2] == 1 {

	} else {
		t.Error(`Get() = false`)
	}
}

// TestNormalize
func TestNormalize(t *testing.T) {
	var v Vec3
	v.Set(1, 1, 1)
	v.Normalize()
	if v.Get()[0] == 0.5773502691896258 &&
		v.Get()[1] == 0.5773502691896258 &&
		v.Get()[2] == 0.5773502691896258 {
	} else {
		t.Error(`Normalize() = false`)
	}
}

// TestScale
func TestScale(t *testing.T) {
	var v, v2 Vec3
	v.Set(1, 1, 1)
	v2.Set(1, 1, 1)
	v.Scale(2, v2)
	if v.Get()[0] == 3 &&
		v.Get()[1] == 3 &&
		v.Get()[2] == 3 {
	} else {
		t.Error(`Scale() = false`)
	}
}

// TestAdd
func TestAdd(t *testing.T) {
	var v, v2 Vec3
	v.Set(1, 1, 1)
	v2.Set(1, 1, 1)
	v.Add(v2)
	if v.Get()[0] == 2 &&
		v.Get()[1] == 2 &&
		v.Get()[2] == 2 {
	} else {
		t.Error(`Add() = false`)
	}
}

// TestSubtract
func TestSubtract(t *testing.T) {
	var v, v2 Vec3
	v.Set(1, 1, 1)
	v2.Set(1, 1, 1)
	v.Subtract(v2)
	if v.Get()[0] == 0 &&
		v.Get()[1] == 0 &&
		v.Get()[2] == 0 {
	} else {
		t.Error(`Subtract() = false`)
	}
}

// TestDot
func TestDot(t *testing.T) {
	var v, v2 Vec3
	v.Set(1, 1, 1)
	v2.Set(1, 1, 1)
	if v.Dot(v2) == 3 {
	} else {
		t.Error(`Dot() = false`)
	}
}

// TestCross
func TestCross(t *testing.T) {
	var v, v2, v3 Vec3
	v.Set(1, 0, 0)
	v2.Set(0, 1, 0)
	v3 = v.Cross(v2)

	if v3.Get()[0] == 0 &&
		v3.Get()[1] == 0 &&
		v3.Get()[2] == 1 {
	} else {
		t.Error(`Cross() = false`)
	}

}
