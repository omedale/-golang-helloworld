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
}
