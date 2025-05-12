package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mjarmoc/pull-push/modules/http"
	"github.com/mjarmoc/pull-push/modules/s3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.Flags().StringVarP(&bucket, "bucket", "b", "", "S3 Bucket name")
	s3Cmd.Flags().StringVarP(&url, "url", "u", "", "File url to download")
	s3Cmd.Flags().StringVarP(&file, "file", "f", "", "File path to upload")
}

var bucket string
var url string
var file string

var s3Cmd = &cobra.Command{
  Use:   "s3",
  Short: "Pull-Push to S3 Bucket",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Pull-Push to S3 Bucket")
	process(url, bucket, file)
  },
}



func process(url string, bucket string, file string) {
	fmt.Printf("Downloading file %s and uploading to S3 bucket %s with path %s", url, bucket, file)
	ctx := context.Background()
	start:= time.Now()
	puller := http.New()
	chunks := puller.Pull(ctx, url)
	pusher := s3.New(ctx)
	pusher.Push(ctx, bucket, file)
	var wg sync.WaitGroup
	for i:=0; i<chunks; i++ {
		wg.Add(1);
		go func() {
			defer wg.Done()
		 	bytes := puller.PullChunk(ctx, i)
			pusher.PushChunk(ctx, i, bytes)
		}()
	}
	wg.Wait()
	pusher.Complete(ctx)
	fmt.Println("Total execution time:", time.Since(start))
}
