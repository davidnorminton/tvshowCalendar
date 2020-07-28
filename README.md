# TV Show calendar

__Tv Show Calendar__ creates an ics calendar file and populates it with the release dates of tv shows which are chosen by the user, via either a command line or web interface. 

  - The app is compatiable with calendar applications including(and tested on Ubuntu gnome calendar and google calendar) 
  - Written in golang


## Installation

```shell
go get github.com/davidnorminton/tvshowCalendar
cd tvshowCalendar
go build main.go
```

##Terminal usage

Tv show calendar uses the following flags:

#### Start the web interface
```shell
./tvshowCalendar -S
```

#### Search for a tv hsow
```shell
./tvshowCalendar -s="A cool show"
```

#### Add a show to the save list
```shell
./tvshowCalendar -a="A cool show"
```

#### List all shows in our save list
```shell
./tvshowCalendar -l
```

#### Output a tv shows details
```shell
./tvshowCalendar -d="a cool show"
```

#### Remove a show from the save list
```shell
./tvshowCalendar -r="was a cool show"
```

#### Update the ics file
```shell
./tvshowCalendar -u
```

#### Output the latest release dates
```shell
./tvshowCalendar -e
```

## Web interface usage

## How to use the ics file

### Home page
![Home Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/home-page.png "Screenshot of the home page")

### Search
![Search Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/search-page.png "Screenshot of a typical search page")

### Show Details
![Show Details Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/details-page.png "Screenshot of a typical show details page")

By default the on ubuntu the ics file is save at /home/USER/.local/share/tvshows.

### gnome calendar

### google calendar

  - Open Google Calendar.
  - In the top right, click Settings(gear icon) > Settings.
  - Click Import & Export.
  - Click Select the ics file from your computer and select the file you exported.
  - Choose which calendar to add the imported events to. By default, events will be imported into your primary calendar.
  - Click Import.