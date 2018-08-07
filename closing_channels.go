package main
import (
    "fmt"
    "time"
)

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("Received job: ", j)
            } else {
                fmt.Println("Received all jobs")
                done <- true
                return
            }
        }
    }()

    for i := 0; i < 3; i++ {
        jobs <- i
    }
    time.Sleep(1*time.Second)
    close(jobs)
    fmt.Println("Sent all jobs")
    <-done
}
