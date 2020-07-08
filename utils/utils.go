// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package utils is used for general utilities that are used throughtout the programs packages
package utils

import (
	"fmt"
	"os"
)

// GetHomeDir retrieves the current users home directoy
func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error getting user home directory! %v", err)
	}
	return home, nil
}
