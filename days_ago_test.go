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
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last year", expected: "2008-01-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last month", expected: "2009-10-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past year", expected: "2008-01-01"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past month", expected: "2009-10-01"},
		{inputTime: time.Date(2017, time.September, 10, 23, 0, 0, 0, time.UTC), inputString: "past week", expected: "2017-09-03"},
		{inputTime: time.Date(2017, time.September, 10, 23, 0, 0, 0, time.UTC), inputString: "last week", expected: "2017-09-03"},
	}

	for _, eachTest := range tests {
		got := GetDateStringFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expected, got, fmt.Sprintf("%s", eachTest.inputString))
	}
}

func TestGetDateStringRangeFromWordString(t *testing.T) {
	tests := []struct {
		inputTime     time.Time
		inputString   string
		expectedStart string
		expectedEnd   string
	}{
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "yesterday", expectedStart: "2009-11-09", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "today", expectedStart: "2009-11-10", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "tomorrow", expectedStart: "2009-11-10", expectedEnd: "2009-11-11"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day after tomorrow", expectedStart: "2009-11-10", expectedEnd: "2009-11-12"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "day before yesterday", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 days ago", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 months ago", expectedStart: "2009-09-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "2 years ago", expectedStart: "2007-01-01", expectedEnd: "2008-12-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 days", expectedStart: "2009-11-08", expectedEnd: "2009-11-10"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 months", expectedStart: "2009-09-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last 2 years", expectedStart: "2007-01-01", expectedEnd: "2008-12-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last year", expectedStart: "2008-01-01", expectedEnd: "2008-12-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "last month", expectedStart: "2009-10-01", expectedEnd: "2009-10-31"},

		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past year", expectedStart: "2008-01-01", expectedEnd: "2008-12-31"},
		{inputTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), inputString: "past month", expectedStart: "2009-10-01", expectedEnd: "2009-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "past week", expectedStart: "2017-09-03", expectedEnd: "2017-09-09"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last week", expectedStart: "2017-09-03", expectedEnd: "2017-09-09"},
	}

	for _, eachTest := range tests {
		gotStart, gotEnd := GetDateStringRangeFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expectedStart, gotStart, fmt.Sprintf("%v Start", eachTest.inputString))
		assert.Equal(t, eachTest.expectedEnd, gotEnd, fmt.Sprintf("%v End", eachTest.inputString))
	}
}
