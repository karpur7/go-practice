package main

import (
    "fmt"
    "container/heap"
)

type Item struct {
    value string
    priority int
    index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i],pq[j] = pq[j],pq[i]
    pq[i].index = j
    pq[j].index = i
}

func (pq *PriorityQueue) Push(el interface{}) {
    item := el.(*Item)
    item.index = len(*pq)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(*pq)
    item := old[n-1]
    item.index = -1
    *pq = old[:n-1]
    return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
    item.value = value
    item.priority = priority
    heap.Fix(pq, item.index)
}

func main() {
    items := map[string]int{
             "banana" : 3, "apple" : 4, "pear" : 2,
    }

    pq := make(PriorityQueue, len(items))

    i := 0
    for val, pr := range items {
        pq[i] = &Item {
                value : val,
                priority : pr,
                index : i}
        i++
    }
    heap.Init(&pq)
    item := &Item {
            value : "orange",
            priority : 6 }
    heap.Push(&pq, item)
    pq.update(item, item.value, 7)
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        fmt.Printf("%.2d:%s",item.priority, item.value)
    }
}
