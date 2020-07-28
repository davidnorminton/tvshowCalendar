// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package calendar includes functions to create an ICS file style event,
// these events are based on a list of TV Shows which are provided by the save list file.
// These events are then written to an ICS file ready to be added to a calendar application
package calendar

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davidnorminton/tvshowCalendar/episodate"
	"github.com/davidnorminton/tvshowCalendar/showlist"
	"github.com/davidnorminton/tvshowCalendar/utils"
)

// IcsFile is the file name of our calendar file
const IcsFile = "tvshows.ics"

// updateCalendar starts the routine to add show air dates to the calendar ics file,
// the shows that are added are based on the shows that are listed in the save show list file
func UpdateCalendar() error {
	err := showlist.RunSaveFileChecks()
	if err != nil {
		return fmt.Errorf("Error in FIle Safety checks")
	}
	if err := getShowsInList(); err != nil {
		return fmt.Errorf("Error writting to file!")
	}
	return nil
}

func getSaveFileLoc() string {
	home, _ := utils.GetHomeDir()
	return home + showlist.SaveDir + showlist.SaveFile
}

type IcsEpisode struct {
	Name    string
	Season  string
	Episode string
	Date    string
}

// getShowsInList retrieves the save show list, then for each show get the latest episodes
// REFACTOR
func getShowsInList() error {

	filename := getSaveFileLoc()
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Problem opening %s", filename)
	}

	defer file.Close()
	var calendarData = []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		html, err := episodate.GetShowData(episodate.GetApiShowUrl(line))
		if err != nil {
			return fmt.Errorf("Error getting %s show Data", line)
		}
		var data episodate.Show
		json.Unmarshal([]byte(html), &data)
		showname := data.TvShow.Name

		for _, val := range data.TvShow.Episodes {

			if CheckAirDate(val.AirDate) != nil {
				continue
			}

			episode := IcsEpisode{
				Name:    showname,
				Season:  strconv.Itoa(val.Season),
				Episode: strconv.Itoa(val.Episode),
				Date:    val.AirDate,
			}

			calendarData = append(calendarData, formatEvent(episode))

		}
	}
	addEventToCalendar(calendarData)

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error writting to file")
	}
	return nil
}

// CheckAirDate Checks if the date provided is prior to current date minus 1 month
// this prevents unnecessary episodes from the past creating a bigger file than we need
func CheckAirDate(date string) error {
	splitDate := strings.Split(date, " ")
	layout := "2006-01-02"
	t, err := time.Parse(layout, splitDate[0])
	if err != nil {
		return err
	}

	if t.Before(time.Now().AddDate(0, -1, 0)) {
		return fmt.Errorf("Date too far in past\n")
	}
	return nil
}

// GetIcsFileLocation retrieves the ics file location
func GetIcsFileLocation() (string, error) {
	home, err := utils.GetHomeDir()
	if err != nil {
		return "", fmt.Errorf("Problem getting home directory")
	}
	return home + showlist.SaveDir + IcsFile, nil
}

// addEventToCalendar is the main loop that write sthe events to the ics file
func addEventToCalendar(events []string) {
	icsfile, err := GetIcsFileLocation()
	if err != nil {
		fmt.Println("Error getting calendar file location")
	}
	file, err := os.OpenFile(icsfile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	_, _ = datawriter.WriteString(startOfIcsFile())
	for _, val := range events {
		_, _ = datawriter.WriteString(val)
	}
	_, _ = datawriter.WriteString("END:VCALENDAR\n")
	datawriter.Flush()
	file.Close()
}

// startOfIcsFile is the first lines of text that are  required by the ICS file
func startOfIcsFile() string {
	return "BEGIN:VCALENDAR\n" +
		"CALSCALE:GREGORIAN\n" +
		"PRODID:-//Ximian//NONSGML Evolution Calendar//EN\n" +
		"VERSION:2.0\n" +
		"X-EVOLUTION-DATA-REVISION:" +
		"2020-06-29T20:39:28.749163Z(1))" +
		"\n"
}
