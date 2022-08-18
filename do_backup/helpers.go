package do_backup

import (
	"time"

	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/backup/middleware"
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

type DoBackupProcess struct {
	Telemetry                telemetry.Telemetry
	BackupLog                *backup_log.BackupLog
	Backup                   backup.Backup
	Middleware               middleware.Middleware
	Target                   target.Target
	TargetSize               int64
	BackupDuration           time.Duration
	BackupMiddlewareDuration time.Duration
	TargetMiddlewareDuration time.Duration
	TargetDuration           time.Duration
	BackupError              error
	BackupStdError           string
	BackupMiddlewareError    error
	TargetMiddlewareError    error
	TargetError              error
	backupStart              time.Time
	backupMiddlewareStart    time.Time
	targetMiddlewareStart    time.Time
	targetStart              time.Time
}

func (d *DoBackupProcess) BackupStart() {
	d.backupStart = time.Now()
	logBackupStart(d.Telemetry, d.Backup)
}

func (d *DoBackupProcess) BackupFinish() {
	d.BackupDuration = time.Since(d.backupStart)
}

func (d *DoBackupProcess) BackupMiddlewareStart() {
	d.backupMiddlewareStart = time.Now()
	logBackupMiddlewareStart(d.Telemetry, d.Backup, d.Middleware)
}

func (d *DoBackupProcess) BackupMiddlewareFinish() {
	d.BackupMiddlewareDuration = time.Since(d.backupMiddlewareStart)
}

func (d *DoBackupProcess) TargetMiddlewareStart() {
	d.targetMiddlewareStart = time.Now()
	logTargetMiddlewareStart(d.Telemetry, d.Backup, d.Target, d.Middleware)
}

func (d *DoBackupProcess) TargetMiddlewareFinish() {
	d.TargetMiddlewareDuration = time.Since(d.targetMiddlewareStart)
}

func (d *DoBackupProcess) TargetStart() {
	d.targetStart = time.Now()
	logTargetStart(d.Telemetry, d.Backup, d.Target)
}

func (d *DoBackupProcess) TargetFinish() {
	d.TargetDuration = time.Since(d.targetStart)
}

func (d DoBackupProcess) BackupErr() {
	b := d.Backup
	d.BackupLog.SaveEvent(
		b.Source.Name(), b.ID, "---", "---",
		int(d.BackupDuration.Seconds()), 0, 0, 0, 0,
		d.BackupError, d.BackupStdError)
	logBackupFailed(d.Telemetry, b, int(d.BackupDuration.Seconds()), d.BackupError)
}

func (d DoBackupProcess) BackupOK() {
	logBackupDone(d.Telemetry, d.Backup, int(d.BackupDuration.Seconds()))
}

func (d DoBackupProcess) BackupMiddlewareErr() {
	b := d.Backup
	m := d.Middleware
	d.BackupLog.SaveEvent(b.Source.Name(), b.ID, "---", "---",
		int(d.BackupDuration.Seconds()),
		int(d.BackupMiddlewareDuration.Seconds()),
		0, 0, 0,
		d.BackupMiddlewareError, "")
	logBackupMiddlewareFailed(
		d.Telemetry, b, m,
		int(d.BackupMiddlewareDuration.Seconds()),
		d.BackupMiddlewareError)
}

func (d DoBackupProcess) BackupMiddlewareOK() {
	b := d.Backup
	m := d.Middleware
	logBackupMiddlewareDone(d.Telemetry, b, m, int(d.BackupMiddlewareDuration.Seconds()))
}

func (d DoBackupProcess) TargetMiddlewareErr() {
	b := d.Backup
	t := d.Target
	m := d.Middleware
	d.BackupLog.SaveEvent(b.Source.Name(), b.ID, t.Name(), t.ID,
		int(d.BackupDuration.Seconds()),
		int(d.BackupMiddlewareDuration.Seconds()),
		0,
		int(d.TargetMiddlewareDuration.Seconds()),
		0,
		d.TargetMiddlewareError, "")

	logTargetMiddlewareFailed(d.Telemetry, b, t, m,
		int(d.TargetMiddlewareDuration.Seconds()),
		d.TargetMiddlewareError)
}

func (d DoBackupProcess) TargetMiddlewareOK() {
	b := d.Backup
	t := d.Target
	m := d.Middleware
	logTargetMiddlewareDone(d.Telemetry, b, t, m, int(d.TargetMiddlewareDuration.Seconds()))
}

func (d DoBackupProcess) SaveErr() {
	b := d.Backup
	t := d.Target
	d.BackupLog.SaveEvent(
		b.Source.Name(), b.ID, t.Name(), t.ID,
		int(d.BackupDuration.Seconds()),
		int(d.BackupMiddlewareDuration.Seconds()),
		int(d.TargetDuration.Seconds()),
		int(d.TargetMiddlewareDuration.Seconds()),
		d.TargetSize,
		d.TargetError, "")
	logTargetFailed(d.Telemetry, b, t, int(d.TargetDuration.Seconds()), d.TargetError)

}

func (d DoBackupProcess) SaveOK() {
	b := d.Backup
	t := d.Target
	d.BackupLog.SaveEvent(
		b.Source.Name(), b.ID, t.Name(), t.ID,
		int(d.BackupDuration.Seconds()),
		int(d.BackupMiddlewareDuration.Seconds()),
		int(d.TargetDuration.Seconds()),
		int(d.TargetMiddlewareDuration.Seconds()),
		d.TargetSize,
		d.TargetError, "")
	logTargetDone(d.Telemetry, b, t, int(d.TargetDuration.Seconds()))
}
