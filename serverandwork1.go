package main

import(
	"fmt"
	"time"
)


func sayHello(s string){
	for i :=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go sayHello("Hello")
	go sayHello("Divya")
	go sayHello("hi")
	time.Sleep(time.Second)
fmt.Println("done")
	fmt.Println("Success")
}


