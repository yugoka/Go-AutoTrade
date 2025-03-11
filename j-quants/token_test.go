package jquants

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

// RoundTripFunc は http.RoundTripper を簡単に実装するためのヘルパー
type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestSaveAndLoadTokens(t *testing.T) {
	// 一時ディレクトリに移動してからテスト実施
	tmpDir, err := os.MkdirTemp("", "jquants_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	origWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(origWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	token := &tokenData{
		RefreshToken:       "test_rt",
		RefreshTokenExpiry: now.Add(7 * 24 * time.Hour),
		IDToken:            "test_it",
		IDTokenExpiry:      now.Add(24 * time.Hour),
	}

	saveTokens(token)

	loaded, err := loadTokens()
	if err != nil {
		t.Fatalf("Failed to load tokens: %v", err)
	}
	if loaded.RefreshToken != token.RefreshToken || loaded.IDToken != token.IDToken {
		t.Error("Saved tokens and loaded tokens do not match")
	}
}

func TestIsExpiringOrExpired(t *testing.T) {
	// トークンの期限が未来(1分後)で、threshold が大きいと expiring と判定される
	future := time.Now().Add(1 * time.Minute)
	if !isExpiringOrExpired(future, 2*time.Minute) {
		t.Error("Expected token to be expiring with threshold 2m")
	}
	// threshold 0 なら expiring ではない
	if isExpiringOrExpired(future, 0) {
		t.Error("Expected token not to be expiring with threshold 0")
	}
}

func TestGetRefreshTokenByCredentials(t *testing.T) {
	// HTTP リクエストをフックするために http.DefaultTransport を差し替える
	origTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = origTransport }()

	http.DefaultTransport = RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != "https://api.jquants.com/v1/token/auth_user" {
			return nil, fmt.Errorf("unexpected URL: %s", req.URL.String())
		}
		respBody := `{"refreshToken": "test_rt"}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(respBody)),
			Header:     make(http.Header),
		}, nil
	})

	token, exp, err := getRefreshTokenByCredentials("dummy@mail", "dummy_pass")
	if err != nil {
		t.Fatalf("Error in getRefreshTokenByCredentials: %v", err)
	}
	if token != "test_rt" {
		t.Errorf("Expected refresh token 'test_rt', got: %s", token)
	}
	if time.Now().After(exp) {
		t.Error("Expected expiry time to be in the future")
	}
}

func TestGetIDTokenByRefreshToken(t *testing.T) {
	origTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = origTransport }()

	http.DefaultTransport = RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		expectedURL := "https://api.jquants.com/v1/token/auth_refresh?refreshtoken=test_rt"
		if req.URL.String() != expectedURL {
			return nil, fmt.Errorf("unexpected URL: %s", req.URL.String())
		}
		respBody := `{"idToken": "test_it"}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(respBody)),
			Header:     make(http.Header),
		}, nil
	})

	token, exp, err := getIDTokenByRefreshToken("test_rt")
	if err != nil {
		t.Fatalf("Error in getIDTokenByRefreshToken: %v", err)
	}
	if token != "test_it" {
		t.Errorf("Expected ID token 'test_it', got: %s", token)
	}
	if time.Now().After(exp) {
		t.Error("Expected expiry time to be in the future")
	}
}
