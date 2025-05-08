package telegram_output

import (
	"bytes"
	"strconv"

	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/utils/file_size_utils"
)

func BackupLogTelegram(l backup_log.BackupLog) string {
	out := new(bytes.Buffer)
	out.WriteString("=== log ===\n")
	out.WriteString("\n")

	for _, log := range l.Events {
		var strStatus, emojiStatus string

		if log.Success {
			strStatus = "OK"
			emojiStatus = "\u2705"
		} else {
			strStatus = "ERROR"
			emojiStatus = "\u274C"
		}

		out.WriteString("Success: " + strStatus + " " + emojiStatus + "\n")
		out.WriteString("Backup: " + log.SourceName + ": " + log.BackupID + "\n")
		out.WriteString("Backup Time: " + strconv.Itoa(log.BackupDuration) + "s" +
			" (+" + strconv.Itoa(log.BackupMiddlewaresDuration) + "s)" + "\n")
		if !(log.TargetName == "---" && log.TargetID == "---") {
			out.WriteString("Target: " + log.TargetName + ": " + log.TargetID + "\n")
		}
		out.WriteString("Upload Time: " + strconv.Itoa(log.TargetDuration) + "s" +
			" (+" + strconv.Itoa(log.TargetMiddlewaresDuration) + "s)" + "\n")
		// Only show file size if target is not empty
		if !(log.TargetName == "---" && log.TargetID == "---") {
			out.WriteString("File Size: " + file_size_utils.PrettyFileSize(log.TargetFileSize) + "\n")
		}
		if log.Error != nil {
			out.WriteString("Error: " + log.Error.Error() + "\n")
		}
		out.WriteString("Time Total: " + strconv.Itoa(log.TotalDuration()) + "s" + "\n")
		out.WriteString("\n")
	}
	return out.String()
}

func BackupLogToTelegramString(l backup_log.BackupLog) string {
	return BackupLogTelegram(l)
}

func BackupErrorLogTelegram(l backup_log.BackupLog) string {
	noErrors := true
	out := new(bytes.Buffer)
	out.WriteString("=== errors ===\n")
	out.WriteString("\n")
	for _, log := range l.Events {
		if log.Success {
			continue
		}
		noErrors = false
		out.WriteString("Backup: " + log.BackupID + "\n")
		out.WriteString("Target: " + log.TargetID + "\n")
		out.WriteString("Error: " + log.StdErr + "\n")
		out.WriteString("\n")
	}
	if noErrors {
		return ""
	}
	return out.String()
}

func BackupErrorLogToTelegramString(l backup_log.BackupLog) string {
	return BackupErrorLogTelegram(l)
}
