package handler

import (
	"reflect"
	"testing"
)

func TestNewCustomerHandler(t *testing.T) {
	type args struct {
		params *CustomerHandlerParams
	}
	tests := []struct {
		name string
		args args
		want *customerHandler
	}{
		{
			name: "Valid nil",
			args: args{
				params: &CustomerHandlerParams{
					Service:   nil,
					Log:       nil,
					Validator: nil,
				},
			},
			want: NewCustomerHandler(&CustomerHandlerParams{
				Service:   nil,
				Log:       nil,
				Validator: nil,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
