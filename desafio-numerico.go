package main

import "fmt"

func main() {

	//desafio 1
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("O número %d é divisível por 3.\n", i)
		} else {
			fmt.Printf("O número %d não é divisível por 3.\n", i)
		}
	}

	//desafio 2
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d - PIN\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d - PAN\n", i)
		}
	}
}
