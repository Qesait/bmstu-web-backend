package app

import (
	"fmt"
	"path/filepath"

	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (app *Application) uploadImage(c *gin.Context, image *multipart.FileHeader, UUID string) (*string, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(image.Filename)
	if extension != ".jpg" && extension != ".jpeg" {
		return nil, fmt.Errorf("разрешены только jpeg изображения")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.Minio.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("%s/%s/%s", app.config.Minio.Endpoint, app.config.Minio.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) deleteImage(c *gin.Context, UUID string) error {
	imageName := UUID + ".jpg"
	fmt.Println(imageName)
	err := app.minioClient.RemoveObject(c, app.config.Minio.BucketName, imageName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

// func (app *Application) getCustomer() string {
// 	return "5f58c307-a3f2-4b13-b888-c80ad08d5ed3"
// }

// func (app *Application) getModerator() *string {
// 	moderaorId := "796c70e1-5f27-4433-a415-95e7272effa5"
// 	return &moderaorId
// }
