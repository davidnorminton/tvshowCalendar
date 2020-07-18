// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package utils is used for general utilities that are used throughtout the programs packages
package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// GetHomeDir retrieves the current users home directoy
func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error getting user home directory! %v", err)
	}
	return home, nil
}

// FmtDate formats a date string in the format 2020-10-06 to October 6, 2020
func FmtDate(date string) string {
	layoutISO := "2006-01-02"
	layout := "Jan 2, 2006"
	split := strings.Split(date, " ")

	t, _ := time.Parse(layoutISO, split[0])

	return t.Format(layout)
}
