package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func generateFileName(extension string) string {
	hashed := md5.New()
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	hashed.Write(randomBytes)
	hash := hex.EncodeToString(hashed.Sum(nil))
	return fmt.Sprintf("%s%s", hash, extension)
}

func saveFile(file multipart.File, filename string, dstDir string) (string, error) {
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	outPath := filepath.Join(dstDir, filename)
	outFile, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return outPath, nil
}

func openFile(outPath string) (*os.File, error) {
	outFile, err := os.OpenFile(outPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return outFile, nil
}

func ConvertAndSaveImage(file *multipart.FileHeader) (string, string, string, error) {
	srcFile, err := file.Open()
	if err != nil {
		return "", "", "", err
	}
	defer srcFile.Close()

	ext := filepath.Ext(file.Filename)
	fileName := generateFileName("")
	uploadDir := "uploads"
	fileNameWithExt := fmt.Sprintf("%s%s", fileName, ext)
	outPath, err := saveFile(srcFile, fileNameWithExt, uploadDir)
	if err != nil {
		return "", "", "", err
	}

	outFile, err := openFile(outPath)
	if err != nil {
		return "", "", "", err
	}
	defer outFile.Close()

	return fileName, uploadDir, ext, err
}
