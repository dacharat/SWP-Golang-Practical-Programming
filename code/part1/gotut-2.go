// package main

// import (
// 	"fmt"
// 	"math"
// )

// func foo() {
// 	fmt.Println("the square root of 4 is", math.Sqrt(4))
// }

// func main() {
// 	foo()
// }

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("A number from 1-100: ", rand.Intn(100))
}
