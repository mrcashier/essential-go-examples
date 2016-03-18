package main

import "fmt"

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

type circle struct {
	ratio int
}

func (c circle) diameter() float32  {
	return float32(c.ratio) / 2
}

func main() {
	p := point{20, 40}

	fmt.Println(p)
	fmt.Println(p.x)

	p2 := point{
		x: 10,
		y: 15,
	}

	fmt.Println(p2)

	r := rect{
		pos:    NewPoint(20, 30),
		width:  200,
		height: 400,
	}

	fmt.Println(r)
	fmt.Printf("Area from rectangle %v is %v \n", r, r.area())

	c := circle{55}

	fmt.Printf("Diameter of circle %v is %v", c, c.diameter() )
}
