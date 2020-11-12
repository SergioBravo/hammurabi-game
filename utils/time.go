package utils

import (
	"fmt"
	"time"
)

// FormatTime ...
func FormatTime(t time.Time) string {
	h, m, s := t.Clock()
	return fmt.Sprintf("%v:%v:%v", h, m, s)

}
