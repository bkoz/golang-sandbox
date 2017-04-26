package main

// import "fmt"
import "github.com/bkoz/golang-sandbox/frimage"

func main() {
	const (
		a int = iota
		b
	)

	// fmt.Errorf("error: %d", a)
	// fmt.Println("b = ", b)

	frimage.Write()
}
