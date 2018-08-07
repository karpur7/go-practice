package main
import (
    "fmt"
    "time"
)

func sync_func(chn chan string){
    for i := 0; i < 5; i++ {
        fmt.Println("sleeping")
        time.Sleep(time.Second)
        fmt.Println("slept")
    }
    chn <- "Done"
}

func main() {
    chn := make(chan string, 1)
    go sync_func(chn)

    <-chn
}
