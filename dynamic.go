package main

import (
	//"fmt"
	"net/http"
	"html/template"
)

type pageData struct{
	Title string
	Name string
	Age int
	Phone int
	Company string
}

func myName(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("personal.html")
	data := pageData{Title: "Assignment", Name : "Divya", Age: 22 ,Phone: 1234567891, Company : "Kloudone"}
	t.Execute(w,data)
	
}


func helloEveryone(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "divya.html")
}


func main(){

	http.HandleFunc("/",  myName)
	http.HandleFunc("/divya", helloEveryone)
	http.ListenAndServe(":9000",nil)


}