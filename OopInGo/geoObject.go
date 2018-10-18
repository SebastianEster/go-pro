package OopInGo

import "fmt"

type Painter interface {
	Paint()
}

type Position struct {
	x int
	y int
}

type GeoObject struct {
	color    int
	position Position
}

type Circle struct {
	GeoObject
	radius int
}

type Rectangle struct {
	GeoObject
	height int
	width  int
}

type Triangle struct {
	GeoObject
	a, b, c Position
}

func (geoObject GeoObject) Paint() {
	fmt.Printf("GeoObject at position %d/%d with color %d\n",
		geoObject.position.x,
		geoObject.position.y,
		geoObject.color)
}

func (circle Circle) Paint() {
	fmt.Printf("Circle at position %d/%d with color %d and radius %d\n",
		circle.position.x,
		circle.position.y,
		circle.color,
		circle.radius)
}

func (rectangle Rectangle) Paint() {
	fmt.Printf("Rectangle at position %d/%d with color %d, height %d and width %d\n",
		rectangle.position.x,
		rectangle.position.y,
		rectangle.color,
		rectangle.height,
		rectangle.width)
}

func (triangle Triangle) Paint() {
	fmt.Printf("Triangle at position %d/%d with color %d and edges at %v, %v and %v\n",
		triangle.position.x,
		triangle.position.y,
		triangle.color,
		triangle.a,
		triangle.b,
		triangle.c)
}