// Package frray_test
package frray

import (
	"testing"
)

// TestGetOrigin
func TestGetOrigin(t *testing.T) {
	var r Ray
	r.SetOrigin(1, 1, 1)
	if r.Get()[0] == 1 && r.Get()[1] == 1 && r.Get()[2] == 1 {

	} else {
		t.Error(`GetOrigin() = false`)
	}
}