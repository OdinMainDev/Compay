package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const maxUnixTimeSec = 1<<63 - 1

var (
	ErrorUtilsDateEmptyString        = fmt.Errorf("%s", "EmptyString")
	ErrorUtilsParseError             = func(err error) error { return fmt.Errorf("%s: %s", "ParseError", err.Error()) }
	ErrorUtilsFirstDateParseError    = func(err error) error { return fmt.Errorf("%s: %s", "FirstDateParseError", err.Error()) }
	ErrorUtilsSecondDateParseError   = func(err error) error { return fmt.Errorf("%s: %s", "SecondDateParseError", err.Error()) }
	ErrorUtilsDateMissingSecondValue = fmt.Errorf("%s", "MissingSecondValue")
)

func ParseUnixRange(s string) (start, end int64, err error) {
	if s == "" {
		return 0, 0, ErrorUtilsDateEmptyString
	}

	if strings.Contains(s, "-") {
		parts := strings.SplitN(s, "-", 2)
		if parts[0] == "" {
			start = 0
		} else {
			start, err = strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				return 0, 0, ErrorUtilsFirstDateParseError(err)
			}
		}
		if parts[1] == "" {
			return 0, 0, ErrorUtilsDateMissingSecondValue
		}
		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, 0, ErrorUtilsSecondDateParseError(err)
		}
	} else {
		var v int64
		v, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, 0, ErrorUtilsParseError(err)
		}
		start, end = v, 0
	}

	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > maxUnixTimeSec {
		start = 0
	}
	if end > maxUnixTimeSec {
		end = 0
	}

	return start, end, nil
}
