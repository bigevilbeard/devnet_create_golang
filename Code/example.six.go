package main

import "fmt"
import "time"

func main() {

	for i := 3; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println("Launch Rocket!")
}
