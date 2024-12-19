package utils

import (
	"fmt"
	"regexp"
)

var Providers = map[string]string{
	"99895": "uzmobile",
	"99899": "uzmobile",
	"99877": "uzmobile",
	"99890": "beeline",
	"99891": "beeline",
	"99893": "ucell",
	"99894": "ucell",
	"99855": "ucell",
	"99850": "ucell",
	"99833": "humans",
	"99897": "mobiuz",
	"99888": "mobiuz",
	"99898": "perfectum",
	"99820": "oq",
}

func PredictProvider(phoneNumber string) string {
	for prefix, provider := range Providers {
		pattern := fmt.Sprintf("^%s[0-9]\\d{6}$", prefix)
		matched, _ := regexp.MatchString(pattern, phoneNumber)
		if matched {
			return provider
		}
	}
	return ""
}