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
	"strings"
	"tvshowCalendar/utils"
)

// SaveDir is the directory relative to the users home directory that will store our data files
// SaveFile is the name of the file that will store our list of saved shows
const (
	SaveDir  = "/.local/share/tvshows/"
	SaveFile = "showlist.txt"
)

// AddTvShow starts the routine to save a TV Show to the saved list
func AddTvShow(show string) (string, error) {

	show = strings.ReplaceAll(show, " ", "-")

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
	fmt.Println("we are here " + show)
	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(home+SaveDir+SaveFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error: Problem opening save file! %v", err)
	}

	defer f.Close()

	err = CheckIfShowInList(show)
	if err != nil {
		return err
	}
	fmt.Println(show)
	if _, err = f.WriteString(show + "\n"); err != nil {
		return fmt.Errorf("There was an error writting to file! %v", err)
	}

	return nil

}

func GetSavelistFileLocation() (string, error) {
	home, err := utils.GetHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting user home directory! %v", err)
	}

	return home + SaveDir + SaveFile, nil
}

// CheckIfShowInLIst checks if the show has already been added to the save list
func CheckIfShowInList(show string) error {

	filename, err := GetSavelistFileLocation()
	if err != nil {
		return err
	}

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

func RemoveShowFromFile(show string) (string, error) {
	filename, err := GetSavelistFileLocation()
	if err != nil {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Problem opening %s", filename)
	}

	scanner := bufio.NewScanner(file)

	newList := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != show {
			newList = append(newList, line)
		}
	}

	defer file.Close()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	appendListToFile(newList)
	return "File updated removed " + show, nil
}

func appendListToFile(list []string) error {
	home, _ := GetSavelistFileLocation()
	f, err := os.OpenFile(home, os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("Error: Problem opening save file! %v", err)
	}

	defer f.Close()

	for _, val := range list {
		if _, err = f.WriteString(val + "\n"); err != nil {
			return fmt.Errorf("There was an error writting to file 2! %v", err)
		}
	}
	return nil
}

// ListShows Lists all the shows in the save file list to the terminal
func ListShows() {

	if err := RunSaveFileChecks(); err != nil {
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
