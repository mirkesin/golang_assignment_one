package main

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	fetchSig := fetchSignalInstance()

	start := time.Now()
	go func() {
		for {
			switch {
			case <-fetchSig:
				if time.Now().Sub(start).Nanoseconds() < 990000000 {
					t.Log("There exists a two crawls who were executed less than 1 sec apart.")
					t.Log("Solution is incorrect.")
					t.FailNow()
				}
				start = time.Now()
			}
		}
	}()

	main()
}
