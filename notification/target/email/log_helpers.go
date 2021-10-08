package email

import (
	"github.com/rs/zerolog/log"
)

func logSkipped() {
	log.Info().
		Str("log-id", "notification-email-skipped").
		Msg("Email notification was skipped because backup was successful and SendOK is not set or set to false")
}

func logSend() {
	log.Info().
		Str("log-id", "notification-email-sent").
		Msg("Email notification was sent")
}
