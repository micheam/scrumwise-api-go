package wiseman

import (
	"fmt"
	"net/url"
	"time"
)

type NullString struct {
	Str   string
	Valid bool
}

type NullURL struct {
	Str   url.URL
	Valid bool
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

type NullTimeUnit struct {
	TimeUnit TimeUnit
	Valid    bool // Valid is true if Time is not NULL
}

type TimeUnit struct {
	Value float64
	Typ   string
}

func (t TimeUnit) String() string {
	if t.Value == -1 {
		return "none"
	}
	return fmt.Sprintf("%.2f %s", t.Value, t.Typ)
}
