package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"runtime"
)

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return //naked return - only use in short functions
}
func Sqrt(x float64) float64 {
	var z float64 = 10
	for i:=0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z
}
var (
	bool3 bool
	bigInt uint64 	= 1<<64-1
)

func main() {
	defer fmt.Println("Tchüss!")
    fmt.Printf("hello, 世界\n")

    fmt.Println("The time is", time.Now())
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println("My favorite number is", rand.Intn(10), ", yours?")

    fmt.Printf("Now you have %g problems. \n", math.Pi)
    fmt.Println("Or just", add(4, 1))

    a, b := swap("hello", "world")
    fmt.Println("Now some word reversal:",a,b)
    fmt.Println(split(17))

    bool1, bool2 := true, "YASSSSS"
    fmt.Println("Random bools ",bool1, bool2, bool3)

    const f = "%T(%v)\n"
    fmt.Printf(f, bigInt, bigInt)
    var counter int = 3
    for counter < 100 {
    	counter += counter
    }
    fmt.Println(counter)
    fmt.Println(Sqrt(2))

    fmt.Print("Go runs on " )
    switch os := runtime.GOOS; os {
    case "darwin":
    	fmt.Println("OS X.")
    case "linux":
    	fmt.Println("Linux");
    default:
    	fmt.Printf("%s", os)

    }
}
