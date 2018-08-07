package main
import "fmt"

func main() {
    arr := []int{1,2,3,4}
    for _,el :=range arr {
        fmt.Println(el)
    }

    s := make([]string, 5)
    s[0] = "hi"
    s[1] = "how"
    s[2] = "are"
    s[3] = "you"
    s[4] = "?"

    for _,el :=range s {
        fmt.Println(el)
    }

    m := map[int]string{1:"hello", 2:"karpur"}
    fmt.Println(m)

    for k,v := range m {
        fmt.Println(k,v)
    }
}
