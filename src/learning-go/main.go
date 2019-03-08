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

	collection()

	loops()

	errorHandlingAndDeferPanicAndRecover()
}

func errorHandlingAndDeferPanicAndRecover(){
	defer func(){
		if err := recover(); err != nil{
fmt.Println("handling error: ", err)
			panic("MyError message")
		}else{
fmt.Println("Nothing happened")
		}
	}()
	xs := []int{1}
	fmt.Println("xs:", xs[0])
}

func loops(){
	for i := 0; i < 5; i++{
	  fmt.Println("Counter:", i)
	}

	xs := []int{1,2,3,4}
	for i, e := range xs{
	  fmt.Printf("i=%d, and e = %d\n", i, e)
	}

	kv := map[string]string{"foo":"bar", "baz":"quz"}
	for k, v := range(kv){
	  fmt.Printf("k=%s, v=%s\n", k, v)
	}
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

func collection(){
	// array
	xs := [2]int{1,2}
	// slice
	ys := []int{1,2,3}

	// map
	kv := map[string]string{"foo":"bar", "quz":"quz", "abc":"def"}
	fmt.Printf("xs=%d, ys=%d, zs=%s", xs, ys, kv)
}
