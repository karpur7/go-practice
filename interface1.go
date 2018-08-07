package main
import ("fmt")

type Animal interface {
    speak()
}

type Dog struct {
    name string
}

func (d Dog) speak() {
    fmt.Println("Bow Bow!!!")
}

type Cat struct {
    name string
}

func (c *Cat) speak() {
    fmt.Println("Meow Meow!!!")
}

func PrintAll(lst ...interface{}) {
    for _,a := range lst {
        fmt.Println(a)
    }
}

func Speak (lst ...Animal) {
    for _,a := range lst {
        a.speak()
    }
}

func main() {
    c := Cat{name:"I'm cat"}
    d := Dog{name:"I'm dog"}
    g := new(Cat)
    g.name = "new cat"
    Speak(&c,d,g)
    PrintAll(c,d,g,"this is karpur")
}
