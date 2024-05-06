package awss3

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Uploader(bucketName string,bucketKey string, filename string) (*manager.UploadOutput, error) {
	// code to upload file

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	
	if err != nil {
		log.Printf("Config error: %v", err)
		
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	fl, err := os.Open(filename);
	if err != nil{
		log.Println("File error:", err)
	}
	// copyParams := &s3.CopyObjectInput{
	// 	Bucket:     aws.String(os.Getenv("S3_BUCKET")),
	// 	CopySource: aws.String(fmt.Sprintf("%s/%s", os.Getenv("S3_BUCKET"), "sourceKey")),
	// 	Key:        aws.String("destinationKey"),
	// }
	// _, err = client.CopyObject(context.TODO(), copyParams)

	uploader := manager.NewUploader(client)
	
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("directory/"+ filename),
		Body:   fl,
	})
	if err != nil {
		log.Println("Upload error:", err)
		return nil, err
	}
	


	return result, nil


}
func Downloader() {
	// code to download file
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Println("error:", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	buffer := manager.NewWriteAtBuffer([]byte{})

	downloader := manager.NewDownloader(client)
	numBytes, err := downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String("my-bucket"), 
		Key:    aws.String("my-key"),
	}, )
	if err != nil {
		log.Fatal("error:", err)
		return
	}
	fmt.Println("Downloaded", numBytes, "bytes")

}

