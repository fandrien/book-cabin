package util

import (
	"fmt"
	"time"
)

func FormatDuration(minutes int) string {
	h := minutes / 60
	m := minutes % 60

	if h == 0 {
		return fmt.Sprintf("%dm", m)
	}

	return fmt.Sprintf("%dh %dm", h, m)
}

// ParseRFC3339 digunakan jika datetime sudah memiliki offset timezone.
// Contoh:
// 2025-12-15T08:00:00+07:00
func ParseRFC3339(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}

// ParseTimeWithTimezone digunakan jika datetime belum memiliki offset,
// tetapi timezone diberikan terpisah.
// Contoh:
// datetime : 2025-12-15T08:00:00
// timezone : Asia/Jakarta
func ParseTimeWithTimezone(
	datetime string,
	timezone string,
) (time.Time, error) {

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(
		"2006-01-02T15:04:05",
		datetime,
		loc,
	)
}

// ParseTimeOffset digunakan jika datetime memiliki offset timezone tanpa titik dua.
// Contoh:
// 2025-12-15T07:15:00+0700
func ParseTimeOffset(value string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-0700", value)
}
