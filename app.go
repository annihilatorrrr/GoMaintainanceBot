// (c) HeimanPictures

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

var UPDATE_CHANNEL = os.Getenv("UPDATE_CHANNEL")

var SUPPORT_GROUP = os.Getenv("SUPPORT_GROUP")

func main() {
	b, err := gotgbot.NewBot((os.Getenv("BOT_TOKEN")), &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewMessage(filters.All, all))
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	updater.Idle()
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s.\nI Am A -- Bot. I Am Sorry To Say That This Bot Is Under Maintainance.. Plese Wait Till My Boss Intimate About My Return.\nDeveloped By @HeimanCreation.", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "UPDATE CHANNEL", Url: UPDATE_CHANNEL},
				{Text: "SUPPORT GROUP", Url: SUPPORT_GROUP},
			}},
		},
	})
	if err != nil {
		fmt.Println("failed to send /start: " + err.Error())
	}
	return nil
}

func all(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s.\nI Am Under Maintainance.. So Use Me Later After I Am Back. Intimation Will Be Provided On Our Channel.\nThank You 🥰🥰", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "SUPPORT GROUP", Url: SUPPORT_GROUP},
			}},
		},
	})
	if err != nil {
		fmt.Println("failed to set all filter: " + err.Error())
	}
	return nil
}