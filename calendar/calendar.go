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

	"tvshows/episodate"
	"tvshows/showlist"
	"tvshows/utils"
)

// IcsFile is the file name of our calendar file
const (
	IcsFile = "tvshows.ics"
)

// updateCalendar starts the routine to add show air dates to the calendar ics file,
// the shows that are added are based on the shows that are listed in the save show list file
func UpdateCalendar() {
	fmt.Println("Updating Calendar")
	err := showlist.RunSaveFileChecks()
	if err != nil {
		fmt.Println(err)
	}
	getShowsInList()
}

// getShowsInList retrieves the save show list, then for each show get the latest episodes
// REFACTOR
func getShowsInList() {
	home, err := utils.GetHomeDir()
	if err != nil {
		fmt.Println("Error: Problem opening file!")
	}

	filename := home + showlist.SaveDir + showlist.SaveFile
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Problem opening %s", filename)
	}

	defer file.Close()
	var calendarData = [][7]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		html, err := episodate.GetShowData(episodate.GetApiShowUrl(line))

		var data episodate.Show
		json.Unmarshal([]byte(html), &data)

		if err != nil {
			fmt.Printf("Error getting %s show Data", line)
		}

		for i := 0; i < len(data.TvShow.Episodes); i++ {

			// check if Air date id prior to current date
			if checkAirDate(data.TvShow.Episodes[i].AirDate) != nil {
				continue
			}

			episode := map[string]interface{}{
				"name":    data.TvShow.Name,
				"season":  strconv.Itoa(data.TvShow.Episodes[i].Season),
				"episode": strconv.Itoa(data.TvShow.Episodes[i].Episode),
				"date":    data.TvShow.Episodes[i].AirDate,
			}

			calendarData = append(calendarData, formatEvent(episode))

		}
	}
	addEventToCalendar(home, calendarData)

	// write calendarData to file ics

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// checkAirDate Checks if the date provided is prior to current date minus 1 month
// this prevents unnecessary episodes from the past creating a bigger file than we need
func checkAirDate(date string) error {
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

// formatEvent formats the data into a readable ics event format
func formatEvent(show map[string]interface{}) (eventData [7]string) {
	name := show["name"].(string)
	season := show["season"].(string)
	episode := show["episode"].(string)
	summary := eventSummary(name, season, episode)
	eventData[0] = "BEGIN:VEVENT"
	eventData[1] = "UID:" + generateUID()
	eventData[2] = "DTSTAMP;VALUE=DATE:" + eventDateStampFormat(show["date"].(string))
	eventData[3] = "DTSTART;VALUE=DATE:" + eventDate(show["date"].(string))
	eventData[4] = "DTEND;VALUE=DATE:" + eventDate(show["date"].(string))
	eventData[5] = "SUMMARY:" + summary
	eventData[6] = "END:VEVENT"
	return
}

// generateUID Generates a random UID
func generateUID() string {
	return strconv.Itoa(rand.Int())
}

// eventDate formats the date into the form of Ymd with no characters
func eventDate(date string) string {
	startDate := strings.Split(date, " ")
	return strings.Replace(startDate[0], "-", "", 2)
}

// eventDateStempFormat formats the date from the api to one which is suitable for the ics file
func eventDateStampFormat(date string) string {
	replace := strings.NewReplacer(" ", "T", "-", "", ":", "")
	return replace.Replace(date)
}

// eventSummary formats the summary to include shows season, episode and name
func eventSummary(name string, season string, episode string) string {
	return fmt.Sprintf("%s, season %s episode %s released", name, season, episode)
}

// addEventToCalendar is the main loop that write sthe events to the ics file
func addEventToCalendar(home string, events [][7]string) {
	icsfile := home + showlist.SaveDir + IcsFile
	file, err := os.OpenFile(icsfile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	_, _ = datawriter.WriteString(startOfIcsFile())
	for i := 0; i < len(events); i++ {
		for j := 0; j < len(events[i]); j++ {
			// write to file
			//fmt.Println(events[i][j])
			_, _ = datawriter.WriteString(events[i][j] + "\n")
		}
	}
	_, _ = datawriter.WriteString("END:VCALENDAR\n")
	datawriter.Flush()
	file.Close()
}

// startOfIcsFile is the first lines of text that are  required by the ICS file
func startOfIcsFile() string {
	return "BEGIN:VCALENDAR\n" + "CALSCALE:GREGORIAN\n" + "PRODID:-//Ximian//NONSGML Evolution Calendar//EN\n" +
		"VERSION:2.0\n" + "X-EVOLUTION-DATA-REVISION:2020-06-29T20:39:28.749163Z(1))\n"
}
