#TV Show calendar

_Tv Show Calendar_ creates an ics calendar file and populates it with the release dates of tv shows which are chosen by the user, via either a command line or web interface. 

  - The app is compatiable with calendar applications including(and tested on Ubuntu calendar and google calendar) 
  - Written in golang


## Installation

```shell
go get github.com/davidnorminton/tvshowCalendar
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