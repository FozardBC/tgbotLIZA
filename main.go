package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	url := os.Getenv("URL")
	token := "7875042854:AAFX1Iql-Db1i_4JAb7eMV3sABRPxnfkrso"

	fmt.Print("\nURL:" + url + "\n")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go server()
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msgText := strings.ToLower(update.Message.Text)

		switch {
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "start":
				// TODO: Замените текст ниже на своё приветственное сообщение
				reply := tgbotapi.NewVideoNote(update.Message.Chat.ID, 7, tgbotapi.FilePath("./shreck2.mp4"))
				bot.Send(reply)

				msg := tgbotapi.NewVideoNote(update.Message.Chat.ID, 20, tgbotapi.FilePath("./shreckZAGAD.mp4"))
				bot.Send(msg)

				msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, `Где плывут те, кто не умеет

Он хотел взлететь, но выбрал воду.
Он не был рыбой, но исчез без следа.

Там, где железо сторожит память,
и волны лижут камень,
тень однажды показала нам,
что даже перья тонут.

Спустись туда, где время течёт,
и тишина знает, о чём ты молчишь
там найдешь ответ, поспеши мне его сообщить`)
				bot.Send(msg2)

			case "bird":
				// TODO: Замените текст ниже на нужное сообщение
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, `А ночь глядит сквозь зыбкий свет,
Под кожей дышит тёплый след.
Отголоски слов в тумане тонут,
Ступень за ступенью нас ведут.
Ты помнишь, как дрожала гладь,
Он выбрал вниз, а не назад.
Лишь взгляд молчит, не отвести —

А в реке — тень без пути.
Нас ждёт та точка на реке,
Дни промелькнули налегке.
Рядом всё, и всё как будто вновь.
Если дойдёшь — найдёшь любовь.
Йди — я рядом, я зову.`)
				bot.Send(reply)

			case "end":
				video := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FilePath("./sunboy.mp3"))
				bot.Send(video)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, `Когда круг замкнётся,
и шагов больше нет…
Там, где началось всё —
ждет завершение след.

В стенах, что шептали,
в подушке, что знала,
ответ ты найдёшь,
где душа твоя спала.`)
				bot.Send(msg)
			}

		case msgText == "апостол андрей":
			// Отправка MP3 как голосовое сообщение

			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, `Те числа, что ты разыскала прежде —
не просто шёпот прошлого,
но ключ с двумя лезвиями.

Одна дверь ведёт туда, где было,
другая — туда, где быть должно.

Разница между ними — вторая часть ответа.
Найди её. И не ошибись.`)

			bot.Send(msg1)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "http://"+url)
			bot.Send(msg)

		}
	}
}

func server() {

	srv := http.Server{
		Addr: "0.0.0.0:8080",
	}

	fs := http.FileServer(http.Dir("./html"))

	http.Handle("/", fs)

	log.Println("Сервер запущен на http://0.0.0.0:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
