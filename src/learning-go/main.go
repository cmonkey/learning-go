package main

import (
	"fmt"
	"log"
	"math"
	"formulas"
)

func main(){
	log.Println("Hello World!") // write to stdlog
	fmt.Println("hello World!") // write to stdout
	println("Hello World!") // write to stderr

	const a bool = true
	const b int = 42
	const c float32 = 3.1415
	const d string = "cmonkey"
	const e complex64 = 12 + 5i
	fmt.Println(a, b, c, d, e)

	shorter()

	fmt.Println(hello(world))

	mymath()

	formulasFunc()

	controlFlow()

	myMake()
}

func controlFlow(){
	x := 2
	if x == 1{
	  fmt.Println("One")
	}else if x == 2{
	  fmt.Println("Two")
	}else{
	  fmt.Println("something else")
	}

	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("something else")

	}
}

func mymath(){
	pi := math.Pi
	circum := pi * 6
	surface := math.Pow(6,2) * pi / 4
	fmt.Printf("circum %f surface %f", circum, surface)
}

func formulasFunc(){
	d := 6.25242
	fmt.Printf("circum %f surface %f", formulas.Circum(d), formulas.Surface(d))
}

func world() string {
	return "World"
}

func hello(f func() string) string{
	return "Hello " + f() + "!"

}

func shorter(){
	a := true
	b := 42
	c := 3.1415
	d := "cmonkey"
	e := 12 + 5i
	fmt.Println(a, b, c , d, e)
}

func myMake(){
	xs := make([]int, 3)
	xs[0] = 1
	xs[1] = 2

	kv := make(map[string]string)
	kv["foo"] = "bar"
	kv["baz"] = "quz"

	fmt.Printf("xs is %d, and ys is %s", xs, kv)
}
