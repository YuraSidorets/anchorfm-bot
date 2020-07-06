package infra

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration contains all configuration options
type Configuration struct {
	BotToken          string `json:"BotToken"`
	EnableDebug       bool   `json:"EnableDebug"`
	AnchorUser		  string `json:"AnchorUser"`
	AnchorPass		  string `json:"AnchorPass"`
	TimeRangeStart	  int64 `json:"TimeRangeStart"`
	WebStationID	  string `json:"WebStationId"`
	UserID		  	  string `json:"UserId"`
}

// Load configuration from file
func (config Configuration) Load(configPath string) (Configuration, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Println(err)
	}
	decoder := json.NewDecoder(file)
	var configuration Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println(err)
		return configuration, err
	}
	return configuration, nil
}
