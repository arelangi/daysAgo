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

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this week", expectedStart: "2017-09-10", expectedEnd: "2017-09-16"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this month", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this year", expectedStart: "2017-01-01", expectedEnd: "2017-12-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "january", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jan", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this january", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for january", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for jan", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in january", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in jan", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this jan", expectedStart: "2017-01-01", expectedEnd: "2017-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev january", expectedStart: "2016-01-01", expectedEnd: "2016-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev jan", expectedStart: "2016-01-01", expectedEnd: "2016-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last january", expectedStart: "2016-01-01", expectedEnd: "2016-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last jan", expectedStart: "2016-01-01", expectedEnd: "2016-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "past january", expectedStart: "2016-01-01", expectedEnd: "2016-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "coming jan", expectedStart: "2018-01-01", expectedEnd: "2018-01-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "next jan", expectedStart: "2018-01-01", expectedEnd: "2018-01-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "february", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "feb", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this february", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for february", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for feb", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in february", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in feb", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this feb", expectedStart: "2017-02-01", expectedEnd: "2017-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev february", expectedStart: "2016-02-01", expectedEnd: "2016-02-29"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev   feb", expectedStart: "2016-02-01", expectedEnd: "2016-02-29"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "past february", expectedStart: "2016-02-01", expectedEnd: "2016-02-29"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last february", expectedStart: "2016-02-01", expectedEnd: "2016-02-29"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: " next  feb", expectedStart: "2018-02-01", expectedEnd: "2018-02-28"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: " upcoming  feb", expectedStart: "2018-02-01", expectedEnd: "2018-02-28"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "march", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "mar", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2018, time.April, 6, 12, 50, 0, 0, time.UTC), inputString: "march", expectedStart: "2018-03-01", expectedEnd: "2018-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this march", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for march", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for mar", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in march", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in mar", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this mar", expectedStart: "2017-03-01", expectedEnd: "2017-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last march", expectedStart: "2016-03-01", expectedEnd: "2016-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev march", expectedStart: "2016-03-01", expectedEnd: "2016-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "coming mar", expectedStart: "2018-03-01", expectedEnd: "2018-03-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "upcoming mar", expectedStart: "2018-03-01", expectedEnd: "2018-03-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "april", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "apr", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for april", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for apr", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in april", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in apr", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this april", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this apr", expectedStart: "2017-04-01", expectedEnd: "2017-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last april", expectedStart: "2016-04-01", expectedEnd: "2016-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev apr", expectedStart: "2016-04-01", expectedEnd: "2016-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last april", expectedStart: "2016-04-01", expectedEnd: "2016-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev apr", expectedStart: "2016-04-01", expectedEnd: "2016-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "coming april", expectedStart: "2018-04-01", expectedEnd: "2018-04-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "next apr", expectedStart: "2018-04-01", expectedEnd: "2018-04-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "may", expectedStart: "2017-05-01", expectedEnd: "2017-05-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in may", expectedStart: "2017-05-01", expectedEnd: "2017-05-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for may", expectedStart: "2017-05-01", expectedEnd: "2017-05-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this may", expectedStart: "2017-05-01", expectedEnd: "2017-05-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev may", expectedStart: "2016-05-01", expectedEnd: "2016-05-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "upcoming may", expectedStart: "2018-05-01", expectedEnd: "2018-05-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "june", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jun", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for june", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for jun", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in june", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in jun", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this june", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this jun", expectedStart: "2017-06-01", expectedEnd: "2017-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last jun", expectedStart: "2016-06-01", expectedEnd: "2016-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev june", expectedStart: "2016-06-01", expectedEnd: "2016-06-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "next june", expectedStart: "2018-06-01", expectedEnd: "2018-06-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "july", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "jul", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in july", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in jul", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for july", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for jul", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this july", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this jul", expectedStart: "2017-07-01", expectedEnd: "2017-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last jul", expectedStart: "2016-07-01", expectedEnd: "2016-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev jul", expectedStart: "2016-07-01", expectedEnd: "2016-07-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "upcoming jul", expectedStart: "2018-07-01", expectedEnd: "2018-07-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "august", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "aug", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for august", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for aug", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in august", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in aug", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this august", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this aug", expectedStart: "2017-08-01", expectedEnd: "2017-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last august", expectedStart: "2016-08-01", expectedEnd: "2016-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev aug", expectedStart: "2016-08-01", expectedEnd: "2016-08-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "coming aug", expectedStart: "2018-08-01", expectedEnd: "2018-08-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "september", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "sep", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this     september", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this     sep", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for     september", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for     sep", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in     september", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in     sep", expectedStart: "2017-09-01", expectedEnd: "2017-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last sep", expectedStart: "2016-09-01", expectedEnd: "2016-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev   sep", expectedStart: "2016-09-01", expectedEnd: "2016-09-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "upcoming   sep", expectedStart: "2018-09-01", expectedEnd: "2018-09-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "october", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "oct", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for october    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for oct    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this october    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this oct    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in october    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in oct    ", expectedStart: "2017-10-01", expectedEnd: "2017-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last oct", expectedStart: "2016-10-01", expectedEnd: "2016-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev oct", expectedStart: "2016-10-01", expectedEnd: "2016-10-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "coming oct", expectedStart: "2018-10-01", expectedEnd: "2018-10-31"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "november", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "nov", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  for   november   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  for   nov   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  this   november   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  this   nov   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  in   november   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "  in   nov   ", expectedStart: "2017-11-01", expectedEnd: "2017-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last nov", expectedStart: "2016-11-01", expectedEnd: "2016-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev nov", expectedStart: "2016-11-01", expectedEnd: "2016-11-30"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "next nov", expectedStart: "2018-11-01", expectedEnd: "2018-11-30"},

		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "december", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "dec", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for december", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "for dec", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in december", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "in dec", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this december", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "this dec", expectedStart: "2017-12-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "prev dec", expectedStart: "2016-12-01", expectedEnd: "2016-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "last dec", expectedStart: "2016-12-01", expectedEnd: "2016-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "upcoming dec  ", expectedStart: "2018-12-01", expectedEnd: "2018-12-31"},

		//Years
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "2017", expectedStart: "2017-01-01", expectedEnd: "2017-12-31"},
		{inputTime: time.Date(2017, time.September, 12, 23, 0, 0, 0, time.UTC), inputString: "2018", expectedStart: "2018-01-01", expectedEnd: "2018-12-31"},
	}

	for _, eachTest := range tests {
		gotStart, gotEnd := GetDateStringRangeFromWordString(eachTest.inputTime, eachTest.inputString)
		assert.Equal(t, eachTest.expectedStart, gotStart, fmt.Sprintf("%v Start", eachTest.inputString))
		assert.Equal(t, eachTest.expectedEnd, gotEnd, fmt.Sprintf("%v End", eachTest.inputString))
	}
}
