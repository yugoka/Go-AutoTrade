package jquants

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// DailyQuote は /prices/daily_quotes の1行分
type DailyQuote struct {
	Date              string  `json:"Date"`
	Code              string  `json:"Code"`
	Open              float64 `json:"Open"`
	High              float64 `json:"High"`
	Low               float64 `json:"Low"`
	Close             float64 `json:"Close"`
	Volume            float64 `json:"Volume"`
	TurnoverValue     float64 `json:"TurnoverValue"`
	AdjustmentOpen    float64 `json:"AdjustmentOpen"`
	AdjustmentHigh    float64 `json:"AdjustmentHigh"`
	AdjustmentLow     float64 `json:"AdjustmentLow"`
	AdjustmentClose   float64 `json:"AdjustmentClose"`
	AdjustmentVolume  float64 `json:"AdjustmentVolume"`
	// ... etc. 必要に応じて
}

// dailyQuotesResponse : JSON全体を受け取るための構造
type dailyQuotesResponse struct {
	DailyQuotes   []DailyQuote `json:"daily_quotes"`
	PaginationKey string       `json:"pagination_key"`
}

// GetDailyQuotesParams : クエリパラメータ
type GetDailyQuotesParams struct {
	Code string
	From string
	To   string
	Date string
}

// GetDailyQuotes は /prices/daily_quotes を全ページ取得し、[]DailyQuote を返す
func (c *JQuantsClient) GetDailyQuotes(params GetDailyQuotesParams) ([]DailyQuote, error) {
	baseURL := "https://api.jquants.com/v1/prices/daily_quotes"
	q := url.Values{}

	if params.Code != "" {
		q.Set("code", params.Code)
	}
	if params.From != "" {
		q.Set("from", params.From)
	}
	if params.To != "" {
		q.Set("to", params.To)
	}
	if params.Date != "" {
		q.Set("date", params.Date)
	}

	extractor := func(respBytes []byte) ([]DailyQuote, string, error) {
		var r dailyQuotesResponse
		if err := json.Unmarshal(respBytes, &r); err != nil {
			return nil, "", fmt.Errorf("failed to unmarshal daily_quotes: %w", err)
		}
		return r.DailyQuotes, r.PaginationKey, nil
	}

	// ここで pagination.go の共通関数を呼び出す
	return DoPaginatedGet[DailyQuote](c, baseURL, q, extractor)
}
