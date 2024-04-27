package utils

import (
	"context"
	"mime/multipart"

	"github.com/JerryJeager/will-be-there-backend/config"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}
	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
