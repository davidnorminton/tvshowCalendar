package web

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"tvshowCalendar/showlist"
)

type IndexData struct {
	ShowList []string
}

// IndexHandler handles the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/html/index.html", "web/html/header.html", "web/html/footer.html")
	if err != nil {
		panic(err)
	}
	showListFile, err := showlist.GetSavelistFileLocation()
	showsInList, err := readLines(showListFile)
	data := IndexData{
		ShowList: showsInList,
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
