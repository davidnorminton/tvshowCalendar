package web

import (
	"fmt"
	"net/http"
	"tvshowCalendar/calendar"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if err := calendar.UpdateCalendar(); err != nil {
		fmt.Fprintf(w, "error")
	} else {
		fmt.Fprintf(w, "updated")
	}
}
