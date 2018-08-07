package main
import "fmt"

func func1(s string) {
    for i:=0; i < 5; i++ {
        fmt.Println(s)
    }
}

func main() {
    func1("hi")

    go func1("hello")

    go func(s string) {
        fmt.Println(s)
    }("world")

    fmt.Scanln()
    fmt.Println("Done")
}
