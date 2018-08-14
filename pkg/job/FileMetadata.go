package job

// FileMetadata represents the file metadata for user-server data set syncing
type FileMetadata struct {
	ModTime int64  `json:"mod_time"`
	Hash    string `json:"hash"`
}
