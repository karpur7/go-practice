package main
import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perimeter() float64
}

type rect struct {
    length, breadth float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
    return r.length * r.breadth
}

func (r rect) perimeter() float64 {
    return 2 * (r.length + r.breadth)
}

func measure(g geometry) (float64, float64) {
    return g.area(), g.perimeter()
}

func main() {
    r := rect{length: 10, breadth: 20}
    c := circle{radius: 20}

    a, p := measure(r)
    fmt.Printf("Area and perimeter of rectangle: %f , %f\n", a,p )
    a, p = measure(c)
    fmt.Printf("Area and perimeter of circle: %f , %f\n", a,p )
    b := new(rect)
    a, p = measure(b)
    fmt.Printf("Area and perimeter of rectangle: %f , %f\n", a,p )
    fmt.Println(*b)
}
