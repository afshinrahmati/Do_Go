package main

import (
	"bufio"
	"fmt"
	"os"

	"myapp/doctor"
)

var s = "Japan"

func main() {
	var whatTosay string
	whatTosay = doctor.Intro()
	fmt.Println(whatTosay)

	reader := bufio.NewReader(os.Stdin)

	userInput , _ := reader.ReadString('\n')
	fmt.Println(userInput  )

}