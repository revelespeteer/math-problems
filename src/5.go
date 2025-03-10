
package main
import "math/rand"
func main() {
    rand.Seed(9)
    n := 100 + rand.Intn(900)
    fmt.Println("What is", n, "times", n, "?")
}