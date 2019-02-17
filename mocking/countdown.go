package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer){
	fmt.Fprint(out, "3")
}
