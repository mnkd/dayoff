package dayoff

import (
	"testing"
	"time"
)

func isTrue(ti time.Time, t *testing.T) {
	if IsDayOff(ti) != true {
		t.Error(ti, "is not a day off.")
	}
}

func isFalse(ti time.Time, t *testing.T) {
	if IsDayOff(ti) != false {
		t.Error(ti, "is a day off.")
	}
}

func TestDayoff(t *testing.T) {
	isTrue(time.Date(2016, 12, 31, 12, 0, 0, 0, time.UTC), t)
	isTrue(time.Date(2017, 1, 3, 12, 0, 0, 0, time.UTC), t)

	isFalse(time.Date(2016, 1, 2, 12, 0, 0, 0, time.UTC), t)
	isFalse(time.Date(2016, 11, 4, 12, 0, 0, 0, time.UTC), t)
}
