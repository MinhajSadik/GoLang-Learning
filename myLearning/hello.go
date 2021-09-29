package myLearning

import (
	"fmt"
)

func Hello() string{
	var ami = "a new person Minhaj"
	// x := "someone"
	fmt.Println(ami)
	return newFunction()
}

func newFunction() string{
	return myFunc()
}

func myFunc() string{
	return "a new person Anika"
}

