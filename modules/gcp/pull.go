package gcp

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/mjarmoc/pull-push/modules/utils"
)

type Puller struct {
	client      *storage.Client
	object      *storage.ObjectHandle
	chunkNumber int16
	chunkSize   int64
}

func NewPuller(ctx context.Context) *Puller {
	client, err := storage.NewClient(ctx)

	if err != nil {
		panic(err)
	}
	return &Puller{client: client}
}

func (p *Puller) Pull(ctx context.Context, bucket, key string) int {
	p.object = p.client.Bucket(bucket).Object(key)
	attrs, err := p.object.Attrs(ctx)
	if err != nil {
	}
	p.chunkNumber = utils.CalculateChunkNumber(attrs.Size) - 1
	p.chunkSize = attrs.Size / int64(p.chunkNumber)
	return int(p.chunkNumber)
}

func (p *Puller) PullChunk(ctx context.Context, chunk int) *[]byte {
	start := int64(chunk) * p.chunkSize
	end := (int64(chunk) + 1) * p.chunkSize
	fmt.Printf("Start downloading chunk %d bytes from %d to %d\n", chunk, start, end)
	rc, err := p.object.NewRangeReader(ctx, start, end-start)
	if err != nil {

	}
	defer rc.Close()

	part, err := io.ReadAll(rc)
	if err != nil {

	}
	fmt.Printf("Completed downloading chunk %d\n", chunk)
	return &part
}
