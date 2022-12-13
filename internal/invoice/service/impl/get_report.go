package impl

import (
	"context"
	"database/sql"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
)

func (s *invoiceService) GetReport(ctx context.Context, params *dto.ReportParams) (*dto.ReportResponse, error) {
	var dateArr []dto.ReportDateStr
	now := time.Now()

	params.ReportDaysInt = mappingDays[params.DateFilter]
	for _, day := range params.ReportDaysInt {
		startDate := now.AddDate(0, 0, -day.StartDate)
		endDate := now.AddDate(0, 0, -day.EndDate)
		dateArr = append(dateArr, dto.ReportDateStr{
			StartDate: dateutils.TimeToDateStr(startDate) + " 00:00:00",
			EndDate:   dateutils.TimeToDateStr(endDate) + " 23:59:59",
		})
	}
	params.ReportDaysStr = dateArr

	report, err := s.repo.GetReport(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrRecordNotFound
		}
		s.log.Warningln("[GetReport] Error on running the repo:", err.Error())
		return nil, err
	}

	return report, nil
}

var mappingDays = map[string][]dto.ReportDateInt{
	"1 Week": {
		dto.ReportDateInt{
			StartDate: 6,
			EndDate:   6,
		},
		dto.ReportDateInt{
			StartDate: 5,
			EndDate:   5,
		},
		dto.ReportDateInt{
			StartDate: 4,
			EndDate:   4,
		},
		dto.ReportDateInt{
			StartDate: 3,
			EndDate:   3,
		},
		dto.ReportDateInt{
			StartDate: 2,
			EndDate:   2,
		},
		dto.ReportDateInt{
			StartDate: 1,
			EndDate:   1,
		},
		dto.ReportDateInt{
			StartDate: 0,
			EndDate:   0,
		},
	},
	"1 Month": {
		dto.ReportDateInt{
			StartDate: 29,
			EndDate:   25,
		},
		dto.ReportDateInt{
			StartDate: 24,
			EndDate:   20,
		},
		dto.ReportDateInt{
			StartDate: 19,
			EndDate:   16,
		},
		dto.ReportDateInt{
			StartDate: 15,
			EndDate:   12,
		},
		dto.ReportDateInt{
			StartDate: 11,
			EndDate:   8,
		},
		dto.ReportDateInt{
			StartDate: 7,
			EndDate:   4,
		},
		dto.ReportDateInt{
			StartDate: 3,
			EndDate:   0,
		},
	},
	"3 Month": {
		dto.ReportDateInt{
			StartDate: 89,
			EndDate:   77,
		},
		dto.ReportDateInt{
			StartDate: 76,
			EndDate:   64,
		},
		dto.ReportDateInt{
			StartDate: 63,
			EndDate:   51,
		},
		dto.ReportDateInt{
			StartDate: 50,
			EndDate:   38,
		},
		dto.ReportDateInt{
			StartDate: 37,
			EndDate:   25,
		},
		dto.ReportDateInt{
			StartDate: 24,
			EndDate:   12,
		},
		dto.ReportDateInt{
			StartDate: 11,
			EndDate:   0,
		},
	},
	"1 Year": {
		dto.ReportDateInt{
			StartDate: 364,
			EndDate:   310,
		},
		dto.ReportDateInt{
			StartDate: 309,
			EndDate:   258,
		},
		dto.ReportDateInt{
			StartDate: 257,
			EndDate:   206,
		},
		dto.ReportDateInt{
			StartDate: 205,
			EndDate:   154,
		},
		dto.ReportDateInt{
			StartDate: 153,
			EndDate:   102,
		},
		dto.ReportDateInt{
			StartDate: 101,
			EndDate:   52,
		},
		dto.ReportDateInt{
			StartDate: 51,
			EndDate:   0,
		},
	},
}
