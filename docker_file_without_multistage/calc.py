import sys

def main():
    if len(sys.argv) < 4:
        print("Usage: calc.py <num1> <operator> <num2>")
        print("Operators: +, -, *, /")
        return

    try:
        num1 = float(sys.argv[1])
        operator = sys.argv[2]
        num2 = float(sys.argv[3])
    except ValueError:
        print("Please provide valid numbers.")
        return

    if operator == '+':
        print(f"{num1} + {num2} = {num1 + num2}")
    elif operator == '-':
        print(f"{num1} - {num2} = {num1 - num2}")
    elif operator == '*':
        print(f"{num1} * {num2} = {num1 * num2}")
    elif operator == '/':
        if num2 == 0:
            print("Error: Division by zero.")
        else:
            print(f"{num1} / {num2} = {num1 / num2}")
    else:
        print("Invalid operator. Use one of: +, -, *, /")

if __name__ == "__main__":
    main()
