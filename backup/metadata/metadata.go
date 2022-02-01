package metadata

type BackupMetadataMeta struct {
	Version int `json:"version"`
}

type BackupMetadata struct {
	Meta               BackupMetadataMeta `json:"meta"`
	CeatedTimestapUnix int                `json:"created_timestamp_unix"`
}
