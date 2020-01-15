package main

import "fmt"
import "./myon"
import "./hueman"

func main() {
		alice := hueman.Reserve("alice")
		fmt.Printf("%s\n", alice.GetName())
		alice.SetName("bill")
		fmt.Printf("%s\n", alice.GetName())
		fmt.Printf("hello, world\n")
		fmt.Printf("2 + 4 = %d\n", myon.Add(2, 4))
		fmt.Printf("-3 - 5 = %d\n", myon.Sub(-3, 5))
		fmt.Printf("5 * 4 = %d\n", myon.Mul(5, 4))
		fmt.Printf("6 / 3 = %d\n", myon.Div(6, 3))
		fmt.Printf("4 / 0 = %d\n", myon.Div(4, 0))
}
