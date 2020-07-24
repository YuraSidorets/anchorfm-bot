# anchorfm-bot
Telegram bot for Anchor.fm written in [Go](https://golang.org) using [go-rod](https://github.com/go-rod/rod) for statistics scraping from Anchor, and [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) for Telegram.


## Setup
- Clone this repository or use `go get github.com/YuraSidorets/anchorfm-bot`
- Build everything: `go build`
- Follow Telegram instructions to create a new bot user and get your token. (https://core.telegram.org/bots#3-how-do-i-create-a-bot)
- Update required fields in the "config/config.json" with your values:
    - `"BotToken"` 
    - `"AnchorUser"`
    - `"AnchorPass"`
    - `"TimeRangeStart"` - start date for Anchor.fm to get your statistics (Unix Epoch)
    - `"WebStationId"` - you can find it in the https://anchor.fm/api/podcast/metadata response on Anchor page
    - `"UserId"`- you can find it in the https://anchor.fm/api/currentuser response on Anchor page
- Run anchorfm-bot using command `anchorfm-bot ./config/config.json`

## Supported commands

- Podcast Totals - /stats 
- Plays By Date - /plays 
- Top Episodes - /topepisodes 
- Geographic locations - /geo 
- Listening platforms - /platforms 
- Gender - /gender 
- Age - /age
- /help
