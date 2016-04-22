package replweather

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Forcast struct {
	Summary string    `json:"summary"`
	Link    string    `json:"url"`
	Date    time.Time `json:"date"`
}

func (forcast *Forcast) ToString() string {
	return fmt.Sprintf(`{"summary": %q, "url": %q, "date": %q}`, forcast.Summary, forcast.Link, forcast.Date)
}

func rssToUtf8(AsIso8859_1 []byte) []byte {
	buf := make([]rune, len(AsIso8859_1))
	for i, b := range AsIso8859_1 {
		buf[i] = rune(b)
	}
	return []byte(strings.Replace(string(buf), `encoding="ISO-8859-1"`, `encoding="UTF-8"`, 1))
}

func nwsItemToForcast(item *Item) *Forcast {
	forcast := new(Forcast)
	forcast.Summary = item.Title
	forcast.Link = item.Link
	if d, err := time.Parse(`Mon, 2 Jan 2006 15:04:05 -0700`, item.PubDate); err == nil {
		forcast.Date = d
	}
	return forcast
}

func GetNWSRSS() ([]*Forcast, error) {
	var (
		feed     RSS
		forcasts []*Forcast
	)
	res, err := http.Get(NWSRSS)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("DEBUG rss raw: %s\n", buf)

	// Convert to UTF-8 and unmarshal
	err = xml.Unmarshal(rssToUtf8(buf), &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Channel.Items {
		forcasts = append(forcasts, nwsItemToForcast(item))
	}
	// if src, err := json.Marshal(forcasts); err == nil {
	// 	fmt.Printf("DEBUG %s\n", src)
	// }
	return forcasts, nil
}
