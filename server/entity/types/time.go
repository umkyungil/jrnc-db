package types

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/volatiletech/null"
)

type NullTime struct {
	null.Time
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	var ts int64
	if !t.Time.Valid {
		return json.Marshal(nil)
	}
	ts = t.Time.Time.Unix()
	return json.Marshal(ts)
}

func NullTimeFrom(t time.Time) NullTime {
	return NullTime{
		Time: null.TimeFrom(t),
	}
}

// DATETIME is MySQL's Data form
const (
	DATETIME = "2006-01-02"
)

// ResourceTime is extend time.Time
type ResourceTime struct {
	mysql.NullTime
}

// MarshalJSON is Input JSON
func (rt ResourceTime) MarshalJSON() ([]byte, error) {
	t := rt.Time
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("ResourceTime.MarshalJSON: year outside of range [0,9999]")
	}
	if rt.Time.Format("2006-01-02") == "0001-01-01" {
		return []byte("null"), nil
	}
	return []byte(t.Format(`"` + DATETIME + `"`)), nil
}

// MarshalText is Input TEXT
func (rt ResourceTime) MarshalText() ([]byte, error) {
	t := rt.Time
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}
	return []byte(t.Format(DATETIME)), nil
}

// UnmarshalJSON is Open JSON
func (rt *ResourceTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		rt.Valid = false
		return
	}
	rt.Time, err = time.Parse(DATETIME, s)
	rt.Valid = true

	// JSTに合わせようとして無駄に9時間加えるので減算
	rt.Time = rt.Time.Add(time.Duration(-9) * time.Hour)
	return

}

// Sql driver interface

// Scan する
// func (rt *ResourceTime) Scan(value interface{}) error {
// 	rt.Time = value.(time.Time)
// 	return nil
// }

// // Value 少し変換
// func (rt ResourceTime) Value() (driver.Value, error) {
// 	if rt.Time.Format("2006-01-02") == "0001-01-01" {
// 		return nil, nil
// 	}
// 	return rt.Time, nil
// }
