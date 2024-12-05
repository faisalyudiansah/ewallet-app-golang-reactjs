package httpdto

import (
	"reflect"
	"testing"

	"ewallet-server-v2/internal/dto/authdto"
)

func TestConvertToLoginResponse(t *testing.T) {
	type args struct {
		log *authdto.LoginDto
	}
	tests := []struct {
		name string
		args args
		want LoginResponse
	}{
		{
			name: "success",
			args: args{
				log: &authdto.LoginDto{
					Token: "token",
				},
			},
			want: LoginResponse{
				Token: "token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToLoginResponse(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToLoginResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
