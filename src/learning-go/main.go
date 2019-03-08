package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"formulas"
	"time"
	"sync"
	"github.com/google/uuid"
	"os"
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

	pointers()

	goroutines()

	//channels()

	//logger()

	randomNumberGenerator()

	myUUid()

	writeStringToFile()
}

func check(e error){

	if e != nil{
		panic(e)
	}
}

func writeStringToFile(){

	f, err := os.Create("/tmp/test.txt")
	check(err)
	defer f.Close()
	n3, err := f.WriteString("hello World\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}

func myUUid(){

	for i:=0; i<5;i++{

		id, _ := uuid.NewRandom()
		fmt.Println(id.String())
	}
}

func randomNumberGenerator(){


	for i := 0; i < 5; i++{

		x := rand.Intn(100)
		y := rand.Float64() * 5
		fmt.Println(x, y)
	}
}

func logger(){

	log.Println("Normal")
	log.Fatal("Fatal")
	// log and then panic
	log.Panic("Panic")
}

func channels(){

	done := make(chan bool)
	ping := make(chan string, 10)
	pong := make(chan string, 10)
	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- "ping"
	<-done
	fmt.Println("Done")
}

func pinger(ping chan string, pong chan string){
	for {
		fmt.Println(<-ping)
		time.Sleep(time.Millisecond * 300)
		pong <- "pong"
	}
}

func ponger(ping chan string, pong chan string){

	for{
		fmt.Println(<-pong)
		time.Sleep(time.Millisecond * 300)
		ping <- "ping"
	}
}

var wg sync.WaitGroup
func say(msg string){
	for i := 0; i < 3; i++{
		fmt.Printf("%s world!\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
func goroutines(){
	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("goodbye")
	wg.Wait()
}

func addOne(x *int){
      *x = *x + 2
}

func pointers(){
	x := 1
	addOne(&x)
	fmt.Println("x is ", x)
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
