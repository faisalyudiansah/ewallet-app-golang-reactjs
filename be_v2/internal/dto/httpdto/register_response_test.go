package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/dto/authdto"
	"ewallet-server-v2/internal/model"
)

func TestConvertToRegisterResponse(t *testing.T) {
	type args struct {
		reg *authdto.RegisterDto
	}
	tests := []struct {
		name string
		args args
		want RegisterResponse
	}{
		{
			name: "success",
			args: args{
				reg: &authdto.RegisterDto{
					User: model.User{
						Name:  "lala",
						Email: "lala@mail.com",
					},
					Wallet: model.Wallet{
						WalletId:     1,
						WalletNumber: "0001",
					},
				},
			},
			want: RegisterResponse{
				Username:     "lala",
				Email:        "lala@mail.com",
				WalletId:     1,
				WalletNumber: "0001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToRegisterResponse(tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToRegisterResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
