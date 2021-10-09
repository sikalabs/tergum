package backup_log

type BackupLogEvent struct {
	SourceName                string
	BackupID                  string
	TargetName                string
	TargetID                  string
	Success                   bool
	BackupDuration            int
	BackupMiddlewaresDuration int
	TargetDuration            int
	TargetMiddlewaresDuration int
	Error                     error
}

type BackupLog struct {
	ExtraName string
	Events    []BackupLogEvent
}

func (l *BackupLog) SaveEvent(
	sourceName string,
	backupID string,
	targetName string,
	targetID string,
	backupDuration int,
	backupMiddlewaresDuration int,
	targetDuration int,
	targetMiddlewaresDuration int,
	err error,
) {
	l.Events = append(l.Events, BackupLogEvent{
		SourceName:                sourceName,
		BackupID:                  backupID,
		TargetName:                targetName,
		TargetID:                  targetID,
		BackupDuration:            backupDuration,
		BackupMiddlewaresDuration: backupMiddlewaresDuration,
		TargetDuration:            targetDuration,
		TargetMiddlewaresDuration: targetMiddlewaresDuration,
		Success:                   err == nil,
		Error:                     err,
	})

}

func (l BackupLog) GlobalSuccess() bool {
	for _, log := range l.Events {
		if !log.Success {
			return false
		}
	}
	return true
}

func (l BackupLog) GlobalSuccessString() string {
	if l.GlobalSuccess() {
		return "OK"
	}
	return "ERR"
}
