package httpdto

import (
	"reflect"
	"testing"
	"time"

	"ewallet-server-v2/internal/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertToGetAllBoxesResponse(t *testing.T) {
	tests := []struct {
		name  string
		boxes []model.GameBox
		want  GetAllBoxesResponse
	}{
		{
			name: "success",
			boxes: []model.GameBox{
				{
					GameBoxId: 1,
					Amount:    decimal.NewFromInt(5000),
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: GetAllBoxesResponse{
				BoxNumbers: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToGetAllBoxesResponse(tt.boxes); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
