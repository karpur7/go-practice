package main

import (
    "fmt"
    "encoding/xml"
    "os"
)

type Address struct{
    city, state string
}

type Person struct{
    XMLName xml.Name `xml:"person"`
    Id int `xml:"id,attr"`
    FirstName string `xml:"name>first"`
    LastName string `xml:"name>last"`
    Age int `xml:"age"`
    Height float32 `xml:"height,omitempty"`
    Married bool
    Address
    Comments string `xml:",comment"`
}

func main() {
    p := &Person{ Id : 13, FirstName : "Karpur", LastName : "Tadiparti",}
    p.Address = Address{ city : "Bangalore", state : "Karnataka"}
    p.Comments = " Need more details. "
    enc := xml.NewEncoder(os.Stdout)
    enc.Indent(" ", "    ")
    if err := enc.Encode(p);err != nil {
        fmt.Println("caught error ", err)
    }
    output, err := xml.MarshalIndent(p, "  ", "    ")
    if err == nil {
        os.Stdout.Write(output)
   }
}
