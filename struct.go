package main
import "fmt"

type person struct {
    name string
    age int
}

func main() {
    s1 := person{"name1", 20}
    fmt.Println(s1)

    s2 := person{name: "name2", age: 30 }
    fmt.Println(s2)

    s3 := person{age:21}
    fmt.Println(s3,s3.name)
    fmt.Println(s3.name)

    s4 := &s1
    fmt.Println(s4.name, s4.age)
}
