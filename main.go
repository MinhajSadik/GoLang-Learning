package main

import "fmt"


func main() {
    fmt.Printf("hello world")
    var s string = "Hello World"
    fmt.Println(s)

    // var s string = "Hello, World"
    for index, character := range(s){
        fmt.Printf("The character %c is in position %d \n", character, index)
    }
}
