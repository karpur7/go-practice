package main
import "fmt"

func sum(nums ...int) {
    res := 0
    for _,num := range nums {
        res += num
    }
    fmt.Printf("%d\n", res)
}

func main() {
    sum(1, 2, 3)
    arr := []int{3, 4, 5, 8}
    sum(arr...)
}
