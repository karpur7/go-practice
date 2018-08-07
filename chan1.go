package main
import (
    "fmt"
    "time"
)

func sync_func(ch chan string) {
    ch <- "hi"
    time.Sleep(time.Duration(2))
    ch <- "hw r u kavitha"
}

func main() {
    chn := make(chan string, 2)

    go sync_func(chn)

    fmt.Println(<-chn)
    fmt.Println("channel is sleeping")
    fmt.Println(<-chn)
}
