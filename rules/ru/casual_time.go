package ru

import (
	"regexp"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func CasualTime(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)((это|этим|этот|этим|до|к|после)?\s*(утр(?:ом|а|у)|вечер(?:у|ом|а)|обеда?))`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			lower := strings.ToLower(strings.TrimSpace(m.String()))

			if (c.Hour != nil || c.Minute != nil) && s == rules.OverWrite {
				return false, nil
			}

			switch {
			case strings.Contains(lower, "после обеда"):
				if o.Afternoon != 0 {
					c.Hour = &o.Afternoon
				} else {
					c.Hour = pointer.ToInt(15)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "вечер"):
				if o.Everning != 0 {
					c.Hour = &o.Everning
				} else {
					c.Hour = pointer.ToInt(18)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "утр"):
				if o.Morning != 0 {
					c.Hour = &o.Morning
				} else {
					c.Hour = pointer.ToInt(8)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "обед"):
				if o.Noon != 0 {
					c.Hour = &o.Noon
				} else {
					c.Hour = pointer.ToInt(12)
				}
				c.Minute = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}
