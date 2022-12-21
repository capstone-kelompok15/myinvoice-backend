package handler

import (
	"reflect"
	"testing"
)

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		params *AuthHandlerParams
	}
	tests := []struct {
		name string
		args args
		want *authHandler
	}{
		{
			name: "Valid",
			args: args{
				params: &AuthHandlerParams{
					Service:   nil,
					Log:       nil,
					Validator: nil,
				},
			},
			want: NewAuthHandler(&AuthHandlerParams{
				Service:   nil,
				Log:       nil,
				Validator: nil,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
