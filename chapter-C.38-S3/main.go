package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_ACCESS_KEY_ID     string = "AKIA***********"
	AWS_SECRET_ACCESS_KEY string = "LReOy************"
	AWS_REGION            string = "ap-southeast-1"
)

func newS3Session() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func main() {
	s3Session, err := newS3Session()
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return
	}

	s3Client := s3.New(s3Session)

	bucketName := "adamstudio-new-bucket"
	fileName := "adamstudio.jpg"

	// --- create bucket
	err = createBucket(s3Client, bucketName)
	if err != nil {
		fmt.Printf("Couldn't create new bucket: %v", err)
		return
	}

	fmt.Println("New bucket successfully created")

	// --- list all buckets
	buckets, err := listBuckets(s3Client)
	if err != nil {
		fmt.Printf("Couldn't list buckets: %v", err)
		return
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
	}

	// --- upload file
	uploader := s3manager.NewUploader(s3Session)
	filePath := filepath.Join("upload", fileName)

	err = uploadFile(uploader, filePath, bucketName, fileName)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")

	// --- download file
	downloader := s3manager.NewDownloader(s3Session)
	err = downloadFile(downloader, bucketName, fileName)

	if err != nil {
		fmt.Printf("Couldn't download file: %v", err)
		return
	}

	fmt.Println("Successfully downloaded file")

	// --- delete file
	err = deleteFile(s3Client, bucketName, fileName)
	if err != nil {
		fmt.Printf("Couldn't delete file: %v", err)
		return
	}

	fmt.Println("Successfully delete file")

}

func listBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createBucket(client *s3.S3, bucketName string) error {
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func uploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}

func downloadFile(downloader *s3manager.Downloader, bucketName string, key string) error {
	file, err := os.Create(filepath.Join("download", key))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		},
	)

	return err
}

func deleteFile(client *s3.S3, bucketName string, fileName string) error {
	ctx := context.Background()
	iter := s3manager.NewDeleteListIterator(client, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(fileName),
	})

	if err := s3manager.NewBatchDeleteWithClient(client).Delete(ctx, iter); err != nil {
		return err
	}
	return nil
}
