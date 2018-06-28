package daysago

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

const Layout = "2006-01-02"

//GetDateStringFromWordString takes as input a sentence or word and returns a time for that
func GetDateStringFromWordString(refTime time.Time, input string) (out string) {
	input = strings.ToLower(input)
	switch input {
	case "yesterday":
		out = refTime.AddDate(0, 0, -1).Format(Layout)
	case "today":
		out = refTime.Format(Layout)
	case "tomorrow":
		out = refTime.AddDate(0, 0, 1).Format(Layout)
	case "day after tomorrow":
		out = refTime.AddDate(0, 0, 2).Format(Layout)
	case "day before yesterday":
		out = refTime.AddDate(0, 0, -2).Format(Layout)
	case "last month", "past month":
		out = now.New(refTime.AddDate(0, -1, 0)).BeginningOfMonth().Format(Layout)
	case "last year", "past year":
		out = now.New(refTime.AddDate(-1, 0, 0)).BeginningOfYear().Format(Layout)
	case "past week", "last week":
		out = now.New(refTime.AddDate(0, 0, -7)).BeginningOfWeek().Format(Layout)
	default:
		re := regexp.MustCompile("[0-9]+")
		numArr := re.FindStringSubmatch(input)
		if len(numArr) > 0 {
			num, _ := strconv.Atoi(numArr[0])
			if strings.Contains(input, "day") {
				out = refTime.AddDate(0, 0, -num).Format(Layout)
			} else if strings.Contains(input, "month") {
				out = refTime.AddDate(0, -num, 0).Format(Layout)
			} else if strings.Contains(input, "years") {
				out = refTime.AddDate(-num, 0, 0).Format(Layout)
			}
		}
	}
	return
}

//GetDateStringRangeFromWordString takes as input a sentence or word and returns a range of time as start and end
func GetDateStringRangeFromWordString(refTime time.Time, input string) (start, end string) {
	var err error
	var re = regexp.MustCompile(` +`)
	input = re.ReplaceAllString(strings.ToLower(strings.TrimSpace(input)), " ")
	refTime = refTime.UTC()

	_, err = now.Parse(input)
	if err != nil {
		switch input {
		case "yesterday":
			end = refTime.Format(Layout)
			start = refTime.AddDate(0, 0, -1).Format(Layout)
		case "today":
			end = refTime.Format(Layout)
			start = end
		case "tomorrow":
			start = refTime.Format(Layout)
			end = refTime.AddDate(0, 0, 1).Format(Layout)
		case "day after tomorrow":
			start = refTime.Format(Layout)
			end = refTime.AddDate(0, 0, 2).Format(Layout)
		case "day before yesterday":
			start = refTime.AddDate(0, 0, -2).Format(Layout)
			end = refTime.Format(Layout)
		case "last month", "past month":
			refTime = now.New(refTime.AddDate(0, -1, 0)).BeginningOfMonth()
			start = refTime.Format(Layout)
			end = now.New(refTime).EndOfMonth().Format(Layout)
		case "last year", "past year":
			refTime = now.New(refTime.AddDate(-1, 0, 0)).BeginningOfYear()
			start = refTime.Format(Layout)
			end = refTime.AddDate(1, 0, -1).Format(Layout)
		case "past week", "last week":
			refTime = now.New(refTime.AddDate(0, 0, -7)).BeginningOfWeek()
			start = refTime.Format(Layout)
			end = now.New(refTime).EndOfWeek().Format(Layout)
		case "this week":
			start = now.New(refTime).BeginningOfWeek().Format(Layout)
			end = now.New(refTime).EndOfWeek().Format(Layout)
		case "this month":
			start = now.New(refTime).BeginningOfMonth().Format(Layout)
			end = now.New(refTime).EndOfMonth().Format(Layout)
		case "this year":
			start = now.New(refTime).BeginningOfYear().Format(Layout)
			end = now.New(refTime).EndOfYear().Format(Layout)
		case "january", "jan", "this january", "this jan", "in jan", "in january", "for january", "for jan":
			start = now.New(refTime).BeginningOfYear().Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 1, -1).Format(Layout)
		case "february", "feb", "this february", "this feb", "in feb", "in february", "for feb", "for february":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 1, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 2, -1).Format(Layout)
		case "march", "mar", "this march", "this mar", "in mar", "in march", "for mar", "for march":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 2, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 3, -1).Format(Layout)
		case "april", "apr", "this april", "this apr", "in apr", "in april", "for apr", "for april":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 3, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 4, -1).Format(Layout)
		case "may", "this may", "in may", "for may":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 4, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 5, -1).Format(Layout)
		case "june", "jun", "this june", "this jun", "in jun", "in june", "for jun", "for june":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 5, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 6, -1).Format(Layout)
		case "july", "jul", "this july", "this jul", "in jul", "in july", "for jul", "for july":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 6, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 7, -1).Format(Layout)
		case "august", "aug", "this august", "this aug", "in aug", "in august", "for aug", "for august":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 7, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 8, -1).Format(Layout)
		case "september", "sep", "this september", "this sep", "in sep", "in september", "for sep", "for september", "this sept", "sept", "in sept", "for sept":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 8, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 9, -1).Format(Layout)
		case "october", "oct", "this october", "this oct", "in oct", "in october", "for oct", "for october":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 9, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 10, -1).Format(Layout)
		case "november", "nov", "this november", "this nov", "in nov", "in november", "for nov", "for november":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 10, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 11, -1).Format(Layout)
		case "december", "dec", "this december", "this dec", "in dec", "in december", "for dec", "for december":
			start = now.New(refTime).BeginningOfYear().AddDate(0, 11, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(0, 12, -1).Format(Layout)
		case "past january", "past jan", "last january", "last jan", "prev january", "prev jan":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 0, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 1, -1).Format(Layout)
		case "past february", "past feb", "last february", "last feb", "prev february", "prev feb":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 1, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 2, -1).Format(Layout)
		case "last march", "last mar", "prev march", "prev mar", "past march", "past mar":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 2, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 3, -1).Format(Layout)
		case "last april", "last apr", "prev april", "prev apr", "past april", "past apr":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 3, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 4, -1).Format(Layout)
		case "last may", "prev may":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 4, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 5, -1).Format(Layout)
		case "last june", "last jun", "prev june", "prev jun", "past june", "past jun":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 5, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 6, -1).Format(Layout)
		case "last july", "last jul", "prev july", "prev jul", "past july", "past jul":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 6, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 7, -1).Format(Layout)
		case "last august", "last aug", "prev august", "prev aug", "past august", "past aug":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 7, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 8, -1).Format(Layout)
		case "last september", "last sep", "prev september", "prev sep", "past september", "past sep":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 8, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 9, -1).Format(Layout)
		case "last october", "last oct", "prev october", "prev oct", "past october", "past oct":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 9, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 10, -1).Format(Layout)
		case "last november", "last nov", "prev november", "prev nov", "past november", "past nov":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 10, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 11, -1).Format(Layout)
		case "last december", "last dec", "prev december", "prev dec", "past december", "past dec":
			start = now.New(refTime).BeginningOfYear().AddDate(-1, 11, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(-1, 12, -1).Format(Layout)
		case "next january", "next jan", "coming january", "coming jan", "upcoming january", "upcoming jan":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 0, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 1, -1).Format(Layout)
		case "next february", "next feb", "coming february", "coming feb", "upcoming february", "upcoming feb":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 1, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 2, -1).Format(Layout)
		case "coming march", "coming mar", "upcoming march", "upcoming mar", "next march", "next mar":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 2, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 3, -1).Format(Layout)
		case "coming april", "coming apr", "upcoming april", "upcoming apr", "next april", "next apr":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 3, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 4, -1).Format(Layout)
		case "coming may", "upcoming may", "next may":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 4, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 5, -1).Format(Layout)
		case "coming june", "coming jun", "upcoming june", "upcoming jun", "next june", "next jun":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 5, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 6, -1).Format(Layout)
		case "coming july", "coming jul", "upcoming july", "upcoming jul", "next july", "next jul":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 6, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 7, -1).Format(Layout)
		case "coming august", "coming aug", "upcoming august", "upcoming aug", "next august", "next aug":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 7, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 8, -1).Format(Layout)
		case "coming september", "coming sep", "upcoming september", "upcoming sep", "next september", "next sep":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 8, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 9, -1).Format(Layout)
		case "coming october", "coming oct", "upcoming october", "upcoming oct", "next october", "next oct":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 9, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 10, -1).Format(Layout)
		case "coming november", "coming nov", "upcoming november", "upcoming nov", "next november", "next nov":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 10, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 11, -1).Format(Layout)
		case "coming december", "coming dec", "upcoming december", "upcoming dec", "next december", "next dec":
			start = now.New(refTime).BeginningOfYear().AddDate(1, 11, 0).Format(Layout)
			end = now.New(refTime).BeginningOfYear().AddDate(1, 12, -1).Format(Layout)

		default:
			re := regexp.MustCompile("[0-9]+")
			numArr := re.FindStringSubmatch(input)
			if len(numArr) > 0 {
				num, _ := strconv.Atoi(numArr[0])
				if strings.Contains(input, "day") {
					start = refTime.AddDate(0, 0, -num).Format(Layout)
					end = refTime.Format(Layout)
				} else if strings.Contains(input, "month") {
					refTime = now.New(refTime.AddDate(0, -num, 0)).BeginningOfMonth()
					start = refTime.Format(Layout)
					end = refTime.AddDate(0, num, -1).Format(Layout)
				} else if strings.Contains(input, "years") {
					refTime = now.New(refTime.AddDate(-num, 0, 0)).BeginningOfYear()
					start = refTime.Format(Layout)
					end = refTime.AddDate(num, 0, -1).Format(Layout)
				}
			}
		}
	} else {
		startTime, _ := now.New(refTime).Parse(input)
		start = startTime.Format(Layout)
		end = now.New(startTime).EndOfYear().Format(Layout)
	}
	return
}
