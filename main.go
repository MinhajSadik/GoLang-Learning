package main

import (
	"fmt"
	"math"
)

// func myEarning(name string) string{
//     return name
// }

// func getName(age int) int{
//     return age
// }

func sayGreeting(n string) {
    fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
    fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)){
    for _, v := range n {
        f(v)
    }
}

func circleArea(r float64) float64{
    return math.Pi * r * r
}

func main() {
    // fmt.Print("hello world -- ")
    // s := "Hello World!"
    // fmt.Println(s)
    // x := "Samira"
    // fmt.Println(myLearning.Hello(), myEarning(x), getName(22))


    // sayGreeting("Minhaj")
    // sayGreeting("Minhaj")
    // sayBye("Ahmed")

    cycleNames([]string{"Minhaj", "Ahmed", "Sadik"}, sayGreeting)
    cycleNames([]string{"Minhaj", "Ahmed", "Sadik"}, sayBye)

    a1 := circleArea(10.8)
    a2 := circleArea(10)

    fmt.Println(a1, a2)
    fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)



    // var s string = "Hello, World"
    // for index, value := range(s){
    //     fmt.Printf("The character %c is in position %d \n", value, index)
    // }

    // var a int32 = 100
    // var b int64 = int64(a)
    // fmt.Println(b)
    // var str int = 32
    // var d string = strconv.Itoa(str)
    // fmt.Printf("%v, %T \n", d , d)

}

