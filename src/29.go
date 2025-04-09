def is_odd(x):
    if x % 2 == 1:
        return True
    else:
        return False

def main():
    num = int(input("请输入一个整数: "))
    print(f"奇数部分是 {is_odd(num)}")

if __name__ == "__main__":
    main()
