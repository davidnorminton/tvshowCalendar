package calendar

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type episode struct {
	Summary string
	Date    string
}

func GetLatestEpisodes() {
	cal, err := GetIcsFileLocation()
	if err != nil {
		fmt.Println("Error getting ICS file!")
	}

	//createSortedList()
	sliceEpisodes(cal)
}

func sliceEpisodes(cal string) {
	file, err := os.Open(cal)
	if err != nil {
		fmt.Println("Error opening ics file")
	}

	defer file.Close()

	list := []map[string]string{}
	temp := map[string]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(list)
		line := scanner.Text()
		split := splitLine(line)
		switch {
		case strings.Contains(line, "BEGIN:VEVENT"):
			temp = map[string]string{}
		case strings.Contains(line, "DTSTART;VALUE=DATE"):
			temp["Date"] = split
		case strings.Contains(line, "SUMMARY"):
			temp["Summary"] = split
		case strings.Contains(line, "END:VEVENT"):
			list = append(list, temp)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sort(list)

}

func splitLine(line string) string {
	split := strings.Split(line, ":")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

func sort(list []map[string]string) {
	fmt.Println(list)
}
