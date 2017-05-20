package main

import "fmt"
import "github.com/bkoz/golang-sandbox/frvector"
import "github.com/bkoz/golang-sandbox/frtrig"

// Main - Testing local packages and built-in golang functions.
func main() {
	const (
		a int = iota
		b
	)

	var z frvector.Vec3
	z.Set(1, 1, 1)
	var localvec = z.Get()

	fmt.Println("localvec = ", localvec)
	// z.Scale(2)
	// fmt.Println("scaled by 2 = ", z.Get())
	fmt.Println("localvec[1] = ", localvec[1])
	fmt.Println("distance = ", z.Normalize())
	fmt.Println("normalized = ", z.Get())

	// Call Pi
	var p frtrig.Pi
	p.Init()
	p.Calc()
	fmt.Println("pi = ", p.Get())

	// frimage.Write()
}
