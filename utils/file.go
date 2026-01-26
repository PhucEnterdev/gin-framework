package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowedExts = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
}

var allowedMimeTypes = map[string]bool{
	"image/jpg":  true,
	"image/jpeg": true,
	"imgae/png":  true,
}

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {

	// check extension in filename
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExts[ext] {
		return "", errors.New("unsupported file extension")
	}
	// check size
	if fileHeader.Size > 10<<20 {
		return "", errors.New("file too large (max 5MB)")
	}

	// check type file
	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("cannot open file")
	}

	defer file.Close()
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", errors.New("cannot read file")
	}

	mimeType := http.DetectContentType(buffer)
	if !allowedMimeTypes[mimeType] {
		return "", fmt.Errorf("invalid mime type: %s", mimeType)
	}

	// change filename
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	// os.ModePerm = 0777 read write execute for all
	err = os.MkdirAll("./upload", os.ModePerm)
	if err != nil {
		return "", errors.New("cannot create upload folder")
	}

	// uploadDir "./upload" +filename "abc.png"
	savePath := filepath.Join(uploadDir, filename)
	if err := saveFile(fileHeader, savePath); err != nil {
		return "", err
	}
	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return nil
}
