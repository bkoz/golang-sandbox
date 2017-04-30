package main

import "fmt"
import "github.com/bkoz/golang-sandbox/frVector"

// Main - Testing local packages and built-in golang functions.
func main() {
	const (
		a int = iota
		b
	)

	var z frVector.Vec3
	z.Set(0.707, 1.414, 3.52)
	var localvec = z.Get()

	fmt.Println("rev. 1.1")
	fmt.Println("localvec = ", localvec)
	z.Scale(2)
	fmt.Println("scaled by 2 = ", z.Get())
	fmt.Println("localvec[1] = ", localvec.GetComponent(1))
	fmt.Println("distance = ", z.Normalize())
	localvec = z.Get()
	fmt.Println("normalized dir = ", localvec)

	// frimage.Write()
}
