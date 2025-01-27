package main

import (
	"net/http"
	"fmt"
	"strconv"
	"time"
	"math/rand"
)
func main() {
	for i := 100000; i < 1000000; i++ {
		pin := rand.Intn(1000000)
		now := time.Now()
		nowUnix := now.Unix()
		x, err := http.Get("https://kahoot.it/reserve/session/" + strconv.Itoa(pin) + "/?" + strconv.FormatInt(nowUnix, 10))
		if err != nil {
			panic("Error thrown during game pin checking.")
		}
		if x.StatusCode != 404 {
			fmt.Println("Game PIN: " + strconv.Itoa(pin) + " - " + x.Status)
		}
	}
}
