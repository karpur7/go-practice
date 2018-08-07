package main
import "fmt"
import "time"

func main() {

    i := 2
    switch i {
    case 1:
        fmt.Println("one")
        fmt.Println("one")
        fmt.Println("one")
    case 2:
        fmt.Println("two")
        fmt.Println("two")
        fmt.Println("two")
        fmt.Println("two")
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Printf("before Noon %d", t.Hour())
    default:
        fmt.Println("after noon")
    }

    whatAmI := func (i interface{}) {
    switch t := i.(type) {
        case bool:
            fmt.Println("boolean type")
        case int:
            fmt.Println("integer type")
        default:
            fmt.Printf("I don't know %T\n", t)
    }
    }

    whatAmI(true)
    whatAmI(2e10)
    whatAmI("karpur")
}
