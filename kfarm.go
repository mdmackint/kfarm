package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
    "github.com/common-nighthawk/go-figure"
)

var stage uint8 = 1

func colour(sc int) string {
	switch sc {
	case 200:
		return "green"
	case 500:
		return "red"
	default:
		return "purple"
	}
}

func progress() {
	for range time.Tick(500 * time.Millisecond) {
		switch stage {
		case 1:
			fmt.Print("\r-")
		case 2:
			fmt.Print("\r\\")
		case 3:
			fmt.Print("\r|")
		case 4:
			fmt.Print("\r/")
		}
		if stage != 4 {
			stage++
		} else {
			stage = 1
		}
	}
}
func main() {
	fmt.Println("Good luck!")
	go progress()
	for i := 100000; i < 1000000; i++ {
		pin := rand.Intn(1000000)
		now := time.Now()
		nowUnix := now.Unix()
		x, err := http.Get("https://kahoot.it/reserve/session/" + strconv.Itoa(pin) + "/?" + strconv.FormatInt(nowUnix, 10))
		if err != nil {
			panic("Error thrown during game pin checking.")
		}
		if x.StatusCode != 404 {
			fmt.Print("\r")
			ascii := figure.NewColorFigure(strconv.Itoa(pin),"",colour(x.StatusCode),true)
            ascii.Print()
		}
	}
}
