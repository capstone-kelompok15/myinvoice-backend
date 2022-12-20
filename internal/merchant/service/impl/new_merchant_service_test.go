package impl

import (
	"reflect"
	"testing"
)

func TestNewMerchantService(t *testing.T) {
	type args struct {
		params *MerchantServiceParams
	}
	tests := []struct {
		name string
		args args
		want *merchantService
	}{
		{
			name: "Valid",
			args: args{&MerchantServiceParams{
				RepoNotif:  nil,
				Repo:       nil,
				Log:        nil,
				Config:     nil,
				Cloudinary: nil,
			}},
			want: NewMerchantService(&MerchantServiceParams{
				RepoNotif:  nil,
				Repo:       nil,
				Log:        nil,
				Config:     nil,
				Cloudinary: nil,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerchantService(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMerchantService() = %v, want %v", got, tt.want)
			}
		})
	}
}
