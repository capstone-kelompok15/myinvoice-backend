package handler

import (
	"reflect"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/sirupsen/logrus"
)

func TestNewBankHandler(t *testing.T) {
	validator, _ := validatorutils.New()
	log := logrus.NewEntry(logrus.New())

	type args struct {
		params *BankHandler
	}
	tests := []struct {
		name string
		args args
		want *bankHandler
	}{
		{
			name: "Init the bank handler",
			args: args{
				params: &BankHandler{
					Service:   nil,
					Log:       log,
					Validator: validator,
				},
			},
			want: &bankHandler{
				service:   nil,
				log:       log,
				validator: validator,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBankHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBankHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
