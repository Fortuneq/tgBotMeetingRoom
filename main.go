package main

import (
	"database/sql"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "537j04222"
	dbname   = "postgres"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("11:00", "11:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:30", "11:30"),
		tgbotapi.NewInlineKeyboardButtonData("12:00", "12:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:30", "12:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("13:00", "13:00"),
		tgbotapi.NewInlineKeyboardButtonData("13:30", "13:30"),
		tgbotapi.NewInlineKeyboardButtonData("14:00", "14:00"),
		tgbotapi.NewInlineKeyboardButtonData("14:30", "14:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("15:00", "15:00"),
		tgbotapi.NewInlineKeyboardButtonData("15:30", "15:30"),
		tgbotapi.NewInlineKeyboardButtonData("16:00", "16:00"),
		tgbotapi.NewInlineKeyboardButtonData("16:30", "16:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("17:00", "17:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:30", "17:30"),
		tgbotapi.NewInlineKeyboardButtonData("18:00", "18:00"),
		tgbotapi.NewInlineKeyboardButtonData("18:30", "18:30"),
	),
)


var cancelnumericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("11:00", "11:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:30", "11:30"),
		tgbotapi.NewInlineKeyboardButtonData("12:00", "12:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:30", "12:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("13:00", "13:00"),
		tgbotapi.NewInlineKeyboardButtonData("13:30", "13:30"),
		tgbotapi.NewInlineKeyboardButtonData("14:00", "14:00"),
		tgbotapi.NewInlineKeyboardButtonData("14:30", "14:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("15:00", "15:00"),
		tgbotapi.NewInlineKeyboardButtonData("15:30", "15:30"),
		tgbotapi.NewInlineKeyboardButtonData("16:00", "16:00"),
		tgbotapi.NewInlineKeyboardButtonData("16:30", "16:30"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("17:00", "17:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:30", "17:30"),
		tgbotapi.NewInlineKeyboardButtonData("18:00", "18:00"),
		tgbotapi.NewInlineKeyboardButtonData("18:30", "18:30"),
	),
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

	bot, err := tgbotapi.NewBotAPI("5420203457:AAHxa3dlya-NkW4i8L62mbgkTEe8Mfo9OVY")
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = true

	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)



				switch update.Message.Command() {
				case "show":
					msg.Text = "Вот все доступные места в переговорку"
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
					show := `SELECT * FROM meetings
					WHERE in_meet = $1`
					rows, err := db.Query(show, false)
					if err != nil {
						log.Fatal(err)
					}

					//showKeyboard := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(msg.Text,msg.Text)})
					//msg.ReplyMarkup = showKeyboard
					for rows.Next() {
						var id int
						var comment string
						var time string
						var in_meet bool
						if err := rows.Scan(&id, &comment, &time, &in_meet); err != nil {
							log.Fatal(err)
						}
						msg.Text = time + " " + comment 
						//showKeyboard.InlineKeyboard = append(showKeyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton())
						//	msg.ReplyMarkup = showKeyboard
						//
						if _, err := bot.Send(msg); err != nil {
							log.Println(err)
						}
					}

					continue
				
				case "show_ordered":
					msg.Text = "Вот все доступные места в переговорку"
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
					show := `SELECT * FROM meetings
					WHERE in_meet = $1`
					rows, err := db.Query(show, true)
					if err != nil {
						log.Fatal(err)
					}

					//showKeyboard := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(msg.Text,msg.Text)})
					//msg.ReplyMarkup = showKeyboard
					for rows.Next() {
						var id int
						var comment string
						var time string
						var in_meet bool
						if err := rows.Scan(&id, &comment, &time, &in_meet); err != nil {
							log.Fatal(err)
						}
						msg.Text = time + " " + comment 
						//showKeyboard.InlineKeyboard = append(showKeyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton())
						//	msg.ReplyMarkup = showKeyboard
						//
						if _, err := bot.Send(msg); err != nil {
							log.Println(err)
						}
					}

					continue
				case "help":
					msg.Text = "Я обрабатываю команды  /start,/show,/cancel"

					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
					continue

				case "cancel":
					dbdel := `UPDATE meetings 
					SET in_meet = false 
					WHERE in_time = $1`
					msg.ReplyMarkup = cancelnumericKeyboard
					msg.Text = "ВВедите время в формате h:m для удаления"
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
					msg.ReplyToMessageID = update.Message.MessageID
					
					msg.Text = update.Message.Text


					_, err := db.Exec(dbdel, update.Message.Text)
					if err != nil{
						panic(err)
					}
				
					continue

				case "start":
					msg.Text = "Выберите себе Время на запись"
					msg.ReplyMarkup = numericKeyboard
					msg.ReplyMarkup = " "

					msg.ReplyToMessageID = update.Message.MessageID

					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}

					msg.ReplyMarkup = msg.Text
					msg.Text = "ВВедите комментарий"
					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}



					continue
				}

			msg.ReplyToMessageID = update.Message.MessageID

			// Okay, we're sending our message off! We don't care about the message
			// we just sent, so we'll discard it.
			if _, err := bot.Send(msg); err != nil {
				// Note that panics are a bad way to handle errors. Telegram can
				// have service outages or network errors, you should retry sending
				// messages or more gracefully handle failures.
				log.Println(err)
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

			//msg.Text = "ВВедите комментарий"


			


			//data := `INSERT INTO meetings(id,in_time,in_meet) VALUES($1, $2, $3);`
			data := `UPDATE meetings 
			SET in_meet = true 
			WHERE in_time = $1`
			//Выполняем наш SQL запрос

			dbcheck := `SELECT in_meet from meetings WHERE in_time = $1`
			var in_meet bool
			row := db.QueryRow(dbcheck,update.CallbackQuery.Data)
			switch err := row.Scan(&in_meet); err {
				case sql.ErrNoRows:
					  fmt.Println("No rows were returned!")
				case nil:
					if in_meet == true{
						msg.Text = fmt.Sprint(err)
						msg.Text = "Сожалеем но на это время уже кто-то записан"
							if _, err := bot.Send(msg); err != nil {
							panic(err)
							}
					} else {
						if _, err = db.Exec(data, update.CallbackQuery.Data); err != nil {
							msg.Text = fmt.Sprint(err)
							msg.Text = "Сожалеем но на это время уже кто-то записан"
							if _, err := bot.Send(msg); err != nil {
								panic(err)
							}
						} else {
							if _, err := bot.Send(msg); err != nil {
								panic(err)
							}
							msg.Text = "Вы были записаны,время указано выше ,удачи в переговорке :)"
							if _, err := bot.Send(msg); err != nil {
								panic(err)
							}
						}
				}
			}
		}
	}
}