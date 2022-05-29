package slack_webhook

import (
	"github.com/rs/zerolog/log"
)

func logSkipped() {
	log.Info().
		Str("log-id", "notification-slack-skipped").
		Msg("Slack notification was skipped because backup was successful and SendOK is not set or set to false")
}

func logSend() {
	log.Info().
		Str("log-id", "notification-slack-sent").
		Msg("Slack notification was sent")
}
