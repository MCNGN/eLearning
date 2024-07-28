package domain

import "time"

type File struct {
    ID         string
    FileName   string
    FileSize   int64
    FileTopic  string
    UploadDate time.Time
}

func NewFile(id, fileName, fileTopic string, fileSize int64) *File {
    return &File{
        ID:         id,
        FileName:   fileName,
        FileSize:   fileSize,
        FileTopic:  fileTopic,
        UploadDate: time.Now(),
    }
}