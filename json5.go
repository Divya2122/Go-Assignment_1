package main


import (
	    "fmt"
		"encoding/json"
		)	

type Preson struct{
	Name string `json:"name"`
	Phone int   `json:"phone"`
	Age int     `json:"age"`
	Gender string `json:"gender"`
	City []string `json:"city"`
}

func main () {

	user := &Preson{Name :"Divya",Phone:9876543210,Age:22,Gender: "female", City: []string{"madhanapalle", "Andhra Pradesh"}}
	data, _:= json.Marshal(user)
	fmt.Println(string(data))


	 user2 := `{"Name":"chandu reddy","Phone":1234567890,"Age":21,"Gender":"feMale","City":["angallu","Andhra Pradesh"]}`
	  data2 := &Preson{}

	  json.Unmarshal([]byte(user2),data2)
	  fmt.Println(data2)
}