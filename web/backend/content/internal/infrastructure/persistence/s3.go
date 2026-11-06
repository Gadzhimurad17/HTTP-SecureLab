package persistence

import (
	"context"
	"fmt"

	"github.com/Gagzhimurad17/NullPoint/web/backend/content/internal/domain/models"
	"github.com/minio/minio-go/v7"
)

type S3Repository struct {
	mc     *minio.Client
	bucket string
}

func NewS3Repository(mc *minio.Client, bucket string) *S3Repository {
	return &S3Repository{
		mc:     mc,
		bucket: bucket,
	}
}

func (r *S3Repository) getPath(userID string, cType models.ContentType) string {
	return fmt.Sprintf("%s/%s", cType, userID)
}
func (r *S3Repository) UploadFile(ctx context.Context, metadata *models.ContentMetadata) error {
	opts := minio.PutObjectOptions{
		ContentType: metadata.MimeType,
	}
	_, err := r.mc.PutObject(ctx, r.bucket, metadata.URL, metadata.File, metadata.Size, opts)
	if err != nil {
		return models.ErrS3Operation
	}

	return nil
}

func (r *S3Repository) Delete(ctx context.Context, userId string, cType models.ContentType) error {
	path := r.getPath(userId, cType)
	if err := r.mc.RemoveObject(ctx, r.bucket, path, minio.RemoveObjectOptions{}); err != nil {
		return models.ErrS3Operation
	}

	return nil
}
