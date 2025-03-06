
func GenerateRandomMathProblem() {
  // Define an array of math operators
  var operators = ["+", "-", "*", "/"]
  
  // Choose a random operator from the array
  var operator = operators[rand.Intn(len(operators))]
  
  // Generate two random numbers between 1 and 10
  var num1 = rand.Intn(10) + 1
  var num2 = rand.Intn(10) + 1
  
  // Define a string to hold the math problem
  var mathProblem = ""
  
  // Build the math problem using the chosen operator and numbers
  switch operator {
    case "+":
      mathProblem = fmt.Sprintf("%d + %d", num1, num2)
    case "-":
      mathProblem = fmt.Sprintf("%d - %d", num1, num2)
    case "*":
      mathProblem = fmt.Sprintf("%d * %d", num1, num2)
    case "/":
      mathProblem = fmt.Sprintf("%d / %d", num1, num2)
  }
  
  // Return the math problem as a string
  return mathProblem
}