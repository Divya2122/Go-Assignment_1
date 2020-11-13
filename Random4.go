package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {

    rand.Seed(time.Now().UnixNano())
}

func main() {

    for i := 0; i < 6; i++ {

        fmt.Printf("%d ", rand.Intn(10))
    }

    fmt.Println()
}