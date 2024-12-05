package httpdto

import (
	"reflect"
	"testing"
	"time"

	"ewallet-server-v2/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestConvertToResetPasswordResponse(t *testing.T) {
	type args struct {
		email                string
		resetPasswordAttempt *model.ResetPasswordAttempt
	}
	tests := []struct {
		name string
		args args
		want ResetPasswordResponse
	}{
		{
			name: "success",
			args: args{
				email: "lala@mail.com",
				resetPasswordAttempt: &model.ResetPasswordAttempt{
					Code:      "token",
					ExpiredAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: ResetPasswordResponse{
				Email:     "lala@mail.com",
				Code:      "token",
				ExpiredAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToResetPasswordResponse(tt.args.email, tt.args.resetPasswordAttempt); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
