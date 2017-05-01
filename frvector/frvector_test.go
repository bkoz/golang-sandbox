package frvector

import "testing"

func TestNormalize(t *testing.T) {
	v := Vec3
	v.Set(1, 1, 1)
	if !Normalize(v) {
		fmt.Println("testing...")
	}
}
