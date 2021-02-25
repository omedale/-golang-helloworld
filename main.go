package main

// import (
// 	"io"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello, world!\n")
// }

// // Route declaration
// func router() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", handler)
// 	return r
// }

// // Initiate web server
// func main() {
// 	router := router()
// 	srv := &http.Server{
// 		Handler:      router,
// 		Addr:         "127.0.0.1:9100",
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}

// 	log.Fatal(srv.ListenAndServe())
// }

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
)

// var c, python, java bool
var c, python, java = true, false, "no!" // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func initialValuetTest() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func typeConvertion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forLoop2() {
	sum := 1
	for sum < 2 {
		sum += sum
	}
	fmt.Println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func switchCase() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//A defer statement defers the execution of a function until the surrounding function returns.
func deferFunction() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X int
	Y int
}

var (
	v1 = Vertex2{1, 2} // has type Vertex
	v2 = Vertex2{X: 1} // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex, : The special prefix & returns a pointer to the struct value.
)

func vertex() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// accessing value by pointers
	t := Vertex{1, 2}
	p := &t
	p.X = 1e9
	fmt.Println(t)

	fmt.Println(v1, p2, v2, v3)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
func slices() {
	//A slice does not store any data, it just describes a section of an underlying array.
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)
}

func slicesDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = append(s, 4, 5)
}

func emptySliceIsNil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
//The make function allocates a zeroed array and returns a slice that refers to that array:
func makeSlice() {
	a := make([]int, 5)
	printMakeSlice("a", a)

	b := make([]int, 0, 5)
	printMakeSlice("b", b)

	c := b[:2]
	printMakeSlice("c", c)

	d := c[2:5]
	printMakeSlice("d", d)
}

func printMakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func displayTictacboard() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

var power = []int{1, 2, 4, 8, 16, 32, 64, 128}

func useRangeWithArray() {
	for i, v := range power {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//You can skip the index or value by assigning to _.
	powerr2 := make([]int, 10)
	for i := range powerr2 {
		powerr2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range powerr2 {
		fmt.Printf("%d\n", value)
	}
}

type MapVertex struct {
	Lat, Long float64
}

var m map[string]MapVertex

func maps() {
	m = make(map[string]MapVertex)
	m["Bell Labs"] = MapVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	m2 := map[string]MapVertex{
		"Bell Labs": MapVertex{
			40.68433, -74.39967,
		},
		"Google": MapVertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]MapVertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	mapValue := make(map[string]int)

	mapValue["Answer"] = 42
	fmt.Println("The value:", mapValue["Answer"])

	mapValue["Answer"] = 48
	fmt.Println("The value:", mapValue["Answer"])

	delete(mapValue, "Answer")
	fmt.Println("The value:", mapValue["Answer"])

	// Test that a key is present with a two-value assignment:
	// elem, ok = m[key]
	// If key is in m, ok is true. If not, ok is false.
	v, ok := mapValue["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	var i int
	k := 3  //Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	v := 42 // When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(add(42, 13))
	fmt.Println(split(17))
	fmt.Println(i, c, python, java)
	fmt.Println(k, v)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	initialValuetTest()
	typeConvertion()

	//Constants
	const World = "世界"
	fmt.Println("Hello", World)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	forLoop()
	forLoop2()
	whileLoop()
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switchCase()

	deferFunction()

	vertex()
	arrays()
	slices()
	slicesDefaults()
	makeSlice()
	displayTictacboard()
	useRangeWithArray()
	maps()
}
