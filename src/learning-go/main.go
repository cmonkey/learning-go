package main

import (
	"fmt"
	"log"
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

}

func shorter(){
	a := true
	b := 42
	c := 3.1415
	d := "cmonkey"
	e := 12 + 5i
	fmt.Println(a, b, c , d, e)
}
