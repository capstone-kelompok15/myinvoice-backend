package impl

import (
	"context"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	cloudinary "github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/cloudinaryutils"
)

func (s *invoiceService) UploadPayment(ctx context.Context, customerID int, invoiceID int, filePath string) error {
	err := s.repo.ValidateInvoiceID(ctx, customerID, invoiceID)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[UploadPayment] Failed on running validate repository:", err.Error())
		}
		return err
	}

	imageURL, err := cloudinary.UploadFile(cloudinary.UploadFileParams{
		Ctx:      ctx,
		Cld:      s.cloudinary,
		Filename: filePath,
	})
	if err != nil {
		s.log.Warningln("[UploadPayment] Failed on upload the file:", err.Error())
		return err
	}

	err = s.repo.UploadPayment(ctx, invoiceID, *imageURL)
	if err != nil {
		s.log.Warningln("[UploadPayment] Failed on inserting repository:", err.Error())
		return err
	}

	return err
}
