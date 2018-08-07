package main
import "fmt"

type rect struct {
    length, width int
}

func (r *rect) area() int {
    return r.length * r.width
}

func (r rect) perimeter() int {
    return 2 * ( r.length + r.width)
}

func main() {
    r1 := rect{length:20, width:30}
    fmt.Println("Area of rectangle: ", r1.area())
    fmt.Println("Perimeter of rectangle: ", r1.perimeter())
    r2 := &r1
    fmt.Println("Area of rectangle: ", r2.area())
    fmt.Println("Perimeter of rectangle: ", r2.perimeter())
}
