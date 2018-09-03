package services

import (
	"net/url"
	"time"

	"github.com/asxcandrew/herd/server/config"
	minio "github.com/minio/minio-go"
)

type storageStruct struct {
	client *minio.Client
	bucket string
	path   string
}

var _storage *storageStruct

func Storage() *storageStruct {
	if _storage == nil {
		minioClient, err := minio.New(
			config.C.StorageEndpoint,
			config.C.StorageAccessKey,
			config.C.StorageSecretKey,
			true,
		)

		if err != nil {
			panic(err)
		}

		exists, err := minioClient.BucketExists(config.C.StorageBucket)

		if err != nil || !exists {
			panic("Bucket does not exist")
		}

		_storage = &storageStruct{
			client: minioClient,
			bucket: config.C.StorageBucket,
			path:   config.C.StoragePath,
		}
	}
	return _storage
}

func (s *storageStruct) PresignedPutObject(objectName string) (*url.URL, error) {
	url, err := s.client.PresignedPutObject(s.bucket, s.preparedName(objectName), time.Minute*5)
	return url, err
}

func (s *storageStruct) RemoveObject(objectName string) error {
	err := s.client.RemoveObject(s.bucket, objectName)
	return err
}

func (s *storageStruct) GetUrlForObject(objectName string) (link *url.URL) {
	link, _ = s.client.PresignedGetObject(s.bucket, s.preparedName(objectName), time.Hour, url.Values{})
	return link
}

func (s *storageStruct) preparedName(objectName string) string {
	return s.path + "/" + objectName
}
