package convert

import (
	"fmt"
	"strconv"
)

func TimeConv(timeFormat string) string {

	var hourString string = timeFormat[:2]

	var minuteString string = timeFormat[3:5]

	var secondString string = timeFormat[6:8]

	switch {
	case string(hourString[0]) == "0":
		hourString = hourString[1:]
	case string(minuteString[0]) == "0":
		minuteString = minuteString[1:]
	case string(secondString[0]) == "0":
		secondString = secondString[1:]
	}

	var hour, _ = strconv.ParseFloat(hourString, 32)
	var minute, _ = strconv.ParseFloat(minuteString, 32)
	var second, _ = strconv.ParseFloat(secondString, 32)

	hour, minute, second = hour/24*10, minute/60*100, second/60*100

	var output = fmt.Sprintf("%.2f:%.2f:%.2f", hour, minute, second)

	return output
}
