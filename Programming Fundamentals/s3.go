package main

import "fmt"

const (
	a     = 42
	b int = 69
)

// iota
const (
	d = 2020 + iota
	e = 2020 + iota
	f = 2020 + iota
	g = 2020 + iota
)
func main() {
	x := a < b
	y:= 69
	z:= y<<1

	// raw string literal
	c:= `here is something
	that matters
	full foramtted
	`
	fmt.Println(x)
	fmt.Printf("%d\t%b\t%#x",z, z, z)
	fmt.Printf(c)
	fmt.Println(d, e, f, g)
}
