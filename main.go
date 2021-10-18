package main

import (
	"fmt"
	"time"
)

func main() {
	ticker(10)
}

func ticker(x int) {
	ticker := time.NewTicker(time.Duration(x) * time.Second)

	for t := range ticker.C {
		fmt.Println(t.Second())
	}

}
