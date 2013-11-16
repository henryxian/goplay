/*    newton's method    */
package main

import (
	"fmt"
    "math"
)

func Sqrt(x float64) float64 {
    var val = x
    var last float64
    for math.Abs(val - last) > 0.000000001{
        last = val
        val = (val + x/val) / 2
    }
    return val
}

func main() {
	fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}