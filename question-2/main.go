package main

import "fmt"

func recursiveFunction(n int) {
	if n == 1 {
		return
	} else {
		recursiveFunction(n / 2)
		fmt.Println(n)
	}

}

func main() {
	recursiveFunction(9)
}
