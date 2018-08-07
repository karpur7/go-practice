package main
import "fmt"

func main() {
    s1 := make([]string, 3, 5)
    s1[0] = "k"
    s1[1] = "a"
    s1[2] = "r"
    s1 = append(s1, "p")
    s1 = append(s1, "u")
    s1 = append(s1, "r")
    fmt.Println(s1,len(s1),cap(s1))

    s2 := make([]string, len(s1))
    copy(s2, s1)
    fmt.Println(s2, len(s2), cap(s2))

    fmt.Println(s2[2:])
    fmt.Println(s2[:len(s2)])

    s3 := []string{"kav"}
    fmt.Println(s3, len(s3), cap(s3))
    s3 = append(s3, "i")
    fmt.Println(s3, len(s3), cap(s3))


    twoD := make([][]string, 3)
    for i := 0 ; i < len(twoD) ; i++ {
        twoD[i] = make([]string, i+1)
        for j := 0 ; j < len(twoD[i]) ; j++ {
            twoD[i][j] = "karpur"
        }
    }
    fmt.Println(twoD)
}
