// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package showList implements functions to interact with the curated list
// used to store the TV Show the user wishes to add to their calendar.
package showlist

import (
	"bufio"
	"fmt"
	"os"
	"tvshows/utils"
)

// SaveDir is the directory relative to the users home directory that will store our data files
// SaveFile is the name of the file that will store our list of saved shows
const (
	SaveDir  = "/.local/share/tvshows/"
	SaveFile = "showlist.txt"
)

// AddTvShow starts the routine to save a TV Show to the saved list
func AddTvShow(show string) (string, error) {

	err := RunSaveFileChecks()
	if err != nil {
		return "", err
	}

	// check if show exists in list first
	appenderr := appendShowToFile(show)
	if err == nil {
		return "", appenderr
	}

	return fmt.Sprintf("TV Show %s has been added to list", show), nil

}

// RunSaveFileChecks performs some file system checks:
// ensure we can get the user home directory
// and that the save file exists within the allocated directory
func RunSaveFileChecks() error {
	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	err = checkSaveFileExists(home)
	if err != nil {
		return err
	}
	return nil
}

// check first if the correct directory and save file exists
// if not create them
func checkSaveFileExists(home string) error {

	if _, err := os.Stat(home + SaveDir); os.IsNotExist(err) {
		os.Mkdir(home+SaveDir, os.ModeDir)
		err = os.Chmod(home+SaveDir, 0755)
		if err != nil {
			return fmt.Errorf("Error: Problem creating directory! %v", err)
		}
	}

	file, err := os.OpenFile(home+SaveDir+SaveFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Error: Problem creating save file! %v", err)
	}

	file.Close()
	return nil

}

// Append the show to our save list
func appendShowToFile(show string) error {

	err := RunSaveFileChecks()
	if err != nil {
		return err
	}

	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(home+SaveDir+SaveFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("Error: Problem opening save file! %v", err)
	}

	defer f.Close()

	err = checkIfShowInList(show)
	if err != nil {
		return err
	}

	if _, err = f.WriteString(show + "\n"); err != nil {
		return fmt.Errorf("There was an error writting to file! %v", err)
	}

	return nil

}

// check if the show has already been added to the save list
func checkIfShowInList(show string) error {

	home, err := utils.GetHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting user home directory! %v", err)
	}

	filename := home + SaveDir + SaveFile
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Problem opening %s", filename)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == show {
			return fmt.Errorf("Show already in list!")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil

}

// ListShows Lists all the shows in the save file list to the terminal
func ListShows() {

	err := RunSaveFileChecks()
	if err != nil {
		fmt.Println(err)
	}

	home, err := utils.GetHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	filename := home + SaveDir + SaveFile
	file, _ := os.Open(filename)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
