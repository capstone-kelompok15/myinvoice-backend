package dateutils

import (
	"strings"
	"time"
)

func StringToDate(date string) (*time.Time, error) {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &time, nil
}

func NowToString() string {
	now := time.Now()
	formatted := now.Format(time.RFC3339Nano)
	return formatted
}

func NowNanoTimeStamp() int64 {
	now := time.Now()
	return now.UnixNano()
}

func TimeToDateStr(date time.Time) string {
	dateStr := strings.Split(date.Local().String(), " ")[0]
	return dateStr
}
