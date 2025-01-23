package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"time"
)

func main() {
	for {
		var input string
		fmt.Println("Welcome to the Guessing Game!")
		fmt.Println("1) Play")
		fmt.Println("2) Leaderboard")
		fmt.Println("3) Quit")
		fmt.Print("Input choice: ")
		_, err := fmt.Scanln(&input)

		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case "1":
			guessing_game()
		case "2":
			leaderboard()
		case "3":
			break
		default:
			fmt.Println("Wrong number chosen. Please choose from the menu listed")
		}
	}
}

func guessing_game() {
	start := time.Now()
	answer := rand.Int64N(10) + 1

	var guess string

	fmt.Println(answer)
	fmt.Println("Guess a number from 1-10")

	for {
		_, err := fmt.Scanln(&guess)
		if err != nil {
			log.Fatal(err)
		}
		guess, err := strconv.ParseInt(guess, 10, 32)
		if guess > answer {
			fmt.Println("Too High")
		} else if guess < answer {
			fmt.Println("Too Low")
		} else {
			elapsed := time.Since(start)
			fmt.Println("You Got It Right!")
			record("Kira Midru", elapsed.String())
			return
		}
	}
}

func leaderboard() {

}

func record(name string, time string) {

}
