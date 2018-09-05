package main
import (
    "fmt"
    "log/syslog"
)

func main() {
    loger,err := syslog.New(syslog.LOG_INFO, "myprg")
    if err != nil {
        fmt.Println("error while connecting to log server : ",err)
        return
    }
    loger.Info("testing with emerg")
}
