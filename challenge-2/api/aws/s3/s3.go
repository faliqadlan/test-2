package s3

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
	"github.com/lithammer/shortuuid"
)

type TaskS3 struct {
	ses *session.Session
}

func New(ses *session.Session) *TaskS3 {
	return &TaskS3{
		ses: ses,
	}
}

func (t *TaskS3) UploadFileToS3(fileHeader multipart.FileHeader) (string, error) {

	var uid = shortuuid.New()

	var manager = s3manager.NewUploader(t.ses)
	var src, err = fileHeader.Open()
	if err != nil {
		log.Warn(err)
		return src.Close().Error(), err
	}
	defer src.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	var input = &s3manager.UploadInput{
		Bucket:       aws.String("karen-givi-bucket"),
		Key:          aws.String(uid),
		ACL:          aws.String("public-read-write"),
		Body:         bytes.NewReader(buffer),
		ContentType:  aws.String(http.DetectContentType(buffer)),
		StorageClass: aws.String("STANDARD"),
	}

	res, err := manager.Upload(input)

	if err != nil {
		log.Info(res)
		log.Error(err)
		return "there's some error", err

	}

	var url = "https://karen-givi-bucket.s3.ap-southeast-1.amazonaws.com/" + uid

	return url, nil
}

func (t *TaskS3) UpdateFileS3(name string, fileHeader multipart.FileHeader) string {

	src, err := fileHeader.Open()
	if err != nil {
		log.Info(err)
		return err.Error()
	}
	defer src.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	var svc = s3.New(t.ses)

	var input = &s3.PutObjectInput{
		Body:         bytes.NewReader(buffer),
		Bucket:       aws.String("karen-givi-bucket"),
		Key:          aws.String(name),
		ACL:          aws.String("public-read-write"),
		ContentType:  aws.String(http.DetectContentType(buffer)),
		StorageClass: aws.String("STANDARD"),
	}

	res, err := svc.PutObject(input)

	if err != nil {
		log.Info(res)
		log.Warn(err)
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Info(aerr.Error())
			}
			return err.Error()
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Info(err.Error())
			return err.Error()
		}
	}

	return "success"
}

func (t *TaskS3) DeleteFileS3(name string) string {
	var svc = s3.New(t.ses)
	// log.Info(name)
	var input = &s3.DeleteObjectInput{
		Bucket: aws.String("karen-givi-bucket"),
		Key:    aws.String(name),
	}

	var res, err = svc.DeleteObject(input)

	if err != nil {
		log.Info(res)
		log.Warn(err)
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Info(aerr.Error())
			}
			return err.Error()
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Info(err.Error())
			return err.Error()
		}
	}

	return "success"
}
