package do_backup

import (
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/target"
	"github.com/sikalabs/tergum/telemetry"
)

const PHASE_START = "START"
const PHASE_DONE = "DONE"
const PHASE_FAILED = "FAILED"

func metaLog(
	tel telemetry.Telemetry,
	b *backup.Backup,
	t *target.Target,
	m *middleware.Middleware,
	method string,
	phase string,
	message string,
) {
	backup_id := ""
	source_name := ""
	scope := ""
	if b != nil {
		backup_id = b.ID
		source_name = b.Source.Name()
		scope = backup_id
	}
	target_id := ""
	target_name := ""
	if t != nil {
		target_id = t.ID
		target_name = t.Name()
		scope = backup_id + "/" + target_id
	}
	middleware_name := ""
	if m != nil {
		middleware_name = m.Name()
		scope = scope + "+" + middleware_name
	}

	message_space := ""
	if message != "" {
		message_space = " "
	}

	log.Info().
		Str("method", method).
		Str("phase", phase).
		Str("backup_id", backup_id).
		Str("target_id", target_id).
		Str("source_name", source_name).
		Str("target_name", target_name).
		Str("middleware_name", middleware_name).
		Msg(phase + "/" + method + "(" + scope + ")" + message_space + message)
	data := map[string]string{
		"method":          method,
		"phase":           phase,
		"backup_id":       backup_id,
		"target_id":       target_id,
		"source_name":     source_name,
		"target_name":     target_name,
		"middleware_name": middleware_name,
	}
	jsonData, _ := json.Marshal(data)
	tel.SendEvent(strings.ToLower("do/"+method+"/"+phase), string(jsonData))
}

func logBackupStart(tel telemetry.Telemetry, b backup.Backup) {
	metaLog(tel, &b, nil, nil, "BACKUP", PHASE_START,
		"")
}

func logBackupDone(tel telemetry.Telemetry, b backup.Backup) {
	metaLog(tel, &b, nil, nil, "BACKUP", PHASE_DONE,
		"Backup "+b.ID+" finished.")
}

func logBackupFailed(tel telemetry.Telemetry, b backup.Backup, err error) {
	metaLog(tel, &b, nil, nil, "BACKUP", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logTargetStart(tel telemetry.Telemetry, b backup.Backup, t target.Target) {
	metaLog(tel, &b, &t, nil, "TARGET", PHASE_START,
		"")
}

func logTargetDone(tel telemetry.Telemetry, b backup.Backup, t target.Target) {
	metaLog(tel, &b, &t, nil, "TARGET", PHASE_DONE,
		"Target "+b.ID+" finished.")
}

func logTargetFailed(tel telemetry.Telemetry, b backup.Backup, t target.Target, err error) {
	metaLog(tel, &b, &t, nil, "TARGET", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logTargetMiddlewareStart(tel telemetry.Telemetry, b backup.Backup, t target.Target, m middleware.Middleware) {
	metaLog(tel, &b, &t, &m, "TARGET_MIDDLEWARE", PHASE_START,
		"")
}

func logTargetMiddlewareDone(tel telemetry.Telemetry, b backup.Backup, t target.Target, m middleware.Middleware) {
	metaLog(tel, &b, &t, &m, "TARGET_MIDDLEWARE", PHASE_DONE,
		"")
}

func logTargetMiddlewareFailed(tel telemetry.Telemetry, b backup.Backup, t target.Target, m middleware.Middleware, err error) {
	metaLog(tel, &b, &t, &m, "TARGET_MIDDLEWARE", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logBackupMiddlewareStart(tel telemetry.Telemetry, b backup.Backup, m middleware.Middleware) {
	metaLog(tel, &b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_START,
		"")
}

func logBackupMiddlewareDone(tel telemetry.Telemetry, b backup.Backup, m middleware.Middleware) {
	metaLog(tel, &b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_DONE,
		"")
}

func logBackupMiddlewareFailed(tel telemetry.Telemetry, b backup.Backup, m middleware.Middleware, err error) {
	metaLog(tel, &b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}
