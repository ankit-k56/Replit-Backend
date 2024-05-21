package awss3

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	defer fl.Close();
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
		Bucket: aws.String("repelit-iam"), 
		Key:    aws.String("defined/"),
	}, )
	if err != nil {
		log.Fatal("error:", err)
		return
	}
	fmt.Println("Downloaded", numBytes, "bytes")

}


func CopyS3Folder (sourcePrefix string, destinationPrefix string)(error){
	// fmt.Println("A"+  sourcePrefix)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Println("error:", err)
		return err
	}

	ctx := context.Background();
	svc := s3.NewFromConfig(cfg)

	paginator := s3.NewListObjectsV2Paginator(svc,&s3.ListObjectsV2Input{
		Bucket: aws.String("repelit-iam"),
		Prefix: aws.String(sourcePrefix),
	})
	// fmt.Println(paginator.Prefix)

	for paginator.HasMorePages(){
		page, err := paginator.NextPage(ctx)
		
		if err != nil{
			log.Fatal("Error in page loading: ", err)
		}
		fmt.Println(len(page.Contents))
		for _, obj := range page.Contents{
			if(obj.Key == nil){
				return nil
			}
			
			sourceKey := *obj.Key
			// if !strings.HasPrefix(sourceKey, sourcePrefix) {
			// 	continue // Skip objects that don't match the prefix
			//   }
			fmt.Println(sourceKey)
		
			destinationKey := destinationPrefix  + sourceKey
			//Copying the object
			_, err := svc.CopyObject(ctx, &s3.CopyObjectInput{
				Bucket: aws.String("repelit-iam"),
				CopySource: aws.String("repelit-iam/"+sourceKey),
				Key: aws.String(destinationKey),

			})
			if err != nil{
				log.Fatal("Error in copying object: ", err)
				return err
			}
		}

	
	}
	return nil
}

func DownLoadFolder() error{
	
	cfg, err := config.LoadDefaultConfig(context.TODO(),config.WithRegion("us-east-1"))
	if err!= nil{
		log.Println("Error in config: ", err)
		return err
	}
	client := s3.NewFromConfig(cfg)
	manager := manager.NewDownloader(client)

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String("repelit-iam"),
		Prefix: aws.String("defined/HtmlCss/"),

	})

	for paginator.HasMorePages(){
		page, err := paginator.NextPage(context.TODO())
		if err != nil{
			log.Fatal("Error in page loading: ", err)
			return err
		}
		for _, obj := range page.Contents{
			if err := DowLoadFile(manager, "awsDownloads", "repelit-iam", *obj.Key); err != nil{
				log.Fatal("Error in downloading file: ", err)
				return err
			}
		}
	}
	return nil



}

func DowLoadFile(downloader *manager.Downloader, targetDirectory, bucketName, key string) error{
	// fmt.Println("A"+  sourcePrefix)
	file := filepath.Join(targetDirectory, key)
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		return err
	}

	// Set up the local file
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	// Download the file using the AWS SDK for Go
	fmt.Printf("Downloading s3://%s/%s to %s...\n", bucketName, key, file)
	_, err = downloader.Download(context.TODO(), fd, &s3.GetObjectInput{
		Bucket: aws.String("repelit-iam"),
		Key:    aws.String(key),
	})

	return err


}

