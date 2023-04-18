# “The Meeting Meter”.

Meetings are expensive, and we can work out exactly how
much each minute of the meeting costs if we know the approximate salary
or billing rate of all the participants.

Given this (perhaps configurable for regular meetings),
the Meeting Meter produces a constantly-updated display of the current cost 
of the meeting.

To take a simple example,
if there are four participants,
  each on an equivalent hourly rate of $60,
  then the meeting costs $4/minute.

As the clock ticks away,
the Meeting Meter will show how much has been spent so far,
encouraging people to make the best use of each other’s time.

## Installation

```bash
go install github.com/tomyfalgui/meeting_meter/cmd/meeting@latest
```

## Usage
```bash
meeting_meter
```
