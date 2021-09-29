package main

type bill struct {
	name string
	itmes map[string] float64
	tip float64

}

// make new bills

func newBill(name string) bill {

	b := bill {
		name: name,
		itmes: map[string]float64{},
		tip: 0,
	}

	return b
}
