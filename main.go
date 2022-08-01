package main

import (
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("7:00", "\xE2\x98\x91"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("7:30", "7:30"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("8:00", "7:00"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("8:30", "7:30"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("9:00", "7:00"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("9:30", "7:30"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("10:00", "7:00"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("10:30", "7:30"),
    ),
)

func main() {
    bot, err := tgbotapi.NewBotAPI("5420203457:AAHxa3dlya-NkW4i8L62mbgkTEe8Mfo9OVY")
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    // Loop through each update.
    for update := range updates {
        // Check if we've gotten a message update.
        if update.Message != nil {
            // Construct a new message from the given chat ID and containing
            // the text that we received.
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

            // If the message was open, add a copy of our numeric keyboard.
       /*     switch update.Message.Text {
            case "open":
                msg.ReplyMarkup = numericKeyboard
            }*/

			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /cancel,/show and /start."
			case "start":
				msg.Text = "Хочешь занять переговорку?"
				msg.ReplyMarkup = numericKeyboard
			case "show":
				msg.Text = "Все записи в переговорку"
			case "cancel":
				msg.Text = "Хочешь отменить запись в переговорку?"
			default:
				msg.Text = "Не знаю такой команды "
			}
            // Send the message.
            if _, err = bot.Send(msg); err != nil {
                panic(err)
            }
        } else if update.CallbackQuery != nil {
            // Respond to the callback query, telling Telegram to show the user
            // a message with the data received.
            callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
            if _, err := bot.Request(callback); err != nil {
                panic(err)
            }

            // And finally, send a message containing the data received.
            msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
            if _, err := bot.Send(msg); err != nil {
                panic(err)
            }
        }
    }
}
