# pomodoro9000
Command line [pomodoro timer](https://en.wikipedia.org/wiki/Pomodoro_Technique), implemented in Go.

## Installation

For Go 1.17 or later:

go install github.com/xyproto/pomodoro9000@latest

## Usage
Usage of pomodoro9000:

    pomodoro9000 [options] [finish time]

Duration defaults to 25 minutes. Finish may be a duration (e.g. "1h2m3s")
or a target time (e.g. "1:00pm" or "13:02:03"). Durations may be expressed
as integer minutes (e.g. "15") or time with units (e.g. "1m30s" or "90s").

Play a bell sound at the end of the timer, unless -silence is set.

## Screenshots
```bash
$ pomodoro9000 -simple
Start timer for 25m0s.

Countdown: 24:43

$ pomodoro9000 -h
Usage of pomodoro9000:

    pomodoro9000 [options] [duration]

Duration defaults to 25 minutes. Durations may be expressed as integer minutes
(e.g. "15") or time with units (e.g. "1m30s" or "90s").

Play a bell sound at the end of the timer, unless -silence is set.
  -silence
        Don't ring bell after countdown
  -simple
        Display simple countdown
```

![screenshot](./screenshot.png)
