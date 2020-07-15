package web

import (
	"fmt"
	"net/http"
	"tvshowCalendar/showlist"
)

func AddShowHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("add")
	_, err := showlist.AddTvShow(query)
	if err != nil {
		fmt.Fprintf(w, "error")
	} else {
		fmt.Fprintf(w, "added")
	}
}
