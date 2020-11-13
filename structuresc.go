package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 22
    return &p
}

func main() {

    fmt.Println(person{"Anitha", 22})

    fmt.Println(person{name:"abhi", age: 96})

    fmt.Println(person{name: "gopi", age: 18})

    fmt.Println(&person{name: "priya", age: 22})

    fmt.Println(newPerson("krish"))

    s := person{name: "chandu", age: 21}
    fmt.Println(s.name)

    sp := s
    fmt.Println(sp.age)
}