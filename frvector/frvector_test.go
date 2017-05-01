package frvector

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	fmt.Println("TestNormalize")
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
