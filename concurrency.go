package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	defer func () {
		fmt.Println(time.Since(startTime))
	}()
	evilNinjas := []string{"tommy", "johnny", "arthur", "polly"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}

	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at target:", target)
	time.Sleep(time.Second)
}
