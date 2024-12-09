package main

import "log"

func main() {
	var myColor string
	myColor = "Green"
	log.Println("myColor is:", myColor)
	changeColor(&myColor)
	log.Println("myColor is:", myColor)

}

func changeColor(color *string) {
	*color = "Red"
	// Test
}
