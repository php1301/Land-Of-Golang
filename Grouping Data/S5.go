package main

import "fmt"

func main() {
	// ex1()
	// ex2()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
	// ex8()
	// ex9()
	ex10()
}

func ex1() {
	a := [5]int{42, 43, 44, 45, 46}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
func ex2() {
	a := []int{42, 43, 44, 45, 46}
	for i, v := range a {
		fmt.Println(i, v)
	}
	// ex3
	fmt.Println(a[:4])
	fmt.Println(a[1:3])
}

func ex4() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50}
	b := append(a, 52)
	fmt.Println(b)
	c := append(b, 53, 54, 55)
	fmt.Println(c)
	x := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println("after appending")
	d := append(c, x...)
	fmt.Println(d)
}

func ex5() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	b := a[:3]
	c := a[6:10]
	d := append(b, c...)
	fmt.Println(d)
}

func ex6() {
	states := make([]string, 49, 100) // allocation slice, 53 length 100 capacity
	states = []string{
		`123`, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}

func ex7() {
	x := make([]string, 50, 50)
	x = []string{"abc", "def"}
	y := []string{`xyz`, `ghi`}
	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println("record", i)
		for j, t := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, t)
		}
	}
}

func ex8() {
	m := map[string][]string{
		"name": {"Phuc", "Hoang", "Pham"},
		"sur":  {"Great", "Cool", "Fun"},
		"sub":  {"Attractive", "Handsome", "DIs"},
	}

	for k, v := range m {
		fmt.Println("record", k)
		for j, v2 := range v {
			fmt.Printf("Index: %v \t value: %v \t \n", j, v2)
		}
	}
}

// add new key to map
func ex9() {
	m := map[string][]string{
		"name": {"Phuc", "Hoang", "Pham"},
		"sur":  {"Great", "Cool", "Fun"},
		"sub":  {"Attractive", "Handsome", "DIs"},
	}
	m["new"] = []string{"New 1", "New 2", "New 3"}
	for k, v := range m {
		fmt.Println("record", k)
		for j, v2 := range v {
			fmt.Printf("Index: %v \t value: %v \t \n", j, v2)
		}
	}
}

// delete key from map 
func ex10() {
	m := map[string][]string{
		"name": {"Phuc", "Hoang", "Pham"},
		"sur":  {"Great", "Cool", "Fun"},
		"sub":  {"Attractive", "Handsome", "DIs"},
	}
	delete(m, "sur")
	for k, v := range m {
		fmt.Println("record", k)
		for j, v2 := range v {
			fmt.Printf("Index: %v \t value: %v \t \n", j, v2)
		}
	}
}
