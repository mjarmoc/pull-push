package gcp2s3

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mjarmoc/pull-push/cmd"
	"github.com/mjarmoc/pull-push/modules/gcp"
	"github.com/mjarmoc/pull-push/modules/s3"
	"github.com/spf13/cobra"
)

func Init() {
	cmd.RootCmd.AddCommand(gcp2s3Cmd)
	gcp2s3Cmd.Flags().StringVarP(&fromBucket, "from", "f", "", "S3 Bucket name")
	gcp2s3Cmd.Flags().StringVarP(&pullFile, "pull", "l", "", "File url to download")
	gcp2s3Cmd.Flags().StringVarP(&toBucket, "to", "t", "", "File path to upload")
	gcp2s3Cmd.Flags().StringVarP(&pushFile, "push", "s", "", "File path to upload")
}

var fromBucket, pullFile, toBucket, pushFile string

var gcp2s3Cmd = &cobra.Command{
	Use:   "gcp2s3",
	Short: "Pull From GCP and Push to S3 Bucket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		process(fromBucket, pullFile, toBucket, pushFile)
	},
}

func process(fromBucket, pullFile, toBucket, pushFile string) {
	fmt.Printf("From %s pulling file %s and to %s pushing %s", fromBucket, pullFile, toBucket, pushFile)
	ctx := context.Background()
	start := time.Now()
	puller := gcp.NewPuller(ctx)
	chunks := puller.Pull(ctx, fromBucket, pullFile)
	pusher := s3.NewPusher(ctx)
	pusher.Push(ctx, toBucket, pushFile)
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
