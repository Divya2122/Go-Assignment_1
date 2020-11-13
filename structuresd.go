package main 
  
import "fmt"
  

type Employee struct { 
    firstName, lastName string 
    age, salary int
} 
  
func main() { 
  
    
    emp1 := &Employee{"divya", "Avula", 21, 25000} 
    emp2 := &Employee{"chandu", "seethi", 22, 6000} 
    emp3 := &Employee{"babu", "Avula", 21, 27000} 
    emp4 := &Employee{"mamatha", "seethaka", 24, 3000} 
  
    
    fmt.Println("First Name:", (*emp1).firstName) 
    fmt.Println("Age:", (*emp1).age) 
    fmt.Println("First Name:", (*emp2).firstName) 
    fmt.Println("Age:", (*emp2).age) 
    fmt.Println("First Name:", (*emp3).firstName) 
    fmt.Println("Age:", (*emp3).age) 
    fmt.Println("First Name:", (*emp4).firstName) 
    fmt.Println("Age:", (*emp4).age) 
} 