package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveProof(file *multipart.FileHeader, directory string, prefix string, ctx *gin.Context) (string, error) {

	extension := filepath.Ext(file.Filename)
	path := directory + prefix + "-proof-" + uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return file.Filename, err
	}

	return path, nil
}

func SaveFile(file *multipart.FileHeader, path string, ctx *gin.Context) (string, error) {

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return path, err
	}

	return path, nil
}
