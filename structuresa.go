package main

import (  
    "fmt"
)

type User struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    
    user1 := User{
        firstName: "Divya",
        age:       22,
        salary:    20000,
        lastName:  "Avula",
    }

    
    user2 := User{"mamatha", "seethaka", 22, 25000}
    user3 := User{"chandu", "seethe", 22, 45000}


    fmt.Println("User 1", user1)
    fmt.Println("User 2", user2)
    fmt.Println("User 3", user3)
}