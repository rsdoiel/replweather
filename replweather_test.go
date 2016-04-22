package replweather

import (
	"fmt"
	"testing"
)

func TestGetNWSRSS(t *testing.T) {
	forcasts, err := GetNWSRSS()
	if err != nil {
		t.Errorf("GetNWSRSS() err %s", err)
	}
	for i, forcast := range forcasts {
		if forcast.Summary == "" {
			t.Errorf("%d: missing Summary -> %s", i, forcast)
		}
		if forcast.Link == "" {
			t.Errorf("%d: missing link -> %s", i, forcast)
		}
		if fmt.Sprintf("%s", forcast.Date) == "" {
			t.Errorf("%d: missing date -> %s", i, forcast)
		}
	}
}
