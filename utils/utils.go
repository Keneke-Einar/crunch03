package utils

import "fmt"

func PrintTen() {
	for i := range 10 {
		fmt.Println(i+1)
	}
}