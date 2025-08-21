package utils

import "strconv"

func ParseInt(val string, defaultValue int) int {
	if v, err := strconv.Atoi(val); err == nil {
		return v
	}
	return defaultValue
}

func ParseFloat(val string, defaultValue float64) float64 {
	if v, err := strconv.ParseFloat(val, 64); err == nil {
		return v
	}
	return defaultValue
}
