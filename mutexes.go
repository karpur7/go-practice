package main
import (
    "fmt"
    "time"
    "sync"
    "sync/atomic"
    "math/rand"
)

func main() {
    var state = make(map[int]int)
    var mutex = &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for i := 0; i < 100; i++ {
        total := 0
        go func(){
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for i := 0; i < 10; i++ {
        go func() {
            for {
                key := rand.Intn(5)
                value := rand.Intn(100)
                mutex.Lock()
                state[key] = value
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    time.Sleep(time.Second)
    final_read_ops := atomic.LoadUint64(&readOps)
    final_write_ops := atomic.LoadUint64(&writeOps)
    fmt.Println("read ops: ", final_read_ops)
    fmt.Println("write ops: ", final_write_ops)
    mutex.Lock()
    fmt.Println("state\n", state)
    mutex.Unlock()
}
