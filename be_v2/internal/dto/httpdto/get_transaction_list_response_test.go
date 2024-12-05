package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/dto/appdto"
	"ewallet-server-v2/internal/dto/pagedto"
	"ewallet-server-v2/internal/model"
)

func TestConvertToGetTransactionListResponse(t *testing.T) {
	type args struct {
		transactionList *appdto.TransactionListDto
	}
	tests := []struct {
		name string
		args args
		want GetTransactionListResponse
	}{
		{
			name: "success",
			args: args{
				transactionList: &appdto.TransactionListDto{
					Entries: []model.Transaction{
						{
							TransactionId: 1,
						},
					},
					PageInfo: pagedto.PageInfoDto{
						Page:     1,
						Limit:    1,
						TotalRow: 1,
					},
				},
			},
			want: GetTransactionListResponse{
				Entries: []GetTransactionListEntryResponse{
					{
						TransactionId: 1,
					},
				},
				PageInfo: pagedto.PageInfoDto{
					Page:     1,
					Limit:    1,
					TotalRow: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToGetTransactionListResponse(tt.args.transactionList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToGetTransactionListResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
