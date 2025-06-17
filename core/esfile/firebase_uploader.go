package esfile

import (
	"compay/core/eslog"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

func (that *ESFile) saveImageToCDN(imageName string, file io.Reader) string {
	bucket := ""
	wc := that.GetStorage().Bucket(bucket).Object(imageName).NewWriter(that.ctx)
	if mime.TypeByExtension(filepath.Ext(imageName)) != "" {
		wc.ContentType = mime.TypeByExtension(filepath.Ext(imageName))
	}

	_, err := io.Copy(wc, file)
	if err != nil {
		eslog.LogError(err)
		return ""
	}

	if err := wc.Close(); err != nil {
		eslog.LogError(err)
		return ""
	}

	result, err := that.generateFirebaseUrl(imageName, bucket)
	if err != nil {
		eslog.LogError(err)
		return ""
	}

	return result
}

// ALIAS FUNCTIONS
func (that *ESFile) UploadFileByForm(handler *multipart.FileHeader) string {
	dstFile, err := handler.Open()
	if err != nil {
		eslog.LogError(err)
		return ""
	}

	defer dstFile.Close()
	return that.saveImageToCDN(GenerateImageName(handler), dstFile)
}

func (that *ESFile) UploadFileByUrl(url string) string {
	id := uuid.New()
	imageName := id.String() + path.Ext(url)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	return that.saveImageToCDN(imageName, resp.Body)
}
