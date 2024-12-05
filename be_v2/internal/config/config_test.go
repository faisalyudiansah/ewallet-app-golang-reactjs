package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initAppConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  AppConfig
		fatal bool
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: AppConfig{
				Environment:  "dev",
				BCryptCost:   20,
				WalletPrefix: "0001",
			},
			fatal: false,
			mock: func() {
				os.Setenv("APP_ENVIRONMENT", "dev")
				os.Setenv("APP_BCRYPT_COST", "20")
				os.Setenv("APP_WALLET_PREFIX", "0001")
			},
			reset: func() {
				os.Unsetenv("APP_ENVIRONMENT")
				os.Unsetenv("APP_BCRYPT_COST")
				os.Unsetenv("APP_WALLET_PREFIX")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if !tt.fatal {
				if got := initAppConfig(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("initAppConfig() = %v, want %v", got, tt.want)
				}
			}
			tt.reset()
		})
	}
}

func Test_initHttpServerConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  HttpServerConfig
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: HttpServerConfig{
				Host:              "localhost",
				Port:              8080,
				GracePeriod:       5,
				MaxUploadFileSize: 5 * 1024,
			},
			mock: func() {
				os.Setenv("HTTP_SERVER_HOST", "localhost")
				os.Setenv("HTTP_SERVER_PORT", "8080")
				os.Setenv("HTTP_SERVER_GRACE_PERIOD", "5")
				os.Setenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB", "5")
			},
			reset: func() {
				os.Unsetenv("HTTP_SERVER_HOST")
				os.Unsetenv("HTTP_SERVER_PORT")
				os.Unsetenv("HTTP_SERVER_GRACE_PERIOD")
				os.Unsetenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := initHttpServerConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initHttpServerConfig() = %v, want %v", got, tt.want)
			}
			tt.reset()
		})
	}
}

func Test_initDbConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  DatabaseConfig
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: DatabaseConfig{
				Host:                  "localhost",
				Port:                  5432,
				DbName:                "lala",
				Username:              "admin",
				Password:              "admin",
				Sslmode:               "disable",
				MaxIdleConn:           10,
				MaxOpenConn:           100,
				MaxConnLifetimeMinute: 60,
			},
			mock: func() {
				os.Setenv("DB_HOST", "localhost")
				os.Setenv("DB_USER", "admin")
				os.Setenv("DB_PASSWORD", "admin")
				os.Setenv("DB_PORT", "5432")
				os.Setenv("DB_NAME", "lala")
				os.Setenv("DB_SSL_MODE", "disable")
				os.Setenv("DB_MAX_IDLE_CONN", "10")
				os.Setenv("DB_MAX_OPEN_CONN", "100")
				os.Setenv("DB_CONN_MAX_LIFETIME", "60")
			},
			reset: func() {
				os.Unsetenv("DB_HOST")
				os.Unsetenv("DB_USER")
				os.Unsetenv("DB_PASSWORD")
				os.Unsetenv("DB_PORT")
				os.Unsetenv("DB_NAME")
				os.Unsetenv("DB_SSL_MODE")
				os.Unsetenv("DB_MAX_IDLE_CONN")
				os.Unsetenv("DB_MAX_OPEN_CONN")
				os.Unsetenv("DB_CONN_MAX_LIFETIME")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := initDbConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initDbConfig() = %v, want %v", got, tt.want)
			}
			tt.reset()
		})
	}
}

func Test_initJwtConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  JwtConfig
		mock  func()
		reset func()
	}{
		{
			name: "sucecss",
			want: JwtConfig{
				Issuer:        "system",
				AllowedAlgs:   []string{"ALGS"},
				TokenDuration: 10,
				SecretKey:     "secret",
			},
			mock: func() {
				os.Setenv("JWT_ISSUER", "system")
				os.Setenv("JWT_SECRET_KEY", "secret")
				os.Setenv("JWT_ALLOWED_ALGS", "ALGS")
				os.Setenv("JWT_DURATION", "10")
			},
			reset: func() {
				os.Unsetenv("JWT_ISSUER")
				os.Unsetenv("JWT_SECRET_KEY")
				os.Unsetenv("JWT_ALLOWED_ALGS")
				os.Unsetenv("JWT_DURATION")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := initJwtConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initJwtConfig() = %v, want %v", got, tt.want)
			}
			tt.reset()
		})
	}
}

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  *Config
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: &Config{
				App: AppConfig{
					BCryptCost: 20,
				},
				HttpServer: HttpServerConfig{
					Port:              8080,
					GracePeriod:       5,
					MaxUploadFileSize: 5 * 1024,
				},
				Database: DatabaseConfig{
					Port:                  5432,
					MaxIdleConn:           10,
					MaxOpenConn:           100,
					MaxConnLifetimeMinute: 60,
				},
				Jwt: JwtConfig{
					AllowedAlgs:   []string{""},
					TokenDuration: 10,
				},
			},
			mock: func() {
				os.Setenv("APP_BCRYPT_COST", "20")
				os.Setenv("HTTP_SERVER_PORT", "8080")
				os.Setenv("HTTP_SERVER_GRACE_PERIOD", "5")
				os.Setenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB", "5")
				os.Setenv("DB_PORT", "5432")
				os.Setenv("DB_MAX_IDLE_CONN", "10")
				os.Setenv("DB_MAX_OPEN_CONN", "100")
				os.Setenv("DB_CONN_MAX_LIFETIME", "60")
				os.Setenv("JWT_DURATION", "10")
			},
			reset: func() {
				os.Unsetenv("APP_BCRYPT_COST")
				os.Unsetenv("HTTP_SERVER_PORT")
				os.Unsetenv("HTTP_SERVER_GRACE_PERIOD")
				os.Unsetenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB")
				os.Unsetenv("DB_PORT")
				os.Unsetenv("DB_MAX_IDLE_CONN")
				os.Unsetenv("DB_MAX_OPEN_CONN")
				os.Unsetenv("DB_CONN_MAX_LIFETIME")
				os.Unsetenv("JWT_DURATION")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exist := true
			_, err := os.Stat(".env")
			if os.IsNotExist(err) {
				os.Create(".env")
				exist = false
			}
			tt.mock()

			if got := InitConfig(); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}

			if !exist {
				os.Remove(".env")
			}
		})
	}
}
