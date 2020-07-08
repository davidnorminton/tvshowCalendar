package episodate

import (
	"testing"
	"tvshows/episodate"
)

func TestGetApiShowUrl(t *testing.T) {
	url := episodate.GetApiShowUrl("Doom Patrol")
	if url != "https://www.episodate.com/api/show-details?q=Doom-Patrol" {
		t.Errorf("url is incorrect!")
	}
}
