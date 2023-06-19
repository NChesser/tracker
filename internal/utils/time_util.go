package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseDuration parses a duration string and returns a time.Duration value.
// The duration string should be in the format "XhYmZs", where X represents hours,
// Y represents minutes, and Z represents seconds.
func ParseDuration(durationString string) (time.Duration, error) {
	parts := strings.Split(durationString, "h")
	if len(parts) != 2 {
		return 0, errors.New("invalid duration format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutesParts := strings.Split(parts[1], "m")
	if len(minutesParts) != 2 {
		return 0, errors.New("invalid duration format")
	}

	minutes, err := strconv.Atoi(minutesParts[0])
	if err != nil {
		return 0, err
	}

	secondsParts := strings.Split(minutesParts[1], "s")
	if len(secondsParts) != 1 {
		return 0, errors.New("invalid duration format")
	}

	seconds, err := strconv.Atoi(secondsParts[0])
	if err != nil {
		return 0, err
	}

	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
	return duration, nil
}
