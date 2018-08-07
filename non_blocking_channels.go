package main
import (
    "fmt"
)

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
        case msg := <- messages :
            fmt.Println(msg)
        default:
            fmt.Println("no message")
    }

    select {
        case messages <- "Hi":
            fmt.Println(<-messages)
        default:
            fmt.Println("no messages")
    }

    select {
        case msg := <- messages:
            fmt.Println(msg)
        case sgl := <- signals:
            fmt.Println(sgl)
        default:
            fmt.Println("no activity")
    }
}
