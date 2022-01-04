package main

import "fmt"

func main() {
	x := fmt.Sprintf("%.2f%s", 1.3333222, "%")
	fmt.Println(x)
}
