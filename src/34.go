def find_factorial(n):
    if n < 0:
        return "Factorials are only defined for non-negative integers."
    elif n == 0 or n == 1:
        return 1
    else:
        factorial = 1
        for i in range(2, n + 1):
            factorial *= i
        return factorial

find_factorial(5)
