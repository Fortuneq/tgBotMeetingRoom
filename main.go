package main

import (
	"database/sql"
	"fmt"
	"log"
	_  "github.com/lib/pq"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Data struct{
	id int
	time string
	in_meet bool
}
var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("11:00", "\xE2\x98\x91"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("11:30", "\xE2\x98\x91"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("12:00", "\xE2\x98\x91"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("12:30", "\xE2\x98\x91"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("13:00", "\xE2\x98\x91"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("13:30", "\xE2\x98\x91"),
    ),
	tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("14:00", "\xE2\x98\x91"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("10:30", "\xE2\x98\x91"),
    ),
)


const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "537j04222"
  dbname   = "postgres"
)
func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
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
            if msg.Text != "hi" {
				data := `INSERT INTO users(id,in_time,in_meet) VALUES($1, $2, $3);`

				//Выполняем наш SQL запрос
				if _, err = db.Exec(data, Data.id,Data.in_time,Data.in_meet); err != nil {
					log.Println(err)
				}
			}

			if _, err := bot.Send(msg); err != nil {
                panic(err)
            }
        }
    }
}
