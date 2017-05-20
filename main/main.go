package main

import "fmt"
import "github.com/bkoz/golang-sandbox/frvector"
import "github.com/bkoz/golang-sandbox/frtrig"

// Console -
type Console struct {
	X int
	Y int
}

// NewConsole - Init console members
func NewConsole() *Console {
	return &Console{X: 5, Y: 10}
}

// Main - Testing local packages and built-in golang functions.
func main() {
	const (
		a int = iota
		b
	)

	var z frvector.Vec3
	z.Set(1, 1, 1)
	var u frvector.Vec3
	u.Set(1, 1, 1)
	var localvec = z.Get()

	fmt.Println("localvec = ", localvec)
	fmt.Println("localvec[1] = ", localvec[1])
	fmt.Println("distance = ", z.Normalize())
	fmt.Println("normalized = ", z.Get())
	z.Scale(2, u)
	fmt.Println("scaled by 2 = ", z.Get())

	// Call Pi
	var p frtrig.Pi
	p.Init()
	p.Calc()
	fmt.Println("pi = ", p.Get())

	var c = *NewConsole()

	fmt.Println("console = ", c)

	// frimage.Write()
}
