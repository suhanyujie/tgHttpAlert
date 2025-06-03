package service

import (
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pkgErr "github.com/pkg/errors"
)

const (
	baseUrl         = "https://api.telegram.org/bot%s/%s"
	botToken        = ""
	chatId001 int64 = 0
)

var (
	botClient *tgbotapi.BotAPI
	pkgOnce   sync.Once
)

func init() {
	pkgOnce.Do(func() {
		botClient = newBot()
	})
}

func newBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	return bot
}

func getBot() *tgbotapi.BotAPI {
	return botClient
}

type AlertMsg struct {
	env        string
	serverFlag string
	app        string
	msg        string
}

func NewAlertMsg(env, serverFlag, app, msg string) *AlertMsg {
	return &AlertMsg{
		env, serverFlag, app, msg,
	}
}

func (am *AlertMsg) Send() error {
	if am.env == "" {
		return pkgErr.New("param env is null")
	}
	if am.serverFlag == "" {
		return pkgErr.New("param serverFlag is null")
	}
	if am.app == "" {
		return pkgErr.New("param app is null")
	}
	if am.msg == "" {
		return pkgErr.New("param msg is null")
	}
	msg := fmt.Sprintf("[%s][%s][%s] %s", am.env, am.serverFlag, am.app, am.msg)
	return sendMsg(msg)
}

// 发送告警消息到群里
func sendMsg(content string) error {
	msg := tgbotapi.NewMessage(chatId001, content)

	if _, err := getBot().Send(msg); err != nil {
		return pkgErr.Wrapf(err, "bot send error")
	}
	return nil
}
