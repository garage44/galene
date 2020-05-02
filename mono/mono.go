package mono

import (
	"time"
)

var epoch = time.Now()

func fromDuration(d time.Duration, hz uint32) uint64 {
	return uint64(d) * uint64(hz) / uint64(time.Second)
}

func toDuration(tm uint64, hz uint32) time.Duration {
	return time.Duration(tm * uint64(time.Second) / uint64(hz))
}

func Now(hz uint32) uint64 {
	return fromDuration(time.Since(epoch), hz)
}

func Microseconds() uint64 {
	return Now(1000000)
}

var ntpEpoch = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

func NTPToTime(ntp uint64) time.Time {
	sec := uint32(ntp >> 32)
	frac := uint32(ntp & 0xFFFFFFFF)
	return ntpEpoch.Add(
		time.Duration(sec) * time.Second +
			((time.Duration(frac) * time.Second) >> 32),
	)
}

func TimeToNTP(tm time.Time) uint64 {
	d := tm.Sub(ntpEpoch)
	sec := uint32(d / time.Second)
	frac := uint32(d % time.Second)
	return (uint64(sec) << 32) + (uint64(frac) << 32) / uint64(time.Second)
}
