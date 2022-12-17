package impl

import (
	"context"
	"fmt"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/fileutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"
)

func (s *invoiceService) GeneratePDF(ctx context.Context, req *dto.GetDetailsInvoicesRequest, downloadBase string) (*string, error) {
	data, err := s.repo.GetDetailInvoiceByID(ctx, req)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[GeneratePDF] Error on running service:", err.Error())
		}
		return nil, err
	}

	tempFileName := fmt.Sprintf("invoice-report-%s.pdf", randomutils.GenerateNRandomString(64))

	err = fileutils.CreatePDFFromHTMLFile(fmt.Sprintf("%s/invoices.html", downloadBase), tempFileName, data)
	if err != nil {
		return nil, err
	}

	return &tempFileName, nil
}
