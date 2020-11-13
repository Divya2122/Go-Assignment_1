package main

import (
	"fmt"
	"net/http"
)

func myName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi, this is Divya! Welcome to my Page.")
}


func helloEveryone(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello!")
}


func main(){

	http.HandleFunc("/",  myName)
	http.HandleFunc("/mahesh", helloEveryone)
	http.ListenAndServe(":9000",nil)


}