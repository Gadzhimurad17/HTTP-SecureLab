package models

import (
	"io"
	"time"

	"github.com/google/uuid"
)

type ContentType string

const (
	Avatar ContentType = "avatar"
)

type File struct {
	ID         uuid.UUID `json:"id"`
	UploadedBy uuid.UUID `json:"uploaded_by"`

	Purpose      ContentType `json:"purpose"`
	StorageKey   string      `json:"storage_key"`
	OriginalName string      `json:"original_name"`
	ContentType  string      `json:"content_type"`
	SizeBytes    int64       `json:"size_bytes"`
	CreatedAt    time.Time   `json:"created_at"`
}

type ContentMetadata struct {
	URL      string
	File     io.Reader
	Purpose  ContentType
	MimeType string
	Size     int64
}
