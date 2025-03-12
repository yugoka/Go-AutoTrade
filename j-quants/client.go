package jquants

import (
	"Go-AutoTrade/config"
	"log"
	"time"
)

// JQuantsClient は J-Quants API 利用のクライアントを表す
type JQuantsClient struct {
	IDToken        string
	IDTokenExpiry  time.Time
	RefreshToken   string
	RefreshExp     time.Time
}

// New はトークン管理を行い、IDトークンをセットしたクライアントを返す
func New() (*JQuantsClient, error) {
	t, err := loadTokens()
	if err != nil {
		log.Println("[INFO] tokens.json not found. Creating new.")
		t = &tokenData{}
	}

	c := &JQuantsClient{
		IDToken:       t.IDToken,
		IDTokenExpiry: t.IDTokenExpiry,
		RefreshToken:  t.RefreshToken,
		RefreshExp:    t.RefreshTokenExpiry,
	}

	// いったんensureToken() で必ずトークンが有効になるようにする
	if err := c.ensureToken(); err != nil {
		return nil, err
	}

	log.Println("[INFO] J-Quants Client init done")
	return c, nil
}

// ensureToken はIDトークンが期限切れであれば再取得する、
// あるいはRefreshトークンも期限切れであれば再発行するなどを担うメソッド
func (c *JQuantsClient) ensureToken() error {
	// RefreshToken がない or 期限切れの場合
	if c.RefreshToken == "" || isExpiringOrExpired(c.RefreshExp, 0) {
		log.Println("[INFO] Refresh token invalid, acquiring new.")
		mail := config.GlobalConfig.JQuantsMailAddress
		pass := config.GlobalConfig.JQuantsPassword

		rt, rtExp, err := getRefreshTokenByCredentials(mail, pass)
		if err != nil {
			return err
		}
		c.RefreshToken = rt
		c.RefreshExp = rtExp
		log.Println("[INFO] Acquired new refresh token.")
	}

	// IDToken がない or 期限切れの場合
	if c.IDToken == "" || isExpiringOrExpired(c.IDTokenExpiry, 0) {
		log.Println("[INFO] ID token invalid, acquiring new.")
		it, itExp, err := getIDTokenByRefreshToken(c.RefreshToken)
		if err != nil {
			return err
		}
		c.IDToken = it
		c.IDTokenExpiry = itExp
		log.Println("[INFO] Acquired new ID token.")
	}

	// 最新のトークン情報を保存しておく
	t := &tokenData{
		RefreshToken:       c.RefreshToken,
		RefreshTokenExpiry: c.RefreshExp,
		IDToken:            c.IDToken,
		IDTokenExpiry:      c.IDTokenExpiry,
	}
	saveTokens(t)

	return nil
}
