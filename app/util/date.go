package util

import (
	"time"

	"github.com/pkg/errors"
)

const dateFmtYMD = "2006-01-02"
const dateFmtFull = "2006-01-02 15:04:05"
const dateFmtJS = "2006-01-02T15:04:05Z"

func TimeToYMD(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(dateFmtYMD)
}

func TimeFromYMD(s string) (*time.Time, error) {
	return load(s, dateFmtYMD)
}

func TimeToString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(dateFmtFull)
}

func TimeFromString(s string) (*time.Time, error) {
	return load(s, dateFmtFull)
}

func TimeFromJS(s string) (*time.Time, error) {
	return load(s, dateFmtJS)
}

func load(s string, f string) (*time.Time, error) {
	if len(s) == 0 {
		return nil, nil
	}
	ret, err := time.Parse(f, s)
	if err != nil {
		return nil, errors.New("invalid date [" + s + "] (expected 2020-01-15)")
	}
	return &ret, nil
}
