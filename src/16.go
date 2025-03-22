import sympy as sp

# Define symbols
x, y = sp.symbols('x y')

# Solve the equation 2*x + 3*y = 10
solution = sp.solve((2*x + 3*y - 10), (x, y))
print(solution)
