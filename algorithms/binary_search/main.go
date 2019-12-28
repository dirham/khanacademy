package main

import (
	"fmt"
	"log"
)

func doSearch(array []int, target int) int {
	min := 0
	max := len(array) - 1
	guess := 0
	for min <= max {
		guess = (max + min) / 2
		if array[guess] == target {
			// for step 2 we need to log guess value on every loop
			log.Println(guess)
			return guess
		} else if array[guess] < target {
			log.Println(guess)
			min = guess + 1
		} else {
			log.Println(guess)
			max = guess - 1
		}
	}
	return -1
}

func main() {
	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
		41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	found := doSearch(primes, 71)
	fmt.Printf("found targe value : %#v on index %#v \n", 71, found)

}
