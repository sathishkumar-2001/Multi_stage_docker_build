package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: calc <num1> <operator> <num2>")
        fmt.Println("Operators: +, -, *, /")
        return
    }

    num1, err1 := strconv.ParseFloat(os.Args[1], 64)
    operator := os.Args[2]
    num2, err2 := strconv.ParseFloat(os.Args[3], 64)

    if err1 != nil || err2 != nil {
        fmt.Println("Please provide valid numbers.")
        return
    }

    switch operator {
    case "+":
        fmt.Printf("%f + %f = %f\n", num1, num2, num1+num2)
    case "-":
        fmt.Printf("%f - %f = %f\n", num1, num2, num1-num2)
    case "*":
        fmt.Printf("%f * %f = %f\n", num1, num2, num1*num2)
    case "/":
        if num2 == 0 {
            fmt.Println("Error: Division by zero.")
        } else {
            fmt.Printf("%f / %f = %f\n", num1, num2, num1/num2)
        }
    default:
        fmt.Println("Invalid operator. Use one of: +, -, *, /")
    }
}
