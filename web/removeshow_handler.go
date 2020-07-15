package web

import (
	"fmt"
	"net/http"
	"tvshowCalendar/showlist"
)

func RemoveShowHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("remove")
	if _, err := showlist.RemoveShowFromFile(query); err != nil {
		fmt.Fprintf(w, "error")
	} else {
		fmt.Fprintf(w, "removed")
	}
}
