package main

import "fmt"

func sentValue(details chan string){

	details <- "Divya"
	details <- "kavya"
	details <- "Abhi"
}

func main (){

	values := make(chan string)
    
	go sentValue(values)

	value :=  <- values

	defer fmt.Println(value)
		  fmt.Println(<-values)
		  fmt.Println(<-values)
}
