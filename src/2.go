func main() {
    // Create a slice of numbers and fill it with 10 random integers
    var numbers []int = make([]int, 10)
    for i := 0; i < len(numbers); i++ {
        numbers[i] = rand.Intn(100)
    }
    
    // Print the slice of numbers
    fmt.Println(numbers)
    
    // Create a map of numbers and their frequencies
    frequencyMap := make(map[int]int)
    for _, number := range numbers {
        frequencyMap[number]++
    }
    
    // Print the map of numbers and their frequencies
    fmt.Println(frequencyMap)
}
