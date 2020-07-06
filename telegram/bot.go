package telegram

import (
	"anchorfm-bot/infra"
	"anchorfm-bot/anchor"
	"log"
	"fmt"
	"sort"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Run starts telegram bot listener
//
//it requires app configuration provided
func Run(config infra.Configuration) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.EnableDebug

	log.Printf("Bot started on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}
	browser := anchor.StartBrowser(config)
	defer anchor.StopBrowser(browser)
	page := anchor.LoginAnchor(browser, config)
	
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "stats":
			totals := anchor.GetTotalsCount(page, config)
			sendMessage(bot, update.Message.Chat.ID, fmt.Sprintf("Total Plays: %d \nEstimated audience size: %v", totals.TotalPlays, totals.AudienceSize))
		case "plays":
			plays := anchor.GetPlaysByEpisode(page, config)

    		keys := []string{}
    		for key := range plays {
        		keys = append(keys, key)
    		}
    		sort.Strings(keys)
			var playsMessage string 
    		for i := range keys {
        		key := keys[i]
        		value := plays[key]
        		playsMessage += fmt.Sprintf("Day: %s, Plays: %d \n", key, value)
    		}
			sendMessage(bot, update.Message.Chat.ID, playsMessage)
		case "topepisodes":
			episodes := anchor.GetTopEpisodes(page, config)

    		keys := []string{}
    		for key := range episodes {
        		keys = append(keys, key)
    		}
    		sort.Strings(keys)
			var topepisodes string 
    		for i := range keys {
        		key := keys[i]
        		value := episodes[key]
        		topepisodes += fmt.Sprintf("%s, Plays: %d \n", key, value)
			}	
			sendMessage(bot, update.Message.Chat.ID, topepisodes)
		case "geo":		
			geo := anchor.GetGeo(page, config)
			sendMessage(bot, update.Message.Chat.ID, getPercentString(geo))
		case "platforms":
			platformsStats := anchor.GetPlatform(page, config)
			sendMessage(bot, update.Message.Chat.ID, getPercentString(platformsStats))
		case "gender":
			genderStats := anchor.GetGender(page, config)
			sendMessage(bot, update.Message.Chat.ID, getPercentString(genderStats))
		case "age":
			ageStats := anchor.GetAge(page, config)
			sendMessage(bot, update.Message.Chat.ID, getPercentString(ageStats))
		case "help":			
			sendMessage(bot, update.Message.Chat.ID, 
				"List of commands:\n"+
				"Podcast Totals - /stats \n"+
				"Plays By Date - /plays \n"+
				"Top Episodes - /topepisodes \n"+
				"Geographic location - /geo \n"+
				"Listening platforms - /platforms \n"+
				"Gender - /gender \n"+
				"Age - /age")
		default:
			sendMessage(bot, update.Message.Chat.ID, "Command was not recognized.")
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func getPercentString(data map[string]float64) string{
	var res string 
    for key, value := range data {
		perc := value * 100
        res += fmt.Sprintf("%.2f %% - %s \n", perc, key)
	}
	return res
}