package minio

import (
	"context"
	"easy-gin-template/internal/config"
	"easy-gin-template/pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"mime/multipart"
)

type Client struct {
	client     *minio.Client
	bucketName string
	logger     *zap.Logger
}

var Minio *Client

func init() {
	Minio = NewMinioClient()
}

func NewMinioClient() *Client {
	minioClient, _ := minio.New(config.CFG.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.CFG.Minio.AccessKeyId, config.CFG.Minio.SecretAccessKey, ""),
		Secure: config.CFG.Minio.UseSSL,
	})

	return &Client{
		client:     minioClient,
		bucketName: config.CFG.Minio.BucketName,
		logger:     log.Logger,
	}
}

func (c *Client) UploadFile(file *multipart.FileHeader, category string) error {
	fileObj, err := file.Open()
	defer func(fileObj multipart.File) {
		err := fileObj.Close()
		if err != nil {
			c.logger.Error(err.Error())
		}
	}(fileObj)

	_, err = c.client.PutObject(context.TODO(), c.bucketName, category+"/"+file.Filename, fileObj, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}
	c.logger.Info("File uploaded successfully.")
	return nil
}

func (c *Client) DownloadFile(objectName string) (*minio.Object, error) {
	obj, err := c.client.GetObject(context.TODO(), c.bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	c.logger.Info("File downloaded successfully.")
	return obj, nil
}

func (c *Client) DeleteFile(objectName string) error {
	err := c.client.RemoveObject(context.TODO(), c.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}
	c.logger.Info("File deleted successfully.")
	return nil
}
