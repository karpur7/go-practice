package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    conn,_ := net.Dial("tcp",":8081")
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Enter message to send")
        message,_ := reader.ReadString('\n')
        fmt.Fprint(conn, message+"\n")
        //conn.Write([]byte(message+"\n"))
        reply,_ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("reply:",reply)
    }
}
