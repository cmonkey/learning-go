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
	"io/ioutil"
	"bufio"
	"encoding/csv"
	"compress/gzip"
	"io"
	"bytes"
	"encoding/base64"
	"crypto/md5"
	"crypto/sha256"
	"crypto/aes"
	cryptoRand "crypto/rand"
	"crypto/cipher"
	"encoding/hex"
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

	writeBytesToFile()

	//readingFiles()

	writeingCSV()

	readingCSV()

	compression()

	readingZIP()

	base64EncodingAndDecoding()

	myMd5()

	sha256Digest()

	encryptByAES()
}

func encryptByAES(){

	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("Hello World!")
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 12)
	io.ReadFull(cryptoRand.Reader, nonce)
	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Nonce: %x\n", nonce)
	fmt.Printf("Cipher: %x\n", ciphertext)
}

func sha256Digest(){

	h := sha256.New()
	h.Write([]byte("Hello World!\n"))
	fmt.Printf("%x", h.Sum(nil))
}

func myMd5(){

	h := md5.New()
	h.Write([]byte("Hello Wrold!\n"))
	fmt.Printf("%xi", h.Sum(nil))
}

func base64EncodingAndDecoding(){

	msg := "Hello World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	enc := "SGVsbG8gV29ybGQh"
	decoded, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(string(decoded))
}

func compression(){

	f, e1 := os.Create("/tmp/test.gz")
	defer f.Close()
	check(e1)
	gz := gzip.NewWriter(f)
	defer func(){

		gz.Flush()
		gz.Close()
	}()

	_, e2 := gz.Write([]byte("hello world\n"))
	check(e2)
}

func readingZIP(){

	f, e1 := os.Open("/tmp/test.gz")
	defer f.Close()
	check(e1)
	gz, e2 := gzip.NewReader(f)
	check(e2)
	var buf bytes.Buffer
	_, e3 := io.Copy(&buf, gz)
	check(e3)
	fmt.Print(string(buf.Bytes()))

}

func writeingCSV(){

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},

	}

	// write a record to console
	f, _ := os.Create("/tmp/names.csv")
	defer f.Close()
	cw := csv.NewWriter(bufio.NewWriter(f))

	for _, record := range records {

		if err := cw.Write(record); err != nil{
			log.Fatal("error writing record to csv:", err)
		}
	}
	cw.Flush()
}

func readingCSV(){

	f, _ := os.Open("/tmp/names.csv")
	defer f.Close()
	cr := csv.NewReader(bufio.NewReader(f))
	names, _ := cr.ReadAll()
	for _, row := range names{
		for i, field := range row {

			fmt.Println(i, field)
		}
	}
}

func check(e error){

	if e != nil{
		panic(e)
	}
}

func readingFiles(){

	data,err := ioutil.ReadFile("/tmp/test2.dat")
	check(err)
	fmt.Print(string(data))

	// read 5 bytes
	f, err := os.Open("/tmp/test2.dat")
	defer f.Close()
	check(err)
	b1 := make([]byte,5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// start reading from the 6th byte
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 6)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes 0 %d: %s\n", n2, o2, string(b2))

	// set the file pointer to the beginning
	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	// read 5 bytes without advancing the reader
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
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

func writeBytesToFile(){
	// write bytes to a file
	data := []byte("Hello There!\n")
	err := ioutil.WriteFile("/tmp/test1.txt", data, 0644)
	check(err)

	// write bytes to a file
	data2 := []byte{115, 111, 109, 101, 10}
	file2, err := os.Create("/tmp/test2.dat")
	defer file2.Close()
	check(err)
	n, err := file2.Write(data2)
	fmt.Printf("wrote %d bytes\n", n)
	check(err)
	file2.Sync()

	// write bytes to a file more efficiently
	file3, err := os.Create("/tmp/test3.dat")
	defer file3.Close()
	writer := bufio.NewWriter(file3)
	n2, err := writer.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n2)
	writer.Flush()

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
