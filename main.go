package main

import (
	"fmt"
	"strconv"

	"example.com/GoLang-Learning/myLearning"
)


func main() {
    fmt.Print("hello world -- ")
    s := "Hello World!"
    fmt.Println(s)
    fmt.Println(myLearning.Hello())

    // var s string = "Hello, World"

    for index, character := range(s){
        fmt.Printf("The character %c is in position %d \n", character, index)
    }

}


// func Sqrt(x float64) float64 {
//     z := 0.0
//     for i := 0; i < 1000; i++ {
//         z -= (z*z - x) / (2 * x)
//     }
//     return z
// }


func Learing() string {
    var a int32 = 100
    var b int64 = int64(a)
    fmt.Println(b)
    var str int = 32
    var d string = strconv.Itoa(str)
    fmt.Println(str, d)

    return "true"
}