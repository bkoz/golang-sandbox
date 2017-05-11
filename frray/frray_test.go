// Package frray
package frray

import (
	"testing"
)

// TestSetGet
func TestSetGet(t *testing.T) {
	var r Ray
	r.SetOrigin(1, 1, 1)
	r.SetDirection(1, 1, 1)
	if r.o.Get()[0] == 1 && r.o.Get()[1] == 1 && r.o.Get()[2] == 1 &&
		r.d.Get()[0] == 1 && r.d.Get()[1] == 1 && r.d.Get()[2] == 1 {

	} else {
		t.Error(`SetGet() = false`)
	}
}
