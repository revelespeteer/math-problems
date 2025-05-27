def find_sum(numbers):
    total = 0
    for num in numbers:
        if isinstance(num, int) and num > 0:
            total += num
    return total

numbers = [1, -2, 3, -4, 5]
print(find_sum(numbers))
