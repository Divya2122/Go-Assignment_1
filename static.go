package main

import (
	//"fmt"
	"net/http"
	"html/template"
)

func myName(w http.ResponseWriter, r *http.Request){
	t, _ := 	template.ParseFiles("sample.html")
	data := "Divyareddy"
	t.Execute(w,data)
	
}


func helloEveryone(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "divya.html")
}


func main(){

	http.HandleFunc("/",  myName)
	http.HandleFunc("/mahesh", helloEveryone)
	http.ListenAndServe(":9000",nil)


}