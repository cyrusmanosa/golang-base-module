package uploadFile

type UploadTools struct {
	MaxFileSize      int
	AllowedFileTypes []string
}

type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}
