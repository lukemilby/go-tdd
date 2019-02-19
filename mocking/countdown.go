package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}
type DefaultSleeper struct {}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1* time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}


func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}