package do_backup

import (
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/target"
)

const PHASE_START = "START"
const PHASE_DONE = "DONE"
const PHASE_FAILED = "FAILED"

func metaLog(
	b *backup.Backup,
	t *target.Target,
	m *middleware.Middleware,
	method string,
	phase string,
	message string,
) {
	backup_id := ""
	scope := ""
	if b != nil {
		backup_id = b.ID
		scope = backup_id
	}
	target_id := ""
	if t != nil {
		target_id = t.ID
		scope = backup_id + "/" + target_id
	}
	middleware_id := ""
	if m != nil {
		middleware_id = m.Name()
		scope = scope + "+" + middleware_id
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
		Str("middleware_id", middleware_id).
		Msg(phase + "/" + method + "(" + scope + ")" + message_space + message)
}

func logBackupStart(b backup.Backup) {
	metaLog(&b, nil, nil, "BACKUP", PHASE_START,
		"")
}

func logBackupDone(b backup.Backup) {
	metaLog(&b, nil, nil, "BACKUP", PHASE_DONE,
		"Backup "+b.ID+" finished.")
}

func logBackupFailed(b backup.Backup, err error) {
	metaLog(&b, nil, nil, "BACKUP", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logTargetStart(b backup.Backup, t target.Target) {
	metaLog(&b, &t, nil, "TARGET", PHASE_START,
		"")
}

func logTargetDone(b backup.Backup, t target.Target) {
	metaLog(&b, &t, nil, "TARGET", PHASE_DONE,
		"Target "+b.ID+" finished.")
}

func logTargetFailed(b backup.Backup, t target.Target, err error) {
	metaLog(&b, &t, nil, "TARGET", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logTargetMiddlewareStart(b backup.Backup, t target.Target, m middleware.Middleware) {
	metaLog(&b, &t, &m, "TARGET_MIDDLEWARE", PHASE_START,
		"")
}

func logTargetMiddlewareDone(b backup.Backup, t target.Target, m middleware.Middleware) {
	metaLog(&b, &t, &m, "TARGET_MIDDLEWARE", PHASE_DONE,
		"")
}

func logTargetMiddlewareFailed(b backup.Backup, t target.Target, m middleware.Middleware, err error) {
	metaLog(&b, &t, &m, "TARGET_MIDDLEWARE", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}

func logBackupMiddlewareStart(b backup.Backup, m middleware.Middleware) {
	metaLog(&b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_START,
		"")
}

func logBackupMiddlewareDone(b backup.Backup, m middleware.Middleware) {
	metaLog(&b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_DONE,
		"")
}

func logBackupMiddlewareFailed(b backup.Backup, m middleware.Middleware, err error) {
	metaLog(&b, nil, &m, "BACKUP_MIDDLEWARE", PHASE_FAILED,
		"Backup "+b.ID+" failed: "+err.Error())
}
