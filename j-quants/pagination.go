package jquants

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// PageDataExtractor は、レスポンスJSONから (dataのスライス, pagination_key) を抽出するための関数型
type PageDataExtractor[T any] func(respBytes []byte) ([]T, string, error)

// DoPaginatedGet は、ページネーション付きの GET リクエストを行い、すべてのページを取得して []T を返す汎用関数。
func DoPaginatedGet[T any](
	c *JQuantsClient,    // トークン管理・認証のためのクライアント
	baseURL string,      // 例: "https://api.jquants.com/v1/prices/daily_quotes"
	params url.Values,   // クエリパラメータ
	extract PageDataExtractor[T], // JSONをどうパースして dataとpagination_keyを取り出すか
) ([]T, error) {

	var result []T
	var paginationKey string

	for {
		// 1. トークンが期限切れであれば更新
		if err := c.ensureToken(); err != nil {
			return nil, fmt.Errorf("failed to ensure token: %w", err)
		}

		// 2. pagination_key の指定
		if paginationKey != "" {
			params.Set("pagination_key", paginationKey)
		} else {
			params.Del("pagination_key")
		}

		fullURL := baseURL + "?" + params.Encode()
		log.Printf("[INFO] GET => %s", fullURL)

		// 3. HTTPリクエスト
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+c.IDToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to do request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("request failed: status=%d, body=%s", resp.StatusCode, string(body))
		}

		// 4. レスポンスを読み取り
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}

		// 5. コールバックで dataPart と nextKey を抽出
		dataPart, next, err := extract(respBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to extract page data: %w", err)
		}

		result = append(result, dataPart...)

		// 6. pagination_key が空なら終了
		if next == "" {
			break
		}
		paginationKey = next
	}

	return result, nil
}
