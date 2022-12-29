package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func init() {
	rootCmd.AddCommand(bucketsCmd)
}

func buckets() {
	c := cos.NewClient(nil, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("S3IT_SECRETID"),
			SecretKey: os.Getenv("S3IT_SECRETKEY"),
		},
	})

	s, _, err := c.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}

	for _, b := range s.Buckets {
		fmt.Printf("%#v\n", b)
	}

}

var bucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "查询存储桶列表",
	Run: func(cmd *cobra.Command, args []string) {
		buckets()
	},
}
