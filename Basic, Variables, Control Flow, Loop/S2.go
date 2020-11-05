package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	foo()
	fmt.Println("this worked")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	n, e := fmt.Println("assigned")
	d, _ := fmt.Println("use underscore")
	fmt.Println(n)
	fmt.Println(e)
	fmt.Println(d)
	identifiers()
}
func foo() {
	fmt.Println("ok")
}

func identifiers() {
	x:= 2 //declare
	x = 3 // reassgin
	fmt.Println(x)
}