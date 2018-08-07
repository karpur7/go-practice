package main
import "fmt"

func main() {
    chn := make(chan string)
    go func(){ chn <- "hello world!!!"}()

    msg := <- chn

    fmt.Println(msg)
}
