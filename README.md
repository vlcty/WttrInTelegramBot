# WttrIntelegramBot

## General info

A Telegram Bot for https://wttr.in/ by @igor_chubin    
His GitHub Page: https://github.com/chubin

This bot provides two commands:

* `/info` Prints a short info about the bot
* `/weather` Requests the current weather through wttr.in. It takes the place you want to know the weather of as argument. Example: `/weather moon`. Then the bot will query wttr.in and send you the picture.

## Build and run

This bot needs an Telegram Bot API key passed by the environment variable `BOTTOKEN`

Build it with:

```
$ go get gopkg.in/telegram-bot-api.v4
$ go build -o WttrInBot WttrIntelegramBot.go
```

Start it with:

```
$ BOTTOKEN="yourBotToken" ./WttrInBot
```

## See it running

Go add my running bot. You can find it under @WttrInBot in Telegram.
