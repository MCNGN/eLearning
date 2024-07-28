package domain

import "time"

type FileMetadata struct {
    FileSize  int64
    FileTopic string
    UploadDate time.Time
}

func NewFileMetadata(fileSize int64, fileTopic string, uploadDate time.Time) FileMetadata {
    return FileMetadata{
        FileSize:  fileSize,
        FileTopic: fileTopic,
        UploadDate: uploadDate,
    }
}
