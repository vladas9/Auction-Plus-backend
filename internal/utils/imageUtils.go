package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func DecodeAndSaveImage(base64Image string, imageDir string, imageUUID string) error {
	if image, err := DecodeBase64Image(base64Image); err != nil {
		return err
	} else {
		if err := SaveImage(image, imageDir, imageUUID); err != nil {
			return err
		}
	}
	return nil
}

func DecodeBase64Image(base64Image string) ([]byte, error) {
	decodedImage, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(base64Image, "data:image/png;base64,"))
	if err != nil {
		return nil, fmt.Errorf("Faled to decode image: %v", err.Error())
	}
	return decodedImage, err
}

func SaveImage(image []byte, imageDir string, imageUUID string) error {
	imagePath := fmt.Sprintf("%s%s.png", imageDir, imageUUID)
	if err := os.WriteFile(imagePath, image, 0644); err != nil {
		return fmt.Errorf("Faled to save image: %v", err.Error())
	}
	return nil
}
