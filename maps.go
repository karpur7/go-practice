package main
import "fmt"

func main() {
    m1 := make(map[string]int, 5)
    m1["karpur"] = 200
    m1["kavitha"] = 200
    m1["kamakshi"] = 400
    fmt.Println(m1)
    delete(m1, "kamakshi")
    fmt.Println(m1)
    _, ok := m1["kamakshi"]
    fmt.Println(ok)
    delete(m1, "kamaskshi")
    fmt.Println(m1)
}
