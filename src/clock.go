package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func fuzzy(now time.Time) string {
	hours := now.Hour()
	minutes := now.Minute()
	glue := "past"
	suffix := ""
	isAM := true

	if hours > 12 {
		isAM = false
		hours = hours - 12
	}

	if minutes > 0 {
		if minutes != 30 && minutes%10 != 0 {
			minutes = (minutes/5 + 1) * 5
		}
		if minutes > 30 {
			minutes = 60 - minutes
			hours++
			glue = "til"
		}
	}

	if minutes == 0 {
		glue = ""
		suffix = " o'clock"
	}
	if hours == 0 {
		suffix = ""
	}
	if hours == 0 && minutes == 0 {
		glue = ""
	}
	if hours == 12 && isAM == false {
		hours = 100
	}

	return trim(fmt.Sprintf(
		"%s %s %s %s",
		getMin(minutes), glue, getHours(hours), suffix,
	))
}

func trim(s string) (r string) {
	space := regexp.MustCompile(` +`)
	r = space.ReplaceAllString(s, " ")
	r = strings.TrimSpace(r)
	return
}
