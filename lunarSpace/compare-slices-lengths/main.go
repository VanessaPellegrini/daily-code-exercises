package main

import (
	"fmt"
)

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	a := []string{"123", "444"}
	b := []string{"123", "444"}
	fmt.Printf("slices equal: %t\n", stringSlicesEqual(a, b))

	b = append(b, "test")
	fmt.Printf("slices equal: %t\n", stringSlicesEqual(a, b))

}
