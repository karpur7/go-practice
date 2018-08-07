package main
import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)

    go func(){
        time.Sleep(1*time.Second)
        ch1 <- "Hi"
    }()

    go func(){
        time.Sleep(2*time.Second)
        ch2 <- "Hw r u"
    }()

    for i := 0; i < 2; i++ {
        select {
            case msg1 := <-ch1:
                fmt.Println(msg1)
            case msg2 := <-ch2:
                fmt.Println(msg2)
        }
    }
}
