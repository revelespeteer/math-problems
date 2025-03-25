import random
from functools import reduce

def game():
    numbers = [1, 2, 3, 4, 5]
    result = reduce(lambda x, y: x * y, (random.choice(numbers) for _ in range(6)))
    return "You got a lucky number! The sum of the chosen numbers is: {}".format(result)

game()
