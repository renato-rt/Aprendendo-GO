package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0
//Primeiro Tipo de for
	for i < 10 {
		i++
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10, j++ {
		fmt.Println("Incrementando j", j)
		time.sleep(time.Second)
	}

}