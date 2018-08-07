package main
import "fmt"

func plus(a int, b int) int {
    return a + b
}

func plus1(a, b, c int) int {
    return a + b + c
}

func main() {
    a, b := 1, 2
    c, d, e := 5, 6, 7
    res1 := plus(a, b)
    res2 := plus1(c, d, e)
    fmt.Printf("%d\n", res1)
    fmt.Printf("%d\n", res2)
}
