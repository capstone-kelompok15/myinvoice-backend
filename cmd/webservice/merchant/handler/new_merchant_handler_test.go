package handler

import (
	"reflect"
	"testing"
)

func TestNewMerchantHandler(t *testing.T) {
	type args struct {
		params *MerchantHandler
	}
	tests := []struct {
		name string
		args args
		want *merchantHandler
	}{
		{
			name: "Valid",
			args: args{
				params: &MerchantHandler{
					Service:   nil,
					Log:       nil,
					Validator: nil,
				},
			},
			want: NewMerchantHandler(
				&MerchantHandler{
					Service:   nil,
					Log:       nil,
					Validator: nil,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerchantHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMerchantHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
