package time

import (
	"time"
)

// Now returns the current time in UTC with no monotonic component.
func Now() time.Time {
	return Canonical(time.Now())
}

// Canonical returns UTC time with no monotonic component.
// Stripping the monotonic component is for time equality.
// See https://github.com/tendermint/tendermint/pull/2203#discussion_r215064334
func Canonical(t time.Time) time.Time {
	return t.Truncate(time.Millisecond).UTC()
}
