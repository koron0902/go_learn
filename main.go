package main

import "fmt"
import "./myon"

func main() {
		myon.Nyan()
		fmt.Printf("hello, world\n")
		fmt.Printf("2 + 4 = %d\n", myon.Add(2, 4))
		fmt.Printf("-3 - 5 = %d\n", myon.Sub(-3, 5))
}
