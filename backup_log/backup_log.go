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
	TargetFileSize            int64
	TargetMiddlewaresDuration int
	Error                     error
	StdErr                    string
}

type BackupLog struct {
	ExtraName string
	Events    []BackupLogEvent
}

func (e *BackupLogEvent) TotalDuration() int {
	return e.BackupDuration + e.BackupMiddlewaresDuration +
		e.TargetDuration + e.TargetMiddlewaresDuration
}

func (l *BackupLog) SaveEventRaw(ev BackupLogEvent) {
	l.Events = append(l.Events, ev)
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
	targetFileSize int64,
	err error,
	stdErr string,
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
		TargetFileSize:            targetFileSize,
		Success:                   err == nil,
		Error:                     err,
		StdErr:                    stdErr,
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
