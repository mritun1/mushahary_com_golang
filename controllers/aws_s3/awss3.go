package aws_s3

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gofiber/fiber/v2"
)

func S3Integrate() (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)),
	)

	if err != nil {
		return nil, err
	}
	s3client := s3.NewFromConfig(cfg)
	return s3client, nil
}

func Uploading(c *fiber.Ctx, my_bucket string, folder string) (*manager.UploadOutput, error) {

	// Get the uploaded file from the request
	file, err := c.FormFile("photo_file")
	if err != nil {
		return nil, err
	}

	//UPLOADING PHOTOS
	// Open the file and create an io.Reader object
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer src.Close()

	s3client, err := S3Integrate()
	if err != nil {
		return nil, err
	}

	newFileName := strings.ReplaceAll(file.Filename, " ", "_")

	//UPLOADING IMAGE TO S3
	Uploader := manager.NewUploader(s3client)
	result, err := Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(my_bucket),
		Key:    aws.String(folder + newFileName),
		Body:   src,
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteS3(my_bucket string, file_loc string) error {
	// Create an S3 client using the loaded configuration
	s3client, err := S3Integrate()
	if err != nil {
		return err
	}

	// Create a DeleteObjectInput object with the bucket and object key
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(my_bucket),
		Key:    aws.String(file_loc),
	}

	// Call the DeleteObject function to delete the object from S3
	_, err = s3client.DeleteObject(context.Background(), input)

	return nil
}
