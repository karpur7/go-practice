package main
import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5)
    limiter := time.Tick(200*time.Millisecond)

    for i := 0; i < 5; i++ {
        requests <- i
    }
    close(requests)

    for req := range requests {
        <- limiter
        fmt.Println("received req ", req, " at ", time.Now())
    }

    burstlimiter := make(chan time.Time, 3)

    for i := 0 ; i < 3 ; i++ {
        burstlimiter <- time.Now()
    }

    go func(){
        t := time.NewTicker(200*time.Millisecond)
        for tic := range t.C {
            burstlimiter <- tic
        }
    }()

    newchannel := make(chan int, 5)
    for j := 0 ; j < 5; j++ {
        newchannel <- j
    }
    close(newchannel)

    for req := range newchannel {
        <- burstlimiter
        fmt.Println("received req ", req, " at ", time.Now())
    }
}
