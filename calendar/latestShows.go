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

func GetLatestEpisodes() ([]map[string]string, error) {
	cal, err := GetIcsFileLocation()
	if err != nil {
		return nil, fmt.Errorf("Error getting ICS file!")
	}

	list, err := sliceEpisodes(cal)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func sliceEpisodes(cal string) ([]map[string]string, error) {
	file, err := os.Open(cal)
	if err != nil {
		return nil, fmt.Errorf("Error opening ics file")
	}

	defer file.Close()

	list := []map[string]string{}
	temp := map[string]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
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
		return nil, fmt.Errorf("Error scanning file!")
	}
	return list, nil
}

func splitLine(line string) string {
	split := strings.Split(line, ":")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}
