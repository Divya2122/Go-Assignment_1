package main 
  
import ( 
    "fmt"
    "math/rand"
) 
  

func main() { 
  
    
    res_1 := rand.Intn(7) 
    res_2 := rand.Intn(8) 
    res_3 := rand.Intn(2) 
  
    
    fmt.Println("Random Number 1: ", res_1) 
    fmt.Println("Random Number 2: ", res_2) 
    fmt.Println("Random Number 3: ", res_3) 
} 