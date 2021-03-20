package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

//Fake sleeper, only counts calls, doesn't actually sleep
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

//This one actually sleeps
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()

		fmt.Fprintln(out, i)

	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func main() {
	sleepr := &DefaultSleeper{}
	Countdown(os.Stdout, sleepr)
}