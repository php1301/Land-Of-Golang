package main

import (
	"fmt"

	"rsc.io/quote"
)
//var is package scope
var y = 42
var z string = "Hello Phuc"
var a int = 42
type hotdog int
var b hotdog

var n int = 42
var m string = "var declaration"
var o bool = true

type ex3 int
var c ex3
var d int
func main() {
	fmt.Println(quote.Go())
	// foo()
	// fmt.Println("this worked")
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// n, e := fmt.Println("assigned")
	// d, _ := fmt.Println("use underscore")
	// fmt.Println(n)
	// fmt.Println(e)
	// fmt.Println(d)
	// fmt.Println("y global scope", y)
	// fmt.Printf("%T\n", z)
	// a = 42
	// b = 43
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// a = int(b)
	// identifiers()
	// ex1()
	d= 69
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	c = 96
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	
	d = int(c)

	s:= fmt.Sprintf("%v\t%v\t%v\t",n,m,o)
	fmt.Println(s)
}
func foo() {
	fmt.Println("ok")
}

func identifiers() {
	x:= 2 //declare
	x = 3 // reassgin
	fmt.Println(x)
}

func ex1(){
	x:=42
	y:="abc"
	z:=true

	fmt.Println(x, y, z)
}