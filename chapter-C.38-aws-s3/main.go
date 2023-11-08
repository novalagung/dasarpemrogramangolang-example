package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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

func newSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			AWS_ACCESS_KEY_ID,
			AWS_SECRET_ACCESS_KEY,
			"",
		),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func main() {
	sess, err := newSession()
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return
	}

	s3Client := s3.New(sess)
	fmt.Println("S3 session & client initialized")

	bucketName := "adamstudio-new-bucket"

	// =============== create bucket ===============
	err = createBucket(s3Client, bucketName)
	if err != nil {
		fmt.Printf("Couldn't create new bucket: %v", err)
		return
	}

	fmt.Println("New bucket successfully created")

	// =============== list all buckets ===============
	buckets, err := listBuckets(s3Client)
	if err != nil {
		fmt.Printf("Couldn't list buckets: %v", err)
		return
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
	}

	// =============== upload file ===============
	uploader := s3manager.NewUploader(sess)

	fileName := "adamstudio.jpg"
	filePath := filepath.Join("upload", fileName)

	err = uploadFile(uploader, filePath, bucketName, fileName)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}

	fmt.Println("Successfully uploaded file!")

	// =============== list objects ===============
	// bucketName := "adamstudio-new-bucket"
	objects, err := listObjects(s3Client, bucketName)
	if err != nil {
		fmt.Printf("Couldn't list objects: %v", err)
		return
	}

	for _, object := range objects.Contents {
		fmt.Printf("Found object: %s, size: %d\n", *object.Key, *object.Size)
	}

	// =============== download file ===============
	downloader := s3manager.NewDownloader(sess)
	// fileName := "adamstudio.jpg"
	// bucketName := "adamstudio-new-bucket"
	downloadPath := filepath.Join("download", fileName)
	err = downloadFile(downloader, bucketName, fileName, downloadPath)
	if err != nil {
		fmt.Printf("Couldn't download file: %v", err)
		return
	}

	fmt.Println("Successfully downloaded file")

	// =============== delete file ===============
	// fileName := "adamstudio.jpg"
	// bucketName := "adamstudio-new-bucket"
	err = deleteFile(s3Client, bucketName, fileName)
	if err != nil {
		fmt.Printf("Couldn't delete file: %v", err)
		return
	}

	fmt.Println("Successfully delete file")

	// =============== presign url ===============
	// fileName := "adamstudio.jpg"
	// bucketName := "adamstudio-new-bucket"
	urlStr, err := presignUrl(s3Client, bucketName, fileName)
	if err != nil {
		fmt.Printf("Couldn't presign url: %v", err)
		return
	}

	fmt.Println("Presign url:", urlStr)
}

func createBucket(client *s3.S3, bucketName string) error {
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func listBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return res, nil
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

func listObjects(client *s3.S3, bucketName string) (*s3.ListObjectsV2Output, error) {
	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func downloadFile(downloader *s3manager.Downloader, bucketName string, key string, downloadPath string) error {
	file, err := os.Create(downloadPath)
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
	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	return err
}

func presignUrl(client *s3.S3, bucketName string, fileName string) (string, error) {
	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}
