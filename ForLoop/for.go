package main

import "fmt"

func main() {

	i := 1
	println("This will work like While loop")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	println("This is normal Loop")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	println("------------")
	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
		fmt.Print(" ")
	}
}
