package date

import "time"

const (
	dateformatapi = "01-02-2006T15:04:05Z"
)

//GetNow :
func GetNow() time.Time {
	return time.Now()
}

//GetNowString :
func GetNowString() string {
	now := GetNow()
	return now.Format(dateformatapi)
}
