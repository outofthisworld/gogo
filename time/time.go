package time

import "fmt"

const MS = 1000
const MSE = 1_000_000
const NS = 1_000_000_000

const MILLISECONDS_STRING = "Milliseconds"
const MICROSECONDS_STRING = "Microseconds"
const NANOSECONDS_STRING = "Nanoseconds"
const SECONDS_STRING = "Seconds"

type Time struct {
	amount uint32
	unit   string
}
type TimeSeconds Time
type TimeMilliseconds Time
type TimeNanoSeconds Time
type TimeMicroSeconds Time

type TimeValue interface {
	ToSeconds() TimeValue
	ToMillseconds() TimeValue
	ToNanoSeconds() TimeValue
	ToMicroSeconds() TimeValue
	ToTime() Time
}

func Seconds(time uint32, formatString *string) TimeSeconds {
	return TimeSeconds(Time{time, SECONDS_STRING})
}

func Millseconds(time uint32) TimeMilliseconds {
	return TimeMilliseconds(Time{time, MILLISECONDS_STRING})
}

func NanoSeconds(time uint32) TimeNanoSeconds {
	return TimeNanoSeconds(Time{time, NANOSECONDS_STRING})
}

func Microseconds(time uint32) TimeMicroSeconds {
	return TimeMicroSeconds(Time{time, MICROSECONDS_STRING})
}

/*
	TimeSeconds
*/
func (s TimeSeconds) ToSeconds() TimeValue {
	return TimeValue(s)
}
func (s TimeSeconds) ToMillseconds() TimeValue {
	return TimeMilliseconds(Time{s.amount * MS, MILLISECONDS_STRING})
}

func (s TimeSeconds) ToNanoSeconds() TimeValue {
	return TimeNanoSeconds((Time{s.ToSeconds().ToTime().amount * NS, NANOSECONDS_STRING}))
}

func (s TimeSeconds) ToMicroSeconds() TimeValue {
	return TimeMicroSeconds(Time{s.ToSeconds().ToTime().amount * MSE, MICROSECONDS_STRING})
}

func (s TimeSeconds) ToTime() Time {
	return Time(s)
}

/*
	TimeNanoSeconds
*/
func (s TimeNanoSeconds) ToSeconds() TimeValue {
	return TimeSeconds(Time{s.amount / NS, SECONDS_STRING})
}

func (s TimeNanoSeconds) ToMillseconds() TimeValue {
	return TimeMilliseconds(Time{s.ToSeconds().ToTime().amount * MS, MILLISECONDS_STRING})
}

func (s TimeNanoSeconds) ToNanoSeconds() TimeValue {
	return TimeNanoSeconds(s)
}

func (s TimeNanoSeconds) ToMicroSeconds() TimeValue {
	return TimeMicroSeconds(Time{s.ToSeconds().ToTime().amount * MSE, MICROSECONDS_STRING})
}

func (s TimeNanoSeconds) ToTime() Time {
	return Time(s)
}

/*
	TimeMicroSeconds
*/
func (s TimeMicroSeconds) ToSeconds() TimeValue {
	return TimeSeconds(Time{s.amount / MSE, SECONDS_STRING})
}

func (s TimeMicroSeconds) ToMillseconds() TimeValue {
	return TimeMilliseconds(Time{s.ToSeconds().ToTime().amount * MS, MILLISECONDS_STRING})
}

func (s TimeMicroSeconds) ToNanoSeconds() TimeValue {
	return TimeNanoSeconds(Time{s.ToSeconds().ToTime().amount * NS, NANOSECONDS_STRING})
}

func (s TimeMicroSeconds) ToMicroSeconds() TimeValue {
	return TimeMicroSeconds(s)
}

func (s TimeMicroSeconds) ToTime() Time {
	return Time(s)
}

/*
	TimeMilliseconds
*/
func (s TimeMilliseconds) ToSeconds() TimeValue {
	return TimeSeconds(Time{s.amount / MS, SECONDS_STRING})
}

func (s TimeMilliseconds) ToMillseconds() TimeValue {
	return TimeMilliseconds(s)
}

func (s TimeMilliseconds) ToNanoSeconds() TimeValue {
	return TimeNanoSeconds(Time{s.ToSeconds().ToTime().amount * NS, NANOSECONDS_STRING})
}

func (s TimeMilliseconds) ToMicroSeconds() TimeValue {
	return TimeMicroSeconds(Time{s.ToSeconds().ToTime().amount * MSE, MICROSECONDS_STRING})
}

func (s TimeMilliseconds) ToTime() Time {
	return Time(s)
}

/*
	Time
*/
func (s Time) AsInt() uint32 {
	return uint32(s.amount)
}

func (s Time) AsString() string {
	return fmt.Sprintf("%d %s", s.amount, s.unit)
}
