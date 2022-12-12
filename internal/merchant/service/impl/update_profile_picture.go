package impl

import (
	"context"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	cloudinary "github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/cloudinaryutils"
)

func (s *merchantService) UpdateProfilePicture(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
	imageURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
		Ctx:      ctx,
		Cld:      s.cloudinary,
		Filename: *newProfilePictureURL,
	})
	if err != nil {
		s.log.Warningln("[UpdateProfilePicture] Failed on upload the file:", err.Error())
		return nil, err
	}

	err = s.repo.UpdateProfilePicture(ctx, userID, imageURL)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return nil, err
		}
		s.log.Warningln("[UpdateProfilePicture] Failed on inserting to the database:", err.Error())
		return nil, err
	}

	return imageURL, nil
}
