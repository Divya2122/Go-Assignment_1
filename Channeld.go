 package main

 import "fmt"
 func main(){

 	myChannel := make(chan string)

 	go func(){
 		myChannel <- "Divya"
 		myChannel <- "Kloudone"
 	}()
 	text := <-myChannel

 	fmt.Println(text)
 	fmt.Println(<-myChannel)
}