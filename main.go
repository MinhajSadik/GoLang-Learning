package main

// func myEarning(name string) string{
//     return name
// }

// func getName(age int) int{
//     return age
// }

// func sayGreeting(n string) {
//     fmt.Printf("Good Morning %v \n", n)
// }

// func sayBye(n string) {
//     fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)){
//     for _, v := range n {
//         f(v)
//     }
// }

// func circleArea(r float64) float64{
//     return math.Pi * r * r
// }

// func getInitals(n string) (string, string){
//     s:= strings.ToUpper(n)
//     names:= strings.Split(s, " ")

//     var initials []string
//     for _, v:= range names{
//         initials = append(initials, v[:1])
//     }
//     if len(initials) > 1 {
//         return initials[0], initials[1]
//     }
//     return initials[0], "_"
// }

// var scores = 99.9

// Group A type => strings, ints, bools, floats, arrays, structs
// func updateName(x string){
//     x = "Ahmed"
//     fmt.Println("main", x)
// }

// Group B type change and add => slices, map, functions
// func updateMenu(y map[string] float64){
//     y["coffee"] = 2.99
//     y["lolipop"] = 2.09
// }

// func updateName(x *string){
//     *x = "Ahmed"
//     // fmt.Println("'main'", x)
// }

// //helper function
// func getInput(prompt string, r *bufio.Reader) (string, error) {
//     fmt.Print(prompt)
//     input, err := r.ReadString('\n')
//     return strings.TrimSpace(input), err

// }

// // user input
// func createBill() bill{
//     reader := bufio.NewReader(os.Stdin)

//     // fmt.Print("Create a new bill name: ")
//     // name, _ := reader.ReadString('\n')
//     // name = strings.TrimSpace(name)
//     name, _ := getInput("Created a new bill name: ", reader)

//     b:= newBill(name)
//     fmt.Println("Bill created - ", b.name)

//     return b
// }
// // prompt
// func promptOptions(b bill){
//     reader := bufio.NewReader(os.Stdin)

//     opt, _:= getInput("Choose Options (a - add item, s - save item, t - add tip):", reader)
//     fmt.Println(opt)
// }


func main() {

    // // user input 
    // mybill := createBill()
    // promptOptions(mybill)
    // fmt.Println(mybill)

    // // Receiver Functions with Pointers
    // mybill:= newBill("minhajSadik")

    // mybill.addItem("onion", 32)
    // mybill.addItem("orange", 29)
    // mybill.addItem("mango", 20)
    // mybill.addItem("apple", 23)
    // mybill.addItem("juice", 39)

    // mybill.updateTip(10)
    // fmt.Println(mybill.format())

    // // Recevier Function in Format Bills

    // mybill:= newBill("minhajSadik")
    // fmt.Println(mybill)
    // mybill.format()
    // fmt.Println(mybill.format())


    // //Struct in GoLang
    // mybill:= newBill("minhajSadik")
    // fmt.Println(mybill)


    // //Pointer in GoLang
    // name:= "Sadik"
    // updateName(name)
    // fmt.Println("memory address of 'name' is:", &name)
    // m:= &name
    // fmt.Println("memory address of 'm':", m)
    // fmt.Println("value at memory address of 'm' ", *m)

    // fmt.Println(name)
    // updateName(m)
    // fmt.Println(name)



    // // Group A type can't change => strings, ints, bools, floats, arrays, structs
    // name:= "Sadik"
    // updateName(name)
    // fmt.Println(name)

    // // Group B type change and add => slices, map, functions
    // menu:= map[string] float64{
    //     "ice creame": 4.93,
    //     "cocoa": 30.88,
    // }
    // updateMenu(menu)
    // fmt.Println(menu)


    // fmt.Print("hello world -- ")
    // s := "Hello World!"
    // fmt.Println(s)
    // x := "Samira"
    // fmt.Println(myLearning.Hello(), myEarning(x), getName(22))


    // sayGreeting("Minhaj")
    // sayGreeting("Minhaj")
    // sayBye("Ahmed")

    // sayHello("Minhaj")
    // showScore()
    // goMap()

    // for _, v:= range points{
    //     fmt.Println(v)
    // }

    // menu:= map[string]float64{
	// 	"mengo": 39,
	// 	"orange": 90,
	// 	"apple": 32,
	// }
	// fmt.Println(menu)
	// fmt.Println(menu["orange"])

    // phoneBook:= map[int]string{
	// 	182: "minhaj",
	// 	990: "sadik",
	// 	902: "Ahmed",
	// }
	// fmt.Println(phoneBook)
	// fmt.Println(phoneBook[182])

    // phoneBook[990] = "Anika"
    // fmt.Println(phoneBook)
    
    // phoneBook[902] = "Ahmed"
    // fmt.Println(phoneBook)

	// // looping map
	// for k, v:= range menu{
	// 	fmt.Println(k, "-", v)
	// }

    // fn1, sn1:= getInitals("Minhaj Sadik")
    // fmt.Println(fn1, sn1)

    // fn2, sn2:= getInitals("Amina Anika")
    // fmt.Println(fn2, sn2)

    // fn3, sn3:= getInitals("ANIKA")
    // fmt.Println(fn3, sn3)


    // cycleNames([]string{"Minhaj", "Ahmed", "Sadik"}, sayGreeting)
    // cycleNames([]string{"Minhaj", "Ahmed", "Sadik"}, sayBye)

    // a1 := circleArea(10.8)
    // a2 := circleArea(10)

    // fmt.Println(a1, a2)
    // fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)



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

