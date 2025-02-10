package time_util

import (
	"fmt"
	"time"
)

func Convert(date_input string, time_input string) (int64, error) {
	dateTimeStr := fmt.Sprintf("%s %s:00", date_input, time_input)
	parsed, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)

	if err != nil {
		return -1, err
	}

	return parsed.Unix(), nil
}
