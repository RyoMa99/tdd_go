package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	// Fprintは、io.Writerに第2引数のstringを送る
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	// io.Writerにすることでモックでも下のような標準出力でも引数にとれる
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
