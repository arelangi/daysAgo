package daysago

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
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
	case "last month", "past month":
		out = now.New(refTime.AddDate(0, -1, 0)).BeginningOfMonth().Format(layout)
	case "last year", "past year":
		out = now.New(refTime.AddDate(-1, 0, 0)).BeginningOfYear().Format(layout)
	case "past week", "last week":
		out = now.New(refTime.AddDate(0, 0, -7)).BeginningOfWeek().Format(layout)
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

//GetDateStringRangeFromWordString takes as input a sentence or word and returns a range of time as start and end
func GetDateStringRangeFromWordString(refTime time.Time, input string) (start, end string) {
	input = strings.TrimSpace(strings.ToLower(input))
	fmt.Println("Bitch, we are looking at ", input)
	switch input {
	case "yesterday":
		end = refTime.Format(layout)
		start = refTime.AddDate(0, 0, -1).Format(layout)
	case "today":
		end = refTime.Format(layout)
		start = end
	case "tomorrow":
		start = refTime.Format(layout)
		end = refTime.AddDate(0, 0, 1).Format(layout)
	case "day after tomorrow":
		start = refTime.Format(layout)
		end = refTime.AddDate(0, 0, 2).Format(layout)
	case "day before yesterday":
		start = refTime.AddDate(0, 0, -2).Format(layout)
		end = refTime.Format(layout)
	case "last month", "past month":
		refTime = now.New(refTime.AddDate(0, -1, 0)).BeginningOfMonth()
		start = refTime.Format(layout)
		end = now.New(refTime).EndOfMonth().Format(layout)
	case "last year", "past year":
		refTime = now.New(refTime.AddDate(-1, 0, 0)).BeginningOfYear()
		start = refTime.Format(layout)
		end = refTime.AddDate(1, 0, -1).Format(layout)
	case "past week", "last week":
		refTime = now.New(refTime.AddDate(0, 0, -7)).BeginningOfWeek()
		start = refTime.Format(layout)
		end = now.New(refTime).EndOfWeek().Format(layout)
	case "this week":
		start = now.New(refTime).BeginningOfWeek().Format(layout)
		end = now.New(refTime).EndOfWeek().Format(layout)
	case "this month":
		start = now.New(refTime).BeginningOfMonth().Format(layout)
		end = now.New(refTime).EndOfMonth().Format(layout)
	case "this year":
		start = now.New(refTime).BeginningOfYear().Format(layout)
		end = now.New(refTime).EndOfYear().Format(layout)
	case "january", "jan", "this january", "this jan":
		start = now.New(refTime).BeginningOfYear().Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 1, -1).Format(layout)
	case "february", "feb", "this february", "this feb":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 1, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 2, -1).Format(layout)
	case "march", "mar", "this march", "this mar":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 2, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 3, -1).Format(layout)
	case "april", "apr", "this april", "this apr":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 3, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 4, -1).Format(layout)
	case "may", "this may":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 4, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 5, -1).Format(layout)
	case "june", "jun", "this june", "this jun":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 5, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 6, -1).Format(layout)
	case "july", "jul", "this july", "this jul":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 6, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 7, -1).Format(layout)
	case "august", "aug", "this august", "this aug":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 7, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 8, -1).Format(layout)
	case "september", "sep", "this september", "this sep":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 8, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 9, -1).Format(layout)
	case "october", "oct", "this october", "this oct":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 9, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 10, -1).Format(layout)
	case "november", "nov", "this november", "this nov":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 10, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 11, -1).Format(layout)
	case "december", "dec", "this december", "this dec":
		start = now.New(refTime).BeginningOfYear().AddDate(0, 11, 0).Format(layout)
		end = now.New(refTime).BeginningOfYear().AddDate(0, 12, -1).Format(layout)
	default:
		re := regexp.MustCompile("[0-9]+")
		numArr := re.FindStringSubmatch(input)
		if len(numArr) > 0 {
			num, _ := strconv.Atoi(numArr[0])
			if strings.Contains(input, "day") {
				start = refTime.AddDate(0, 0, -num).Format(layout)
				end = refTime.Format(layout)
			} else if strings.Contains(input, "month") {
				refTime = now.New(refTime.AddDate(0, -num, 0)).BeginningOfMonth()
				start = refTime.Format(layout)
				end = refTime.AddDate(0, num, -1).Format(layout)
			} else if strings.Contains(input, "years") {
				refTime = now.New(refTime.AddDate(-num, 0, 0)).BeginningOfYear()
				start = refTime.Format(layout)
				end = refTime.AddDate(num, 0, -1).Format(layout)
			}
		}
	}
	return
}
