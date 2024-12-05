package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/model"
)

func TestConvertToTransferResponse(t *testing.T) {
	type args struct {
		transaction *model.Transaction
		walletFrom  *model.Wallet
		walletTo    *model.Wallet
	}
	tests := []struct {
		name string
		args args
		want TransferResponse
	}{
		{
			name: "success",
			args: args{
				transaction: &model.Transaction{
					TransactionId: 1,
				},
				walletFrom: &model.Wallet{
					WalletId: 1,
				},
				walletTo: &model.Wallet{
					WalletId: 1,
				},
			},
			want: TransferResponse{
				TransactionId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToTransferResponse(tt.args.transaction, tt.args.walletFrom, tt.args.walletTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToTransferResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
