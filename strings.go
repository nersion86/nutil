package nutil

import (
	"bytes"
	"encoding/json"
	"time"
)

const (
	stringsDefaultTimeFormat = "2006-01-02 15:04:05"
)

//ConvTimeToString = time -> "YYYY-MM0DD hh:mm:ss"
func ConvTimeToString(t time.Time) string {
	return t.Format(stringsDefaultTimeFormat)
}

func ConvStringToTime(str string) (time.Time, error) {
	return time.Parse(stringsDefaultTimeFormat, str)
}

func GetNowString() string {
	return ConvTimeToString(time.Now())
}

func JSONPretty(v any) (string, error) {

	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return "", err
	}

	newBuffer := bytes.NewBuffer(buf)
	if newBuffer == nil {
		return "", NewBasicError("Can't create *bytes.Buffer")
	}

	return newBuffer.String(), nil
}
