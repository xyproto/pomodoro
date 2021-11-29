package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

func fullscreenCountdown(start, finish time.Time, formatter func(time.Duration) string) {
	err := termbox.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't open display:", err)
		os.Exit(2)
	}
	defer termbox.Close()

	// Leaks a goroutine
	ticker := time.Tick(40 * time.Millisecond)
	quit := make(chan struct{})
	// Leaks if not quit
	go func() {
		defer close(quit)
		for {
			e := termbox.PollEvent()
			// Quit on any of the common keys for quitting
			if strings.ContainsRune("CcDdQqXx", e.Ch) ||
				e.Key == termbox.KeyCtrlC ||
				e.Key == termbox.KeyCtrlD ||
				e.Key == termbox.KeyCtrlQ ||
				e.Key == termbox.KeyCtrlX {
				return
			}
		}
	}()

	for render(start, finish, formatter) {
		select {
		case <-ticker:
		case <-quit:
			termbox.Close()
			os.Exit(1)
			return
		}
	}

}

func render(start, finish time.Time, formatter func(time.Duration) string) bool {
	now := time.Now()
	remaining := -now.Sub(finish)
	if remaining < 0 {
		return false
	}

	const timeFmt = "3:04:05pm"
	screenW, screenH := termbox.Size()
	centerX := screenW / 2
	centerY := screenH / 2

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	startStr := start.Format(timeFmt)
	Point{
		0, 0,
		termbox.ColorBlue, termbox.ColorDefault,
	}.Str("Start")
	Point{
		0, 1,
		termbox.ColorWhite, termbox.ColorDefault,
	}.Str(startStr)

	nowStr := now.Format(timeFmt)
	Point{
		centerX - (len("Now") / 2), 0,
		termbox.ColorBlue, termbox.ColorDefault,
	}.Str("Now")
	Point{
		centerX - (len(nowStr) / 2), 1,
		termbox.ColorWhite, termbox.ColorDefault,
	}.Str(nowStr)

	finishStr := finish.Format(timeFmt)
	Point{
		screenW - len("Finish"), 0,
		termbox.ColorBlue, termbox.ColorDefault,
	}.Str("Finish")
	Point{
		screenW - len(finishStr), 1,
		termbox.ColorWhite, termbox.ColorDefault,
	}.Str(finishStr)

	remainingStr := formatter(remaining)
	Point{
		centerX - (len(remainingStr) * (BigCharWidth + 1) / 2), centerY,
		termbox.ColorBlue, termbox.ColorDefault,
	}.BigStr(remainingStr)

	Point{
		0, centerY + 6,
		termbox.ColorBlue, termbox.ColorWhite,
	}.ProgressBar(screenW, int(start.Sub(now)), int(start.Sub(finish)))

	termbox.Flush()
	return true
}
