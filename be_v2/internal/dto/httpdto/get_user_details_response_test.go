package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertToGetUserDetailResponse(t *testing.T) {
	type args struct {
		user   *model.User
		wallet *model.Wallet
	}
	tests := []struct {
		name string
		args args
		want GetUserDetailsResponse
	}{
		{
			name: "success",
			args: args{
				user: &model.User{
					UserId: 1,
					Name:   "lala",
					Email:  "lala@mail.com",
				},
				wallet: &model.Wallet{
					WalletId:     1,
					WalletNumber: "0001",
					Amount:       decimal.NewFromInt(5000),
				},
			},
			want: GetUserDetailsResponse{
				UserId:       1,
				Name:         "lala",
				Email:        "lala@mail.com",
				WalletId:     1,
				WalletNumber: "0001",
				Amount:       decimal.NewFromInt(5000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToGetUserDetailResponse(tt.args.user, tt.args.wallet); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
