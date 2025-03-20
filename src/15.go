package main

import "math/rand"

func generateRandomMathProblem() (int, int) {
    // Generate two random numbers between 1 and 10
    num1 := rand.Intn(10) + 1
    num2 := rand.Intn(10) + 1
    operation := rand.Intn(3)
    switch operation {
        case 0: // addition
            return num1, num2
        case 1: // subtraction
            return num1 - num2, num2
        case 2: // multiplication
            return num1 * num2, num2
        default: // division
            if num2 != 0 {
                return num1 / num2, num2
            } else {
                return generateRandomMathProblem()
            }
    }
}
