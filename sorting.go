package main
import (
    "fmt"
    "sort"
)

func main() {
    strs := []string{"c","a","d","b"}

    sort.Strings(strs)

    fmt.Println("sorted list: ", strs)

    ints := []int{8,9,3,45,2}
    sort.Ints(ints)

    fmt.Println("sorted list: ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("is ints sorted: ",s)
}
