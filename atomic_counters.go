package main
import (
    "fmt"
    "time"
    "sync/atomic"
)

func main() {
    var counter uint64

    incrementer := func(){
        for {
            atomic.AddUint64(&counter, 1)
            time.Sleep(time.Millisecond)
        }
    }

    for i := 0; i < 50; i++ {
        go incrementer()
    }

    time.Sleep(time.Second)
    count_final := atomic.LoadUint64(&counter)
    fmt.Println("final count: ", count_final)
}
