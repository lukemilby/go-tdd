package main

import (
	"bytes"
	"testing"
)



func TestCountdown(t *testing.T){
	t.Run("print 3 and Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		assertCountdownOutput(t, got, want)
		assertSleepCalls(t, spySleeper, 4)
	})
}
func assertCountdownOutput(t *testing.T, got string, want string){
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertSleepCalls(t *testing.T, s *SpySleeper, want int){
	t.Helper()
	if s.Calls != want {
		t.Errorf("not enought calls to sleeper, want '%d' got '%d'", want, s.Calls)
	}
}

type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)