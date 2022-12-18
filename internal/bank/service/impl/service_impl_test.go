package impl

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewBankService(t *testing.T) {
	type args struct {
		params *BankServiceParams
	}
	tests := []struct {
		name string
		args args
		want *bankService
	}{
		{
			name: "Valid",
			args: args{
				params: &BankServiceParams{
					Repo: nil,
					Log:  logrus.NewEntry(nil),
				},
			},
			want: NewBankService(&BankServiceParams{
				Repo: nil,
				Log:  logrus.NewEntry(nil),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBankService(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBankService() = %v, want %v", got, tt.want)
			}
		})
	}
}
