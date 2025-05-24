def is_even(num):
    return num % 2 == 0

def main():
    # Sample code with even numbers
    print("Even numbers:", [num for num in range(10) if is_even(num)])

if __name__ == "__main__":
    main()
