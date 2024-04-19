package main

import (
	"fmt"
	"time"
)

func main2() {
	startTime := time.Now()

	defer func () {
		fmt.Println(time.Since(startTime))
	}()

	smokeSignal := make(chan bool)

	evilNinjas := []string{"tommy", "johnny", "arthur", "polly"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, smokeSignal)
		<-smokeSignal
	}
}

func attack(target string, smokeSignalFired chan bool) {
	fmt.Println("Throwing ninja stars at target:", target)
	smokeSignalFired <- true
}
