package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSourceOfFundToString(t *testing.T) {
	tests := []struct {
		name           string
		sourceOfFundId int64
		want           string
	}{
		{
			name:           "success",
			sourceOfFundId: 1,
			want:           StringSourceOfFundBankTransfer,
		},
		{
			name:           "not registered",
			sourceOfFundId: -1,
			want:           "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSourceOfFundToString(tt.sourceOfFundId); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvertSourceOfFundToReadable(t *testing.T) {
	tests := []struct {
		name           string
		sourceOfFundId int64
		want           string
	}{
		{
			name:           "success",
			sourceOfFundId: 1,
			want:           ReadableSourceOfFundBankTransfer,
		},
		{
			name:           "not registered",
			sourceOfFundId: -1,
			want:           "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSourceOfFundToReadable(tt.sourceOfFundId); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConvertGoTimeLayoutToReadable(t *testing.T) {
	tests := []struct {
		name   string
		layout string
		want   string
	}{
		{
			name:   "success",
			layout: "2006-01-02",
			want:   "YYYY-MM-DD",
		},
		{
			name:   "not registered",
			layout: "-1",
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertGoTimeLayoutToReadable(tt.layout); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
