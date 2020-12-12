package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// The session the S3 Uploader will use
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	region := os.Getenv("REGION")
	// directory := os.Getenv("DIRECTORY")

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:      aws.String(region),
	},
	)
	// Create a S3 client with additional configuration
	input := s3.ListBucketsInput{}
	svc := s3.New(sess, aws.NewConfig().WithRegion(region).WithEndpoint("eu-central-1.linodeobjects.com"))
	req, resp := svc.ListBucketsRequest(&input)

	err = req.Send()
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}
	// TODO: find out how to compress a directory
	// TODO: find out how to encrypt the compressed files
}
