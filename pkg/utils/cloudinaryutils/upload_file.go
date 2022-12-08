package cloudinary

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadFileParams struct {
	Ctx      context.Context
	Cld      *cloudinary.Cloudinary
	Filename string
}

func UploadFile(params UploadFileParams) (*string, error) {
	resp, err := params.Cld.Upload.Upload(params.Ctx, params.Filename, uploader.UploadParams{
		UseFilename:    api.Bool(true),
		UniqueFilename: api.Bool(true),
		ResourceType:   "auto",
	})
	if err != nil {
		log.Println("[ERROR] Error while uploading to the cloudinary:", err.Error())
		return nil, err
	}
	return &resp.SecureURL, nil
}
