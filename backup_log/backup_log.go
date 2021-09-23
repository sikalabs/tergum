package backup_log

type BackupLogEvent struct {
	BackupID string
	TargetID string
	Success  bool
	Error    error
}

type BackupLog struct {
	ExtraName string
	Events    []BackupLogEvent
}

func (l *BackupLog) SaveEvent(
	backupID string,
	targetID string,
	err error,
) {
	if err == nil {
		l.Events = append(l.Events, BackupLogEvent{
			BackupID: backupID,
			TargetID: targetID,
			Success:  true,
			Error:    nil,
		})
	} else {
		l.Events = append(l.Events, BackupLogEvent{
			BackupID: backupID,
			TargetID: targetID,
			Success:  false,
			Error:    err,
		})
	}

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
