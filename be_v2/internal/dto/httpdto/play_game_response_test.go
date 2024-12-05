package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertToPlayGameResponse(t *testing.T) {
	type args struct {
		gameAttempt *model.GameAttempt
	}
	tests := []struct {
		name string
		args args
		want PlayGameResponse
	}{
		{
			name: "success",
			args: args{
				gameAttempt: &model.GameAttempt{
					GameAttemptId: 1,
					Amount:        decimal.NewFromInt(5000),
				},
			},
			want: PlayGameResponse{
				GameAttemptId: 1,
				PrizeAmount:   decimal.NewFromInt(5000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToPlayGameResponse(tt.args.gameAttempt); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
