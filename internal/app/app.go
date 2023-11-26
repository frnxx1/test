package internal

import (
	"time"
)




func TimeCounter() int64 {
	d := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

	date := time.Until(d)

	return int64(date.Hours()) / 24
}
