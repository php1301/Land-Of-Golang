package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	fav       []string
}
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	fourWheel bool
}

func main() {
	// ex1()
	// ex2()
	// ex3()
	ex4()
}

func ex1() {
	p1 := person{
		firstname: "Phuc",
		lastname:  "Pham",
		fav:       []string{"Vanilla", "Pie"},
	}
	p2 := person{
		firstname: "Dev",
		lastname:  "Night",
		fav:       []string{"Real", "MTFK"},
	}
	fmt.Println(p1, p2)
}

func ex2() {
	p1 := person{
		firstname: "Phuc",
		lastname:  "Pham",
		fav:       []string{"Vanilla", "Pie"},
	}
	p2 := person{
		firstname: "Dev",
		lastname:  "Night",
		fav:       []string{"Real", "MTFK"},
	}
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for i, v := range m {
		fmt.Println("Person", i)
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for _, v2 := range v.fav {
			fmt.Println("Favorite of", v.lastname, "Is", v2)
		}
	}
}

func ex3() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 3,
			color: "blue",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "purple",
		},
		fourWheel: false,
	}
	fmt.Println(truck1.color)
	fmt.Println(sedan1.color)
}

func ex4() {
	m := map[string]int{
		"Bro":  2,
		"Mate": 3,
	}
	anon := struct {
		firstname string
		lastname  string
		anon      bool
		friends   map[string]int
		fav       []string
	}{
		"ok",
		"hmm",
		false,
		m,
		[]string{"Vani", "Pie", "Choco"},
	}
	fmt.Println(anon.firstname)
	fmt.Println(anon.lastname)
	for _, v := range anon.friends {
		fmt.Println(v)
	}
	for _, v := range anon.fav {
		fmt.Println(v)
	}
}
