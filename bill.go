package main

import "fmt"

type bill struct {
	name string
	itmes map[string] float64
	tip float64

}

// make new bills

func newBill(name string) bill {

	b := bill {
		name: name,
		itmes: map[string]float64{"pic": 4.9, "cake": 9.3},
		tip: 0,
	}

	return b
}

// format the bill Receiver Function
func (b bill) format() string {
	fs := "Bill breakedown: \n"
	var total float64 = 0

	// list items 
	for k, v := range b.itmes  {
		fs += fmt.Sprintf("%-25v ...$%v \n", k + ":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%25v ...$%0.2f", "total:", total)
	// fs += fmt.Sprintf("%v ...$%0.2f","total:", total)
	return fs
}