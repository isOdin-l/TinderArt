package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/isOdin-l/TinderArt/pkg/configs"
)

type RustFS struct {
	Client *s3.Client
}

func NewRustFS(cfg *configs.ConfigRustFS) *RustFS {
	s3Config := aws.Config{
		Region: cfg.RustFSRegion,
		EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: cfg.DSN(),
			}, nil
		}),
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			cfg.RustFSAccess_key,
			cfg.RustFSSecret_access_key,
			"",
		)),
	}

	return &RustFS{Client: s3.NewFromConfig(s3Config, func(o *s3.Options) {
		o.UsePathStyle = true
	})}
}

func (storage *RustFS) PutObject(ctx context.Context, bucket, key *string, body io.Reader) error {
	_, err := storage.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
		Body:   body,
	})
	return err
}

func (storage *RustFS) GetObject(ctx context.Context, bucket, key *string) ([]byte, error) {
	resp, errObj := storage.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if errObj != nil {
		return nil, errObj
	}
	defer resp.Body.Close()

	data, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}

	return data, nil
}
