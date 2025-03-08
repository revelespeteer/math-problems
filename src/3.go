func RandomMathProblem() {
	// Generate a random number between 1 and 10
	num1 := rand.Intn(10) + 1
	// Generate another random number between 1 and 10
	num2 := rand.Intn(10) + 1
	// Generate a random operation (+, -, x, /)
	op := []string{"+", "-", "x", "/"}[rand.Intn(4)]
	// Display the problem for the user to solve
	fmt.Printf("What is %d %s %d?\n", num1, op, num2)
}
