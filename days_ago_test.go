package daysago

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDateStringFromWordString(t *testing.T) {
	tests := []struct {
		inputTime   time.Time
		inputString string
		expected    string
	}{
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "yesterday", expected: "2009-11-09"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "today", expected: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "tomorrow", expected: "2009-11-11"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day after tomorrow", expected: "2009-11-12"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day before yesterday", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 days ago", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 months ago", expected: "2009-09-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 years ago", expected: "2007-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 days", expected: "2009-11-08"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 months", expected: "2009-09-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 years", expected: "2007-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last year", expected: "2008-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last month", expected: "2009-10-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past year", expected: "2008-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past month", expected: "2009-10-10"},
	}

	for testKey, eachTest := range tests {
		got := GetDateStringFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expected, got, fmt.Sprintf("%d", testKey))
	}
}
