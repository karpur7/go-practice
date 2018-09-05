package main
import (
    "fmt"
    "sort"
)

type bylength []string

func(s bylength) Less(i,j int) bool {
    return len(s[i]) < len(s[j])
}

func(s bylength) Swap(i,j int){
    s[i],s[j] = s[j],s[i]
}

func(s bylength) Len() int {
    return len(s)
}

func main(){
    str_arr := []string{"banana","kiwi","apple"}
    sort.Sort(bylength(str_arr))
    fmt.Println("sorted array: ", str_arr)
}
