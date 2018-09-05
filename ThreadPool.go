package main
import (
    "fmt"
    "time"
)

func worker(tid int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker thread " , tid, " started job ", j)
        time.Sleep(time.Second)
        fmt.Println("worker thread " , tid, " finished job ", j)
        results <-j*2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    for j := 1; j <= 3; j++ {
        go worker(j, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs);

    for j := 1; j <= 5; j++ {
        <- results
    }
}
