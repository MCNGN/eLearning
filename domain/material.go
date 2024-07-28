package domain

type Material struct {
    File         File
    FileMetadata FileMetadata
}

func NewLearningMaterial(file File, fileMetadata FileMetadata) Material {
    return Material{
        File:         file,
        FileMetadata: fileMetadata,
    }
}