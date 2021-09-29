package main

import (
	"fmt"
	"strconv"

	"example.com/GoLang-Learning/myLearning"
)


func myEarning(name string) string{
    return name
}

func getName(age int) int{
    return age
}

func main() {
    fmt.Print("hello world -- ")
    s := "Hello World!"
    fmt.Println(s)
    x := "Samira"
    fmt.Println(myLearning.Hello(), myEarning(x), getName(32))

    // var s string = "Hello, World"

    // for index, value := range(s){
    //     fmt.Printf("The character %c is in position %d \n", value, index)
    // }

    var a int32 = 100
    var b int64 = int64(a)
    fmt.Println(b)
    var str int = 32
    var d string = strconv.Itoa(str)
    fmt.Printf("%v, %T \n", d , d)

}

