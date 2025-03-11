package jquants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// 必ずgitignoreする
const tokenFilePath = "tokens.json"

type tokenData struct {
	RefreshToken       string    `json:"refresh_token"`
	RefreshTokenExpiry time.Time `json:"refresh_token_expiry"`
	IDToken            string    `json:"id_token"`
	IDTokenExpiry      time.Time `json:"id_token_expiry"`
}

func loadTokens() (*tokenData, error) {
	f, err := os.Open(tokenFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var t tokenData
	if err := json.NewDecoder(f).Decode(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func saveTokens(t *tokenData) {
	f, err := os.OpenFile(tokenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("[ERROR] Failed to open tokens.json: %v\n", err)
		return
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(t); err != nil {
		log.Printf("[ERROR] Failed to write tokens.json: %v\n", err)
	}
}

func isExpiringOrExpired(exp time.Time, threshold time.Duration) bool {
	return time.Now().Add(threshold).After(exp)
}

func getRefreshTokenByCredentials(mail, pass string) (string, time.Time, error) {
	apiURL := "https://api.jquants.com/v1/token/auth_user"
	body := map[string]string{
		"mailaddress": mail,
		"password":    pass,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return "", time.Time{}, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(b))
	if err != nil {
		return "", time.Time{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(resp.Body)
		return "", time.Time{}, fmt.Errorf("failed to get refresh token: status=%d, body=%s", resp.StatusCode, string(resBody))
	}

	var result struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", time.Time{}, err
	}

	log.Println("[INFO] Successfully obtained refresh token from auth_user.")
	return result.RefreshToken, time.Now().Add(7 * 24 * time.Hour), nil
}

func getIDTokenByRefreshToken(refreshToken string) (string, time.Time, error) {
	apiURL := "https://api.jquants.com/v1/token/auth_refresh?refreshtoken=" + url.QueryEscape(refreshToken)

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(resp.Body)
		return "", time.Time{}, fmt.Errorf("failed to get ID token: status=%d, body=%s", resp.StatusCode, string(resBody))
	}

	var result struct {
		IDToken string `json:"idToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", time.Time{}, err
	}

	log.Println("[INFO] Successfully obtained ID token from auth_refresh.")
	return result.IDToken, time.Now().Add(24 * time.Hour), nil
}