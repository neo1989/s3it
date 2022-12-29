package cmd

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var BucketName string
var BucketRegion string
var FilePath string
var KeyPath string

func init() {
	uploadCmd.Flags().StringVarP(&BucketName, "name", "n", "", "BucketName")
	uploadCmd.Flags().StringVarP(&BucketRegion, "region", "r", "", "BucketRegion")
	uploadCmd.Flags().StringVarP(&KeyPath, "path", "p", "", "Path to store")
	uploadCmd.Flags().StringVarP(&FilePath, "file", "f", "", "Localfile path to upload")
	uploadCmd.MarkFlagRequired("name")
	uploadCmd.MarkFlagRequired("region")
	uploadCmd.MarkFlagRequired("file")
	uploadCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传",
	Run: func(cmd *cobra.Command, args []string) {
		upload()
	},
}

func md5sum() string {
	f, err := os.Open(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func upload() {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", BucketName, BucketRegion))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("S3IT_SECRETID"),
			SecretKey: os.Getenv("S3IT_SECRETKEY"),
		},
	})

	filename := fmt.Sprintf("%s%s", md5sum(), path.Ext(FilePath))

	key := fmt.Sprintf("%s/%s", KeyPath, filename)

	_, _, err := client.Object.Upload(
		context.Background(), key, FilePath, nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("AccessUrl: %s/%s\n", os.Getenv("S3IT_BASEURL"), key)

}
