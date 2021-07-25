package repository

import (
	"time"

	"github.com/volatiletech/null/v8"
)

// StringMayNull returns null.String.
// If s is empty string,  this returns as Null.
// Is s is not empty, this returns as String.
func StringMayNull(s string) null.String {
	if s == "" {
		return null.NewString("", false)
	}

	return null.StringFrom(s)
}

// IntMayNull returns null.Int.
// If i is zero,  this returns as Null.
// Is i is not zero, this returns as Int.
func IntMayNull(i int) null.Int {
	if i == 0 {
		return null.NewInt(0, false)
	}

	return null.IntFrom(i)
}

// Int64MayNull returns null.Int64.
// If i is zero,  this returns as Null.
// Is i is not zero, this returns as Int64s.
func Int64MayNull(i int64) null.Int64 {
	if i == 0 {
		return null.NewInt64(0, false)
	}

	return null.Int64From(i)
}

// TimeMayNull returns null.Time.
// If i is zero,  this returns as Null.
// Is i is not zero, this returns as Time.
func TimeMayNull(t time.Time) null.Time {
	if t.IsZero() {
		return null.NewTime(t, false)
	}

	return null.TimeFrom(t)
}
