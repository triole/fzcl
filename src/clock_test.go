package main

import (
	"fmt"
	"testing"
	"time"
)

type params struct {
	hours   int
	minutes int
	out     string
}

func (p params) Name() string {
	return fmt.Sprintf("%d:%02d", p.hours, p.minutes)
}

func TestTime(t *testing.T) {
	var times = []params{
		{0, 0, "midnight"},
		{0, 30, "half past midnight"},
		{2, 42, "quarter til three"},
		{3, 22, "twenty five past three"},
		{3, 25, "half past three"},
		{3, 30, "half past three"},
		{7, 00, "seven o'clock"},
		{12, 0, "twelve o'clock"},
		{12, 2, "five past twelve"},
		{13, 54, "five til two"},
		{13, 56, "two o'clock"},
		{15, 37, "twenty til four"},
		{18, 00, "six o'clock"},
		{23, 44, "quarter til midnight"},
		{23, 48, "ten til midnight"},
	}

	for _, param := range times {
		t.Run(param.Name(), func(t *testing.T) {
			now, _ := time.Parse("15:04", param.Name())
			result := fuzzy(now)
			if result != param.out {
				t.Errorf("%s got %s, want %s", now.Format("15:04"), result, param.out)
			}
		})
	}
}
