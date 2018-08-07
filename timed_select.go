package main
import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string, 1)
    c2 := make(chan string, 1)

    go func(c chan string){
        time.Sleep(2*time.Second)
        c <- "testing channel 1"
    }(c1)

    select {
        case msg := <-c1 :
            fmt.Println(msg)
        case <-time.After(1*time.Second):
            fmt.Println("channel 1 timedout")
    }

    go func(c chan string){
        time.Sleep(2*time.Second)
        c <- "testing channel 2"
    }(c2)

    select {
        case msg := <-c2 :
            fmt.Println(msg)
        case <-time.After(3*time.Second):
            fmt.Println("channel 2 timedout")
    }
}
