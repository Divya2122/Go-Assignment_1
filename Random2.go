
package main 

import ( 
	"fmt"
	"math/rand"
) 

 
func intnrandom(value_1, value_2 int) int { 
	return value_1 + value_2 + rand.Intn(4) 

} 

 
func main() { 

	// Getting result from Intnrandom() function 
	res1 := intnrandom(11, 19) 
	res2 := intnrandom(27, 49) 
	res3 := intnrandom(240, 78) 

	// Displaying results 
	fmt.Println("Result 1: ", res1) 
	fmt.Println("Result 2: ", res2) 
	fmt.Println("Result 3: ", res3) 
} 
