package config

import "fmt"

type S3Config struct {
	MinioEndpoint  string `env:"MINIO_ENDPOINT"`
	MinioAccesKey  string `env:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY"`
	MinioBucket    string `env:"MINIO_BUCKET"`
	MinioUseSSL    bool   `env:"MINIO_USE_SSL=false"`
}

func (mcfg *S3Config) MINIO_DSN() string {
	return fmt.Sprintf(
		"endpoint=%s access_key=%s secret_key=%s bucket=%s ssl=%t",
		mcfg.MinioEndpoint,
		mcfg.MinioAccesKey,
		mcfg.MinioSecretKey,
		mcfg.MinioBucket,
		mcfg.MinioUseSSL,
	)
}
