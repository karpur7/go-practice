package main
import (
    "fmt"
    "errors"
)

var threshold int = 300
var threshold1 int = 200
type myerror struct{}

func error1(val int) (int, error) {
    if val > threshold {
        return -1, errors.New("value is greather than threshold " )
    }
    if val > threshold1 {
        return -2, &myerror{}
    }
    return 0, nil
}


func (e *myerror) Error() string {
    return fmt.Sprintf("my code; my error %d", 100)
}

func main() {
    ret, e := error1(300)
    fmt.Println(ret,e)
    
}
