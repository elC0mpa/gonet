package common

import (
	"fmt"
	"math"
)

func FromBytesToString(bytes float64) string {
	var value float64
	var unit string

	if bytes < 1024 {
		value = bytes
		unit = "B"
	} else if bytes < (math.Pow(1024, 2)) {
		value = bytes / 1024
		unit = "KB"
	} else if bytes < (math.Pow(1024, 3)) {
		value = bytes / math.Pow(1024, 2)
		unit = "MB"
	} else if bytes < (math.Pow(1024, 4)) {
		value = bytes / math.Pow(1024, 3)
		unit = "GB"
	}

	if value == float64(int(value)) {
		return fmt.Sprintf("%d %s", int(value), unit)
	}

	return fmt.Sprintf("%.1f %s", value, unit)
}
