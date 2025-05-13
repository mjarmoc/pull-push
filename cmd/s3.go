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
	RootCmd.AddCommand(http2s3Cmd)
	http2s3Cmd.Flags().StringVarP(&toBucket, "to", "t", "", "S3 Bucket name")
	http2s3Cmd.Flags().StringVarP(&pullFile, "pull", "l", "", "File url to download")
	http2s3Cmd.Flags().StringVarP(&pushFile, "push", "s", "", "File path to upload")
}

var toBucket string
var pullFile string
var pushFile string

var http2s3Cmd = &cobra.Command{
	Use:   "http2s3",
	Short: "Pull From Http and Push to S3 Bucket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		process(pullFile, toBucket, pushFile)
	},
}

func process(url string, bucket string, file string) {
	fmt.Printf("Downloading file %s and uploading to S3 bucket %s with path %s", url, bucket, file)
	ctx := context.Background()
	start := time.Now()
	puller := http.NewPuller()
	chunks := puller.Pull(ctx, url)
	pusher := s3.NewPusher(ctx)
	pusher.Push(ctx, bucket, file)
	var wg sync.WaitGroup
	for i := 0; i < chunks; i++ {
		wg.Add(1)
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
