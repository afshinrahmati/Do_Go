package main

import (
	"fmt"
)

func main() {
	//  whoSay string = "afshin"
	whoSay := "abase"
	whoSay = "cas"
	sayHelloWorld(whoSay)
}

func sayHelloWorld(whoSay string) {
	fmt.Println(whoSay)
}
