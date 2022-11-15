package nutil

import (
	"fmt"
	"testing"
	"time"
)

func TestStringConvert(t *testing.T) {

	obj := map[string]string{
		"M":   "asdf",
		"BCC": "asdf",
	}

	t.Run("Convert JSON Pretty", func(t *testing.T) {
		str, err := JSONPretty(obj)
		if err != nil {
			t.Fail()
		}
		fmt.Println(str)
	})

	t.Run("Convert Time To String", func(t *testing.T) {
		str := ConvTimeToString(time.Now())
		if str == "" {
			t.Fail()
		}
		fmt.Println(str)
	})

}
