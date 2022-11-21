package config

import (
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

type Cloudinary struct {
	APIKey    string
	APISecret string
	CloudName string
}

var cld *cloudinary.Cloudinary

func initCloudinary(config *Cloudinary) error {
	var err error
	url := fmt.Sprintf("cloudinary://%s:%s@%s", config.APIKey, config.APISecret, config.CloudName)
	cld, err = cloudinary.NewFromURL(url)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
		return err
	}
	return nil
}

func GetCloudinaryConn(config *Cloudinary) (*cloudinary.Cloudinary, error) {
	if cld == nil {
		initCloudinary(config)
	}
	return cld, nil
}
