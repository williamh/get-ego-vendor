// Copyright 2019 William Hubbs
// Copyright 2019 Sony Interactive Entertainment
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

func convertVersion(v string) string {
	pseudoVersionForms := []string {
		`^v[0-9]+\.0\.0-[0-9]+-[[:xdigit:]]+$`,
		`^v[0-9]+\.[0-9]+\.[0-9]+-pre\.0\.[0-9]+-[[:xdigit:]]+$`,
		`^v[0-9]+\.[0-9]+\.[0-9]+-0.[0-9]+-[[:xdigit:]]+$`,
	}
	for _, pv := range pseudoVersionForms {
		match, _ := regexp.MatchString(pv, v)
		if match {
			pos := strings.LastIndex(v, "-")
			if pos + 1 < len(v) {
				return v[pos+1:]
			} else {
				return ""
			}
		}
	}
	return v
}
