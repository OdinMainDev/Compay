package esfile

import (
	"mime/multipart"
	"path"
	"slices"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

func (that *ESFile) generateFirebaseUrl(imageName string, bucket string) (string, error) {
	imageStructure := imageStructure{
		ImageName: imageName,
		URL:       "https://storage.cloud.google.com/" + bucket + "/" + imageName,
	}

	_, _, err := that.client.Collection("image").Add(that.ctx, imageStructure)
	if err != nil {
		return "", err
	}
	return "https://storage.googleapis.com/" + bucket + "/" + imageName, nil
}

func GenerateImageName(handler *multipart.FileHeader) string {
	return uuid.New().String() + path.Ext(handler.Filename)
}

func IsValidMimeType(file multipart.File, allowedTypes []string) bool {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false
	}
	if !slices.Contains(allowedTypes, mime.Extension()) {
		return false
	}
	return true
}
