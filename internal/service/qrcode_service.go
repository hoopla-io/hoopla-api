package service

import (
	"encoding/base64"
	"fmt"
)

type QRCodeService struct{}

func NewQRCodeService() *QRCodeService {
	return &QRCodeService{}
}

func (q *QRCodeService) GenerateQRCode(data string) string {
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedData
}

func (q *QRCodeService) DecodeQRCode(encodedData string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode QR code data: %w", err)
	}
	return string(decodedData), nil
}
