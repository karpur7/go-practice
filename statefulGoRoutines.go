package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync/atomic"
)

type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    val int
    resp chan bool
}

func main(){
    var readOps uint64
    var writeOps uint64

    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func(){
        state := make(map[int]int)
        for {
            select {
                case read := <-reads:
                    read.resp <- state[read.key]
                case write := <-writes:
                    state[write.key] = state[write.val]
                    write.resp <- true
            }
            time.Sleep(time.Millisecond)
        }
    }()

    for i := 0 ; i < 100; i++ {
        go func(){
            for {
                read := &readOp {
                    key : rand.Intn(5),
                    resp : make(chan int) }
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for i := 0 ; i < 10 ; i++ {
        go func() {
            for {
                write := &writeOp {
                    key : rand.Intn(5),
                    val : rand.Intn(100),
                    resp : make(chan bool) }
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }


    time.Sleep(time.Second)
    readOps_final := atomic.LoadUint64(&readOps)
    writeOps_final := atomic.LoadUint64(&writeOps)
    fmt.Println("readOps: ", readOps_final)
    fmt.Println("writeOps: ", writeOps_final)
}
