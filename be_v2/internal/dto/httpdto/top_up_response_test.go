package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/model"
)

func TestConvertToTopUpResponse(t *testing.T) {
	type args struct {
		transaction *model.Transaction
		wallet      *model.Wallet
	}
	tests := []struct {
		name string
		args args
		want TopUpResponse
	}{
		{
			name: "success",
			args: args{
				transaction: &model.Transaction{
					TransactionId: 1,
				},
				wallet: &model.Wallet{
					WalletId: 1,
				},
			},
			want: TopUpResponse{
				TransactionId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToTopUpResponse(tt.args.transaction, tt.args.wallet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToTopUpResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
