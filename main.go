package main

import (
	"database/sql"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("11:00", "11:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:30", "11:30"),
		tgbotapi.NewInlineKeyboardButtonData("12:00", "12:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:30","12:30"),

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

		tgbotapi.NewInlineKeyboardButtonData("17:00","17:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:30", "17:30"),
		tgbotapi.NewInlineKeyboardButtonData("18:00", "18:00"),
		tgbotapi.NewInlineKeyboardButtonData("18:30", "18:30"),

	),
	
)


var Marcup = tgbotapi.NewInlineKeyboardMarkup()

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

			switch update.Message.Command() {
			case "help":
				msg.Text = "Бот понимает команды /cancel,/show and /start."
			case "start":
				msg.Text = "Хочешь занять переговорку?"
				msg.ReplyMarkup = numericKeyboard
				//tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ChatConfig().ChatID,update.CallbackQuery.Message.MessageID,Marcup)
				
			case "show":
				msg.Text = "Все записи в переговорку"
				show :=  `SELECT * FROM meetings
				WHERE in_meet = $1`
				rows, err := db.Query(show,false)
				if err != nil {
					log.Fatal(err)
				}
				for rows.Next() {
					var id int
					var is_bool bool
					if err := rows.Scan(&id,&msg.Text,&is_bool); err != nil {
						log.Fatal(err)
					}
					if _, err := bot.Send(msg); err != nil {
						panic(err)
					}
				}
					
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

			data := `INSERT INTO meetings(id,in_time,in_meet) VALUES($1, $2, $3);`

			//Выполняем наш SQL запрос
			if _, err = db.Exec(data, update.CallbackQuery.Message.MessageID, update.CallbackQuery.Data, true); err != nil {
				msg.Text = fmt.Sprint(err)
				msg.Text = "Сожалеем но на это время уже кто-то записан"
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			} else {			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			msg.Text = "Вы были записаны,время указано выше ,удачи в переговорке :)"
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}}
			

		}
	}
}
