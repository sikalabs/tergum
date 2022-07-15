package email

import (
	"github.com/rs/zerolog/log"
)

func logSkipped() {
	log.Info().
		Str("log-id", "notification-email-skipped").
		Msg("Email notification was skipped because backup was successful and SendOK is not set or set to false")
}

func logSend(email string) {
	log.Info().
		Str("log-id", "notification-email-sent").
		Msg("Email notification was sent (" + email + ")")
}

func logFailed(email string) {
	log.Error().
		Str("log-id", "notification-email-failed").
		Msg("Email notification failed to sent (" + email + ")")
}

func logError(errorMessage string) {
	log.Error().
		Str("log-id", "email-error").
		Msg(errorMessage)
}
