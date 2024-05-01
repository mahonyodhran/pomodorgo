package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Pomodorgo!")
	timer_map := createTimers()
	choice := getSessionChoice()
	rounds := getRoundChoice()
	timer := timer_map[choice]
	startSession(rounds, timer)
	fmt.Println("\nSession Completed!")
}

type Timer struct {
	Work int
	Rest int
}

func createTimers() map[int]Timer {
	t1 := Timer{Work: 25, Rest: 5}
	t2 := Timer{Work: 50, Rest: 10}
	t3 := Timer{Work: 100, Rest: 20}
	return map[int]Timer{1: t1, 2: t2, 3: t3}
}

func getSessionChoice() int {
	fmt.Println("\nWhich session would you like?")
	fmt.Println("1: 25 / 5")
	fmt.Println("2: 50 / 10")
	fmt.Println("3: 100 / 20")
	fmt.Print("Choice: ")
	var choice int
	for true {
		_, err := fmt.Scanf("%d", &choice)
		if err == nil && choice > 0 && choice < 4 {
			break
		}
		fmt.Println("Not a valid choice!")
		fmt.Print("Choice: ")
		var dump string
		fmt.Scanln(&dump)
	}
	return choice
}

func getRoundChoice() int {
	fmt.Println("\nHow many rounds would you like?")
	fmt.Println("Choose between 1 - 5")
	fmt.Print("Choice: ")
	var choice int
	for true {
		_, err := fmt.Scanf("%d", &choice)
		if err == nil && choice > 0 && choice < 6 {
			break
		}
		fmt.Println("Not a valid choice!")
		fmt.Print("Choice: ")
		var dump string
		fmt.Scanln(&dump)
	}
	return choice
}

func startSession(rounds int, timer Timer) {
	for i := 0; i < rounds; i++ {
		var round = i + 1
		fmt.Println("\nRound:", round)
		fmt.Println("Let's work.")
		work := time.NewTimer(time.Duration(timer.Work) * time.Minute)
		<-work.C
		fmt.Println("Work Finished - Have a rest!")
		rest := time.NewTimer(time.Duration(timer.Rest) * time.Minute)
		<-rest.C
		fmt.Println("Rest Over!")
	}
}
