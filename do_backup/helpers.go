package do_backup

import (
	"time"

	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/backup/target"
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/telemetry"
)

func sleep(b backup.Backup, tel telemetry.Telemetry, i int) {
	if b.SleepBefore != 0 && i != 0 {
		logSleepStart(tel, b)
		time.Sleep(time.Duration(b.SleepBefore) * time.Second)
		logSleepDone(tel, b)
	}
}

func saveEventBackupErr(
	bl *backup_log.BackupLog,
	b backup.Backup,
	backupDuration time.Duration,
	err error,
	stdErr string,
) {
	bl.SaveEvent(
		b.Source.Name(), b.ID, "---", "---",
		int(backupDuration.Seconds()), 0, 0, 0, 0,
		err, stdErr)
}

func saveEventBackupMiddlewareErr(
	bl *backup_log.BackupLog,
	b backup.Backup,
	backupDuration time.Duration,
	backupMiddlewareDuration time.Duration,
	err error,
) {
	bl.SaveEvent(b.Source.Name(), b.ID, "---", "---",
		int(backupDuration.Seconds()),
		int(backupMiddlewareDuration.Seconds()),
		0, 0, 0,
		err, "")
}

func saveEventTargetMiddlewareErr(
	bl *backup_log.BackupLog,
	b backup.Backup,
	t target.Target,
	backupDuration time.Duration,
	backupMiddlewareDuration time.Duration,
	targetMiddlewareDuration time.Duration,
	err error,
) {
	bl.SaveEvent(b.Source.Name(), b.ID, t.Name(), t.ID,
		int(backupDuration.Seconds()),
		int(backupMiddlewareDuration.Seconds()),
		0,
		int(targetMiddlewareDuration.Seconds()),
		0,
		err, "")
}

func saveEventTargetSaveErr(
	bl *backup_log.BackupLog,
	b backup.Backup,
	t target.Target,
	backupDuration time.Duration,
	backupMiddlewareDuration time.Duration,
	targetMiddlewareDuration time.Duration,
	targetDuration time.Duration,
	size int64,
	err error,
) {
	bl.SaveEvent(
		b.Source.Name(), b.ID, t.Name(), t.ID,
		int(backupDuration.Seconds()),
		int(backupMiddlewareDuration.Seconds()),
		int(targetDuration.Seconds()),
		int(targetMiddlewareDuration.Seconds()),
		size,
		err, "")
}

func saveEventTargetSaveOK(
	bl *backup_log.BackupLog,
	b backup.Backup,
	t target.Target,
	backupDuration time.Duration,
	backupMiddlewareDuration time.Duration,
	targetMiddlewareDuration time.Duration,
	targetDuration time.Duration,
	size int64,
	err error,
) {
	bl.SaveEvent(
		b.Source.Name(), b.ID, t.Name(), t.ID,
		int(backupDuration.Seconds()),
		int(backupMiddlewareDuration.Seconds()),
		int(targetDuration.Seconds()),
		int(targetMiddlewareDuration.Seconds()),
		size,
		err, "")
}
