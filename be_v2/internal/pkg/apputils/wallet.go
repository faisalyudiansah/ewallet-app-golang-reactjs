package apputils

import (
	"fmt"

	"ewallet-server-v2/internal/config"
)

type WalletNumberFormatter interface {
	Format(userId int64) string
}

type walletNumberFormatter struct {
	cfg config.AppConfig
}

func NewWalletNumberFormatter(cfg config.AppConfig) *walletNumberFormatter {
	return &walletNumberFormatter{
		cfg: cfg,
	}
}

func (f *walletNumberFormatter) Format(userId int64) string {
	return fmt.Sprintf("%s%010d", f.cfg.WalletPrefix, userId)
}
