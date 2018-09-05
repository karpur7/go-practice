package main
import (
    "fmt"
    "os"
)

func main(){
    f := CreateFile("/tmp/temporaryFile")
    defer CloseFile(f)
    WriteFile(f)
}

func CreateFile(p string) *os.File {
    fmt.Println("Creating the file")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func WriteFile(f *os.File) {
    fmt.Println("writing some data")
    fmt.Fprintln(f,"hi karpur")
}

func CloseFile(f *os.File) {
    fmt.Println("Closing the file")
    f.Close()
}
