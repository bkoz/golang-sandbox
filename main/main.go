package main

import "fmt"
import "github.com/bkoz/golang-sandbox/frvector"

// Main - Testing local packages and built-in golang functions.
func main() {
	const (
		a int = iota
		b
	)

	var z frvector.Vector

	z.Set(0, 0, 0, 0.707, 1.414, 3.52)
	var localvec = z.Get()

	fmt.Println("localvec = ", localvec)
	z.Scale(2)
	fmt.Println("scaled = ", z.Get())
	fmt.Println("localvec origin [1] = ", localvec.GetOrigin(1))
	fmt.Println("localvec dir [1] = ", localvec.GetDirection(1))
	fmt.Println("distance = ", z.Normalize())
	localvec = z.Get()
	fmt.Println("normalized dir [1] = ", localvec)

	// frimage.Write()
}
