package main
import (
    "fmt"
    "encoding/xml"
)

type Email struct {
    Where string `xml:"Where,attr"`
    Addr string
}

type Address struct {
    City,State string
}

type Person struct {
    XMLName xml.Name `xml:"Person"`
    Name string `xml:"FullName"`
    Email []Email
    Groups []string `xml:"Group>Value"`
    Phone string
    Address
}

func main() {
    person := &Person{Name: "none",Phone: "none"}
    data := `
            <Person>
                <FullName>Karpur Tadiparti</FullName>
                <Company>Nokia Networks</Company>
                <Email Where="work">
                    <Addr>karpur.tadiparti@nokia.com</Addr>
                </Email>
                <Email Where="personal">
                    <Addr>karpur7@gmail.com</Addr>
                </Email>
                <Group>
                    <Value>MN ATF</Value>
                    <Value>RUBY</Value>
                </Group>
                <City>Bangalore</City>
                <State>Karnataka</State>
            </Person>
    `
    err := xml.Unmarshal([]byte(data),&person)
    if err != nil {
        fmt.Printf("error: %v\n",err)
        return
    }
    fmt.Printf("XMLName:%#v\n", person.XMLName)
    fmt.Printf("Name:%q\n", person.Name)
    fmt.Printf("Email:%v\n", person.Email)
    fmt.Printf("Phone:%q\n", person.Phone)
    fmt.Printf("Groups:%v\n", person.Groups)
    fmt.Printf("Address:%v\n", person.Address)
}
