package anchor

import (
	"anchorfm-bot/infra"
	"log"
	"time"
	"fmt"
	"encoding/json"

	"github.com/go-rod/rod"	
)

type Totals struct {
	AudienceSize     float32 `json:"averageCt"`
	TotalPlays       int     `json:"ct"`
}

type AnchorResponse struct {
	Data struct {
		Rows	[][]interface{} `json:"rows"`
	} `json:"data"`
}

// LoginAnchor returns your anchor dashboard page using creds provided
func LoginAnchor(browser *rod.Browser, config infra.Configuration) *rod.Page {
	page := browser.Page("https://anchor.fm/login")
	page.WaitLoad()

	log.Printf("input user email")
	page.Element("#email").Click().Input(config.AnchorUser)
	
	log.Printf("input user pass")
	page.Element("#password").Click().Input(config.AnchorPass)
	
	log.Printf("submit login form")	
	page.Element("button.ButtonBase__root___2GNnu").Click()
	page.WaitLoad()
	page.Element("div.css-vax5dl").Text()
	return page
}

// GetAge returns plays percent by age
func GetAge(page *rod.Page, config infra.Configuration) map[string]float64 {
	url := fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/playsByAgeRange?timeRangeStart=%v&timeRangeEnd=%v&", config.WebStationID, config.TimeRangeStart, time.Now().Unix())	
	result := getData(page, url)
	total := make(map[string]float64)
	for _, play := range result.Data.Rows {
		name := play[0].(string)
		plays := play[1].(float64)
		total[name] += plays
	}	
	return total
}

// GetGender returns plays percent by gender
func GetGender(page *rod.Page, config infra.Configuration) map[string]float64 {
	url := fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/playsByGender?timeRangeStart=%v&timeRangeEnd=%v&", config.WebStationID, config.TimeRangeStart, time.Now().Unix())	
	result := getData(page, url)
	total := make(map[string]float64)
	for _, play := range result.Data.Rows {
		name := play[0].(string)
		plays := play[1].(float64)
		total[name] += plays
	}	
	return total
}

// GetPlatform returns percent of plays by platform
func GetPlatform(page *rod.Page, config infra.Configuration) map[string]float64 {
	url := fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/playsByApp?userId=%v&timeRangeStart=%v&timeRangeEnd=%v&", config.WebStationID, config.UserID, config.TimeRangeStart, time.Now().Unix())	
	result := getData(page, url) 
	total := make(map[string]float64)
	for _, play := range result.Data.Rows {
		name := play[0].(string)
		plays := play[1].(float64)
		total[name] += plays
	}	
	return total
}

// GetGeo returns geolocation based on plays 
func GetGeo(page *rod.Page, config infra.Configuration) map[string]float64 {
	url := fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/playsByGeo?limit=200&resultGeo=geo2", config.WebStationID)
	result := getData(page, url) 
	total := make(map[string]float64)
	for _, play := range result.Data.Rows {
		name := play[0].(string)
		plays := play[1].(float64)
		total[name] += plays
	}	
	return total
}

// GetTopEpisodes returns total plays for each episode by Date
func GetTopEpisodes(page *rod.Page, config infra.Configuration) map[string]int64 {
	url :=  fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/topEpisodes?limit=10&timeRangeStart=%v&timeRangeEnd=%v", config.WebStationID, config.TimeRangeStart, time.Now().Unix())	
	result := getData(page, url) 
	total := make(map[string]int64)
	for _, play := range result.Data.Rows {
		name := play[0].(string)
		plays := int64(play[1].(float64))
		total[name] += plays
	}	
	return total
}

// GetPlaysByEpisode returns total plays for each episode
func GetPlaysByEpisode(page *rod.Page, config infra.Configuration) map[string]int64 {
	url := fmt.Sprintf("https://anchor.fm/api/proxy/v3/analytics/station/webStationId:%v/playsByEpisode?timeInterval=604800&limit=3", config.WebStationID)	
	result := getData(page, url) 
	total := make(map[string]int64)
	for _, play := range result.Data.Rows {
		timestamp := int64(play[0].(float64))
		t := time.Unix(timestamp, 0)
		dateStr := t.Format(`01/02/06`)

		plays := int64(play[2].(float64))
		total[dateStr] += plays
	}	
	return total
}

// GetTotalsCount returns total plays for podcast
func GetTotalsCount(page *rod.Page, config infra.Configuration) Totals {
	url := "https://anchor.fm/api/podcast/analytics/downloads/all"	
	page.Navigate(url)
	page.WaitLoad()
	
	totalsString := page.ElementByJS(`() => document.body`).Text()

	var data Totals 
	err := json.Unmarshal([]byte(totalsString), &data)
	if err != nil {
		log.Println(err)
	}
	return data
 }

 func getData(page *rod.Page, url string) AnchorResponse {
	page.Navigate(url)
	page.WaitLoad()
	
	playsString := page.ElementByJS(`() => document.body`).Text()

	var data AnchorResponse 
	err := json.Unmarshal([]byte(playsString), &data)
	if err != nil {
		log.Println(err)
	}
	return data
 }

// StartBrowser starts headless browser
func StartBrowser(config infra.Configuration) *rod.Browser {
	browser := rod.New().Connect()
	log.Printf("Browser started")	
	return browser
}

// StopBrowser stops headless browser
func StopBrowser(browser *rod.Browser){
	log.Println("Before stop browser")
	browser.Close()
	log.Printf("Browser stopped")	
}