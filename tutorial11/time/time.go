package time

import (
	"database/sql/driver"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) Value() (driver.Value, error) {
	return t.Format("2006-01-02 15:04:05"), nil
}

func (t *Time) Scan(value any) error {
	t.Time = value.(time.Time)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 21)
	b = append(b, '"')
	b = t.AppendFormat(b, "2006-01-02 15:04:05")
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = Time{now}
	return
}
