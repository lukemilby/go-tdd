package selects

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string){
	aDuration := mesureResponseTime(a)
	bDuration := mesureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func mesureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
