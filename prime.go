package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}
	if n < 0 {
		return false, "Negative nos. are not prime, by definition!"
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime, it is divisble by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is Prime.", n)
}

func main() {
	intro()                              // Starts the intro for our primeChecker program
	doneChan := make(chan bool)          // Channel to track for program exit
	go readUserInput(os.Stdin, doneChan) // Another goroutine to keep running the prime checker
	<-doneChan                           // main goroutine will be blocked until it receives a message on this channel
	close(doneChan)
}

func intro() {
	fmt.Println("Is it a Prime ?")
	fmt.Println("---------------")
	fmt.Println("Enter a whole number to check for prime. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumber(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a correct number.", false
	}
	_, msg := checkPrime(numToCheck)
	return msg, false
}
