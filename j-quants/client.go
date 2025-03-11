package jquants

import (
	"Go-AutoTrade/config"
	"log"
)

// JQuantsClient は J-Quants API 利用のクライアント
type JQuantsClient struct {
	IDToken string
}

// New はトークン管理を行い、IDトークンをセットしたクライアントを返す
func New() (*JQuantsClient, error) {
	t, err := loadTokens()
	if err != nil {
		log.Println("[INFO] tokens.json not found. Creating new.")
		t = &tokenData{}
	}

	if t.RefreshToken == "" || isExpiringOrExpired(t.RefreshTokenExpiry, 0) {
		log.Println("[INFO] Refresh token invalid, acquiring new.")
		mail := config.GlobalConfig.JQuantsMailAddress
		pass := config.GlobalConfig.JQuantsPassword

		rt, rtExp, err := getRefreshTokenByCredentials(mail, pass)
		if err != nil {
			return nil, err
		}
		t.RefreshToken = rt
		t.RefreshTokenExpiry = rtExp
		log.Println("[INFO] Acquired new refresh token.")
		saveTokens(t)
	}

	if t.IDToken == "" || isExpiringOrExpired(t.IDTokenExpiry, 0) {
		log.Println("[INFO] ID token invalid, acquiring new.")
		it, itExp, err := getIDTokenByRefreshToken(t.RefreshToken)
		if err != nil {
			return nil, err
		}
		t.IDToken = it
		t.IDTokenExpiry = itExp
		log.Println("[INFO] Acquired new ID token.")
		saveTokens(t)
	}

	return &JQuantsClient{IDToken: t.IDToken}, nil
}
