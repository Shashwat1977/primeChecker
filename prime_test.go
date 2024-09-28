package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_checkPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is Prime."},
		{"not prime", 10, false, "10 is not prime, it is divisble by 2"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative", -1, false, "Negative nos. are not prime, by definition!"},
	}

	for _, test := range primeTests {
		result, message := checkPrime(test.testNum)
		if result != test.expected && test.expected {
			t.Errorf("%d expected to be true, but was false", test.testNum)
		}
		if result != test.expected && !test.expected {
			t.Errorf("%d expected to be false, but was true", test.testNum)
		}
		if test.msg != message {
			t.Errorf("Expected: %v, Actual: %v", test.msg, message)
		}
	}

}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	_ = w.Close()
	inp, _ := io.ReadAll(r)
	os.Stdout = oldOut
	if string(inp) != "-> " {
		t.Errorf("Invalid prompt, expected: -> , Got: %v", string(inp))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldOut
	if !strings.Contains(string(out), "Is it a Prime ?") {
		t.Errorf("Invalid Intro, Got: %v", string(out))
	}
}

func Test_checkNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"prime", "3", "3 is Prime."},
		{"NaN", "abc", "Please enter a correct number."},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
	}

	for _, e := range tests {
		inputReader := strings.NewReader(e.input)
		scanner := bufio.NewScanner(inputReader)

		res, _ := checkNumber(scanner)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("Expected: %s, Got: %s", e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	inputReader := strings.NewReader("1\nq")
	go readUserInput(inputReader, doneChan)
	<-doneChan
	close(doneChan)
}
