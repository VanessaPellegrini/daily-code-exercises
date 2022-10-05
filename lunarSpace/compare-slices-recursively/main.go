package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	nutrition string
	race      string
	weigth    int
	age       int
}

func firstHit() {
	animal := Animal{
		nutrition: "carnivore",
		race:      "bulldog",
		weigth:    8,
		age:       20,
	}

}

func main() {
	a := []string{"12345", "321234"}
	b := []string{"12345", "321234"}
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))

	b = append(b, "test")
	fmt.Printf("slices equal: %t\n", reflect.DeepEqual(a, b))
}
