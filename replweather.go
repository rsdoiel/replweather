package replweather

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// NWSRSS URL for the National Weather Service RSS feed
	NWSRSS = "http://www.weather.gov/rss_page.php?site_name=nws"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []*Item  `xml:"item"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	GUID        string    `xml:"guid"`
	Description string    `xml:"description,omitempty"`
	PubDate     time.Time `xml:"pubdate"`
}

type Forcast struct {
	Summary string    `json:"summary"`
	Link    string    `json:"url"`
	Date    time.Time `json:"date"`
}

func (forcast *Forcast) ToString() string {
	return fmt.Sprintf(`{"summary": %q, "link": %q, "date": %q}`, forcast.Summary, forcast.Link, forcast.Date)
}

func toUtf8(iso8859_1_buf []byte) []byte {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return []byte(string(buf))
}

func itemToForcast(item *Item) *Forcast {
	forcast := new(Forcast)
	forcast.Summary = item.Title
	forcast.Link = item.Link
	forcast.Date = item.PubDate
	return forcast
}

func GetNWSRSS() ([]*Forcast, error) {
	var (
		forcasts []*Forcast
		feed     RSS
	)

	res, err := http.Get(NWSRSS)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// Convert to UTF-8
	asText := bytes.Replace(toUtf8(buf), []byte(`encoding="ISO-8859-1"`), []byte(`encoding="UTF-8"`), 1)

	err = xml.Unmarshal(asText, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Channel.Items {
		forcasts = append(forcasts, itemToForcast(item))
	}
	return forcasts, nil
}
