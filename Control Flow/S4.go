package main

import (
	"fmt"
)

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
	ex8()
}

func ex1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// ASCII
func ex2() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U\n", i)
	}
}

func ex3() {
	bd := 2001
	for bd <= 2020 {
		fmt.Printf("%d\n", bd)
		bd++
	}

}

func ex4() {
	bd := 2001
	fmt.Println("for ;;")
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
func ex5() {
	for i := 4; i <= 100; i++ {
		fmt.Println("modulus of", i, "by 4", i%4)
	}
}

func ex6() {
	x := 4
	if x == 3 {
		fmt.Println("yes")
	} else if x == 4 {
		fmt.Println("woohoo")
	} else {
		fmt.Println("wtf")
	}
}

func ex7() {
	switch {
	// default go true
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}

func ex8() {
	x:=2
	switch x {
	case 2:
		fmt.Println("2")
		fallthrough // next case
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("no")
	}
}
