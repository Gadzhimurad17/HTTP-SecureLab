package repo

import (
	"context"

	"github.com/Gagzhimurad17/NullPoint/web/backend/content/internal/domain/models"
)

type S3Repository interface {
	UploadFile(ctx context.Context, metadata *models.ContentMetadata) error
	Delete(ctx context.Context, userId string, cType models.ContentType) error
}
