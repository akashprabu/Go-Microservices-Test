package datastructures

import "gobasics/data-structures/helpers"

const numPool = 21

func Channel(intChan chan int) {
	value := helpers.GenerateRandomNumber(numPool)
	intChan <- value
}
