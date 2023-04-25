[![Go Report Card](https://goreportcard.com/badge/github.com/tomyfalgui/meeting_meter)](https://goreportcard.com/report/github.com/tomyfalgui/meeting_meter)
[![Go Reference](https://pkg.go.dev/badge/github.com/tomyfalgui/meeting_meter.svg)](https://pkg.go.dev/github.com/tomyfalgui/meeting_meter)

# “The Meeting Meter”.

Meetings are expensive, and we can work out exactly how
much each minute of the meeting costs if we know the approximate salary
or billing rate of all the participants.

`meeting` is a command-line tool for calculating the cost of an ongoing meeting

## Installation

```bash
go install github.com/tomyfalgui/meeting_meter/cmd/meeting@latest
```

## Usage
```bash
meeting [OPTIONS] [ARGS]
```

### Options

```bash
-f frequency of printing (time duration string)
```

### Arguments

Arguments is a list of int which represent the hourly cost in cents of a meeting's participants.
