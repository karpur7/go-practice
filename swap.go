package main
import "fmt"

func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    a, b := 5, 12
    fmt.Printf("before swapping %d, %d\n",a,b)
    a, b = swap(a,b)
    fmt.Printf("after swapping %d, %d\n",a,b)
}
