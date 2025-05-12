package http

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mjarmoc/pull-push/modules/utils"
)

type Puller struct {
	client http.Client
	url string
	chunkNumber int16
	chunkSize int64
}

func New() (*Puller) {
	client := http.Client{}
	return &Puller{client: client}
}

func (p *Puller) Pull(ctx context.Context, url string) int{
	head, err:= p.client.Head(url)
	if err != nil {

	}
	p.chunkNumber = utils.CalculateChunkNumber(head.ContentLength) - 1
	// Check Accept-Ranges: bytes
	p.url = url
	p.chunkSize = head.ContentLength / int64(p.chunkNumber)
	return int(p.chunkNumber)
}

func(p *Puller) PullChunk(ctx context.Context, chunk int) (*[]byte) {
	request, err := http.NewRequest("GET", p.url, nil)
	if err != nil {

	}
	start := int64(chunk) * p.chunkSize
	end := (int64(chunk) + 1) * p.chunkSize
	rangeHeader := fmt.Sprintf("bytes=%d-%d", start, end)
	request.Header.Set("Range", rangeHeader)
	fmt.Printf("Start downloading chunk %d bytes from %d to %d\n", chunk, start, end)
	response, err2 := p.client.Do(request)
	if err2 != nil {

	}

	part, err3 := io.ReadAll(response.Body)
	if err3 !=nil {

	}
	fmt.Printf("Completed downloading chunk %d\n", chunk)
	return &part
}