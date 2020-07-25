// events component of the calendar package, creates and formats the ics events
//

package calendar

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

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
func formatEvent(show IcsEpisode) string {
	ev := EventData{show.Name, show.Date, show.Season, show.Episode}

	eventData := "BEGIN:VEVENT\n" +
		"UID:" + ev.generateUID() + "\n" +
		"DTSTAMP;VALUE=DATE:" + ev.eventDateStampFormat() + "\n" +
		"DTSTART;VALUE=DATE:" + ev.eventDate() + "\n" +
		"DTEND;VALUE=DATE:" + ev.eventDate() + "\n" +
		"SUMMARY:" + ev.eventSummary() + "\n" +
		"END:VEVENT\n"
	return eventData
}
