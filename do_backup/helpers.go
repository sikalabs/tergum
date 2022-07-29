package do_backup

import (
	"time"

	"github.com/sikalabs/tergum/backup"
	"github.com/sikalabs/tergum/telemetry"
)

func sleep(b backup.Backup, tel telemetry.Telemetry, i int) {
	if b.SleepBefore != 0 && i != 0 {
		logSleepStart(tel, b)
		time.Sleep(time.Duration(b.SleepBefore) * time.Second)
		logSleepDone(tel, b)
	}
}
