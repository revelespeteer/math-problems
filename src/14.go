package main

import "math/rand"

func randomNum() int {
    return rand.Intn(10)
}

func main() {
    num := randomNum()
    fmt.Println("The random number is:", num)
}
