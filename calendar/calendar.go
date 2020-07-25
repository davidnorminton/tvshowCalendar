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
	"math/rand"
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
	var calendarData = [][7]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		html, err := episodate.GetShowData(episodate.GetApiShowUrl(line))
		if err != nil {
			return fmt.Errorf("Error getting %s show Data", line)
		}
		var data episodate.Show
		json.Unmarshal([]byte(html), &data)

		for _, val := range data.TvShow.Episodes {

			if CheckAirDate(val.AirDate) != nil {
				continue
			}

			episode := IcsEpisode{
				Name:    val.Name,
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

type EventData struct {
	Name, Date, Season, Episode string
}

// generateUID Generates a random UID
func (e EventData) generateUID() string { return strconv.Itoa(rand.Int()) }

// eventDate formats the date into the form of Ymd with no characters
func (e EventData) eventDate() string {
	startDate := strings.Split(e.Date, " ")
	return strings.Replace(startDate[0], "-", "", 2)
}

// eventDateStempFormat formats the date from the REST api to one which is suitable for the ics file
func (e EventData) eventDateStampFormat() string {
	replace := strings.NewReplacer(" ", "T", "-", "", ":", "")
	return replace.Replace(e.Date)
}

// eventSummary formats the summary to include shows season, episode and name
func (e EventData) eventSummary() string {
	return fmt.Sprintf("%s, S%s E%s", e.Name, e.Season, e.Episode)
}

// formatEvent formats the data into a readable ics event format
func formatEvent(show IcsEpisode) (eventData [7]string) {
	ev := EventData{
		show.Name,
		show.Date,
		show.Season,
		show.Episode,
	}

	eventData = [7]string{}

	//summary := eventSummary(name, season, episode)
	eventData[0] = "BEGIN:VEVENT"
	eventData[1] = "UID:" + ev.generateUID()
	eventData[2] = "DTSTAMP;VALUE=DATE:" + ev.eventDateStampFormat()
	eventData[3] = "DTSTART;VALUE=DATE:" + ev.eventDate()
	eventData[4] = "DTEND;VALUE=DATE:" + ev.eventDate()
	eventData[5] = "SUMMARY:" + ev.eventSummary()
	eventData[6] = "END:VEVENT"
	return
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
func addEventToCalendar(events [][7]string) {
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
	for i := 0; i < len(events); i++ {
		for j := 0; j < len(events[i]); j++ {
			_, _ = datawriter.WriteString(events[i][j] + "\n")
		}
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
