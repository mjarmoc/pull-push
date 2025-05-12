package s3

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Pusher struct {
	client *s3.Client
	bucket string
	key string
	uploadId string
	parts []types.CompletedPart
}

func New(ctx context.Context) (*Pusher) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println(err)
		panic("Couldn't load default configuration. Have you set up your AWS account?")
	}
	client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
		o.Region = "us-east-1"
		o.UsePathStyle = true
	})
	return &Pusher{client: client}
}

func (p *Pusher) Push(ctx context.Context, bucket string, key string) {
	p.bucket = bucket
	p.key = key
	upload, err := p.client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{Bucket: aws.String(p.bucket), Key: aws.String(p.key)})
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	p.uploadId = *upload.UploadId
}

func (p *Pusher) PushChunk(ctx context.Context, chunk int, filePart *[]byte) {
	fmt.Printf("Start pushing chunk %d\n", chunk)
	part, err := p.client.UploadPart(ctx, &s3.UploadPartInput{
		Bucket: aws.String(p.bucket), 
		Key: aws.String(p.key), 
		UploadId: aws.String(p.uploadId),
		PartNumber: aws.Int32(int32(chunk) + 1),
		Body: bytes.NewReader(*filePart),
	})
	if err != nil {
		panic(err)
	}
	completePart := types.CompletedPart{ETag: part.ETag, PartNumber: aws.Int32(int32(chunk) + 1)}
	p.parts = append(p.parts, completePart)
	fmt.Printf("Completed pushing chunk %d\n", chunk)
}

func (p *Pusher) Complete(ctx context.Context) {
	 slices.SortFunc(p.parts, func(a, b types.CompletedPart) int {return cmp.Compare(*a.PartNumber, *b.PartNumber)})
	_, err := p.client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(p.bucket), 
		Key: aws.String(p.key), 
		UploadId: aws.String(p.uploadId),
		MultipartUpload: &types.CompletedMultipartUpload{Parts: p.parts},
	})
	if err != nil {
		fmt.Print(err)
	}
}
