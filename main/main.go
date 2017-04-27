package main

import "fmt"
import "github.com/bkoz/golang-sandbox/frvector"

func main() {
	const (
		a int = iota
		b
	)

	var z frvector.VecStruct

	z.Set(0.707, 1.414, 2.52)
	var localvec = z.Get()

	fmt.Println("localvec = ", localvec[0], localvec[1], localvec[2])

	// frimage.Write()
}
