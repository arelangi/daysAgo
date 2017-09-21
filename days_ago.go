package daysago

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

const layout = "2006-01-02"

//GetDateStringFromWordString takes as input a sentence or word and returns a time for that
func GetDateStringFromWordString(refTime time.Time, input string) (out string) {
	input = strings.ToLower(input)
	switch input {
	case "yesterday":
		out = refTime.AddDate(0, 0, -1).Format(layout)
	case "today":
		out = refTime.Format(layout)
	case "tomorrow":
		out = refTime.AddDate(0, 0, 1).Format(layout)
	case "day after tomorrow":
		out = refTime.AddDate(0, 0, 2).Format(layout)
	case "day before yesterday":
		out = refTime.AddDate(0, 0, -2).Format(layout)
	default:
		re := regexp.MustCompile("[0-9]+")
		numArr := re.FindStringSubmatch(input)
		if len(numArr) > 0 {
			num, _ := strconv.Atoi(numArr[0])
			if strings.Contains(input, "day") {
				out = refTime.AddDate(0, 0, -num).Format(layout)
			} else if strings.Contains(input, "month") {
				out = refTime.AddDate(0, -num, 0).Format(layout)
			} else if strings.Contains(input, "years") {
				out = refTime.AddDate(-num, 0, 0).Format(layout)
			}
		}
	}
	return
}
