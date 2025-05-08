package telegram

import (
	"github.com/rs/zerolog/log"
)

func logSkipped() {
	log.Info().
		Str("log-id", "notification-telegram-skipped").
		Msg("Telegram notification was skipped because backup was successful and SendOK is not set or set to false")
}

func logSend() {
	log.Info().
		Str("log-id", "notification-telegram-sent").
		Msg("Telegram notification was sent")
}
