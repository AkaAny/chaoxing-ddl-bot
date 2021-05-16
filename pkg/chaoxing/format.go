package chaoxing

import (
	"strings"
	"time"
)

const (
	CHAOXING_TIME_FORMAT = "2006-01-02 15:04"
	HOUR_TO_UTC          = 8
)

func ParseTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(CHAOXING_TIME_FORMAT, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	t = t.Local().Add(-HOUR_TO_UTC * time.Hour) //获取到的就是本地时间
	return t, err
}

func Trim(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.TrimSpace(str)
	return str
}
