# TV Show calendar

__Tv Show Calendar__  creates an ics calendar file and add events of upcoming tv show release dates which are chosen prior by the user. The shows to track can be added either via the command line or a web interface. 
In order to get the tv show data the app uses the [episodate REST api](https://www.episodate.com/api) to retrieve the tv show data. The app is compatiable with calendar applications including(and tested) on Ubuntu gnome calendar and google calendar.

## Installation

```shell
go get github.com/davidnorminton/tvshowCalendar
cd tvshowCalendar
go build tvshowCalendar
```

If you wish to clone the repository from github, first create the directory structure $GOPATH/src/github.com/davidnorminton then run:
```shell
git clone https://github.com/davidnorminton/tvshowCalendar
```

## Terminal usage

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


### Home page
![Home Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/home-page.png "Screenshot of the home page")

### Search
![Search Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/search-page.png "Screenshot of a typical search page")

### Show Details
![Show Details Page](https://raw.githubusercontent.com/davidnorminton/tvshowCalendar/master/img/details-page.png "Screenshot of a typical show details page")


## How to use the ics file


By default the on ubuntu the ics file is save at /home/USER/.local/share/tvshows.

### gnome calendar

  - Open gnome calendar
  - Click on manage calendars
  - Manage Calendars
  - add a calendar
  - Import Calendar > then select the ics file

### google calendar

  - Open Google Calendar.
  - In the top right, click Settings(gear icon) > Settings.
  - Click Import & Export.
  - Click Select the ics file from your computer and select the file you exported.
  - Choose which calendar to add the imported events to. By default, events will be imported into your primary calendar.
  - Click Import.

## Todo

  - ability to save a files other than the default
  - refactor the code to make it simple to change the tv show REST api
