package main
import "fmt"

func callbyval(val int) {
    val = 20
}

func callbyref(ptr *int) {
    *ptr = 20
}

func main() {
    i := 0
    fmt.Println("initial val of i: ", i)
    callbyval(i)
    fmt.Println("after callbyval val of i: ", i)
    callbyref(&i)
    fmt.Println("after callbyref val of i: ", i)
}
