package main

import (
    "log"
    "os"
    "strings"
    "gopkg.in/telegram-bot-api.v4"
    "io"
    "net/http"
    "errors"
)

func main() {
    log.Println("Starting Wttr.in Telegram Bot")

    bot, err := tgbotapi.NewBotAPI(GetBotToken())

	if err != nil {
		log.Panic(err)
	}

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

        if update.Message.IsCommand() {
            command := update.Message.Text

            log.Printf("[%s] %s", update.Message.From.UserName, command)

            if ( command == "/info") {
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Type \"/weather Place\" to get the current weather")
                bot.Send(msg)
            } else if strings.HasPrefix(command, "/weather") {
                if err := FetchWttrPicture(command); err != nil {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
                    bot.Send(msg)
                } else {
                    msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "/tmp/wttr.png")
                    bot.Send(msg)
                }
            } else {
                // Do nothing
            }
        }
	}
}

func FetchWttrPicture(place string) error {
    place = strings.Replace(place, "/weather", "", -1)
    place = strings.Trim(place, " ")

    if len(place) == 0 {
        return errors.New("No place given")
    }

    log.Printf("Place is %s\n", place)

    img, err := os.Create("/tmp/wttr.png")

    if err != nil {
        return err
    }

    defer img.Close()

    resp, err := http.Get("https://wttr.in/" + place + ".png")
    defer resp.Body.Close()

    if err != nil {
        return err
    }

    b, err := io.Copy(img, resp.Body)
    log.Println("Bytes: ", b)

    if err != nil {
        return err
    }

    return nil
}

func GetBotToken() string {
    botToken := os.Getenv("BOTTOKEN")

    if len(botToken) != 0 {
        log.Printf("Bot token is: %s", botToken)
    } else {
        log.Fatal("Did not get a bot token")
    }

    return botToken
}
