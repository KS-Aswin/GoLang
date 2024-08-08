package main

import "fmt"

func setValueToZero(value int) {
    value = 0
}

func setPointerToZero(ptr *int) {
    *ptr = 0
}

func main() {
    number := 1
    fmt.Println("initial value:", number)

    setValueToZero(number)
    fmt.Println("after setValueToZero:", number)

    setPointerToZero(&number)
    fmt.Println("after setPointerToZero:", number)

    fmt.Println("memory address of number:", &number)
}
