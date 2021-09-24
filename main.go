package main

import (
	"fmt"
)


func main() {
    fmt.Print("hello world -- ")
    s := "Hello World!"
    fmt.Println(s)

    // var s string = "Hello, World"

    for index, character := range(s){
        fmt.Printf("The character %c is in position %d \n", character, index)
    }
}


func Sqrt(x float64) float64 {
    z := 0.0
    for i := 0; i < 1000; i++ {
        z -= (z*z - x) / (2 * x)
    }
    return z
}