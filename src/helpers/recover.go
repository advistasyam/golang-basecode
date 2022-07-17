package helpers

import "fmt"

func Recover(name string) {
	if r := recover(); r != nil {
		fmt.Printf("Recovered! (%v)", name)
	}
}
