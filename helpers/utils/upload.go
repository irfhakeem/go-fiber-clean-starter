package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
)

func Uploads(file multipart.FileHeader, path string) error {
	parts := strings.Split(path, "/")
	fileID := parts[1]
	dirPath := fmt.Sprintf("%s/%s", "uploads", parts[0])

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", dirPath, fileID)

	uploadedFile, err := file.Open()
	if err != nil {
		return err

	}
	defer uploadedFile.Close()

	filename := file.Filename
	extension := strings.ToLower(filepath.Ext(filename))

	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	var maxSize int64
	if extensions[extension] {
		maxSize = 2 * 1024 * 1024
	} else {
		return dto.ErrUnsupportedFileType
	}

	if file.Size > maxSize {
		return dto.ErrFileTooLarge
	}

	targetFile, err := os.Create(filePath)
	if err != nil {
		return err

	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return err

	}

	return nil
}
