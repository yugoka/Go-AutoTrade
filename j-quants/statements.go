package jquants

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Statement は /fins/statements の1レコード分の財務情報を表す構造体
type Statement struct {
	DisclosedDate                              string `json:"DisclosedDate"`
	DisclosedTime                              string `json:"DisclosedTime"`
	LocalCode                                  string `json:"LocalCode"`
	DisclosureNumber                           string `json:"DisclosureNumber"`
	TypeOfDocument                             string `json:"TypeOfDocument"`
	TypeOfCurrentPeriod                        string `json:"TypeOfCurrentPeriod"`
	CurrentPeriodStartDate                     string `json:"CurrentPeriodStartDate"`
	CurrentPeriodEndDate                       string `json:"CurrentPeriodEndDate"`
	CurrentFiscalYearStartDate                 string `json:"CurrentFiscalYearStartDate"`
	CurrentFiscalYearEndDate                   string `json:"CurrentFiscalYearEndDate"`
	NextFiscalYearStartDate                      string `json:"NextFiscalYearStartDate"`
	NextFiscalYearEndDate                        string `json:"NextFiscalYearEndDate"`
	NetSales                                    string `json:"NetSales"`
	OperatingProfit                             string `json:"OperatingProfit"`
	OrdinaryProfit                              string `json:"OrdinaryProfit"`
	Profit                                      string `json:"Profit"`
	EarningsPerShare                            string `json:"EarningsPerShare"`
	DilutedEarningsPerShare                     string `json:"DilutedEarningsPerShare"`
	TotalAssets                                 string `json:"TotalAssets"`
	Equity                                      string `json:"Equity"`
	EquityToAssetRatio                          string `json:"EquityToAssetRatio"`
	BookValuePerShare                           string `json:"BookValuePerShare"`
	CashFlowsFromOperatingActivities            string `json:"CashFlowsFromOperatingActivities"`
	CashFlowsFromInvestingActivities            string `json:"CashFlowsFromInvestingActivities"`
	CashFlowsFromFinancingActivities            string `json:"CashFlowsFromFinancingActivities"`
	CashAndEquivalents                          string `json:"CashAndEquivalents"`
	ResultDividendPerShare1stQuarter            string `json:"ResultDividendPerShare1stQuarter"`
	ResultDividendPerShare2ndQuarter            string `json:"ResultDividendPerShare2ndQuarter"`
	ResultDividendPerShare3rdQuarter            string `json:"ResultDividendPerShare3rdQuarter"`
	ResultDividendPerShareFiscalYearEnd         string `json:"ResultDividendPerShareFiscalYearEnd"`
	ResultDividendPerShareAnnual                string `json:"ResultDividendPerShareAnnual"`
	DistributionsPerUnit                        string `json:"DistributionsPerUnit(REIT)"`
	ResultTotalDividendPaidAnnual               string `json:"ResultTotalDividendPaidAnnual"`
	ResultPayoutRatioAnnual                     string `json:"ResultPayoutRatioAnnual"`
	ForecastDividendPerShare1stQuarter          string `json:"ForecastDividendPerShare1stQuarter"`
	ForecastDividendPerShare2ndQuarter          string `json:"ForecastDividendPerShare2ndQuarter"`
	ForecastDividendPerShare3rdQuarter          string `json:"ForecastDividendPerShare3rdQuarter"`
	ForecastDividendPerShareFiscalYearEnd       string `json:"ForecastDividendPerShareFiscalYearEnd"`
	ForecastDividendPerShareAnnual              string `json:"ForecastDividendPerShareAnnual"`
	ForecastDistributionsPerUnit                string `json:"ForecastDistributionsPerUnit(REIT)"`
	ForecastTotalDividendPaidAnnual             string `json:"ForecastTotalDividendPaidAnnual"`
	ForecastPayoutRatioAnnual                   string `json:"ForecastPayoutRatioAnnual"`
	NextYearForecastDividendPerShare1stQuarter    string `json:"NextYearForecastDividendPerShare1stQuarter"`
	NextYearForecastDividendPerShare2ndQuarter    string `json:"NextYearForecastDividendPerShare2ndQuarter"`
	NextYearForecastDividendPerShare3rdQuarter    string `json:"NextYearForecastDividendPerShare3rdQuarter"`
	NextYearForecastDividendPerShareFiscalYearEnd string `json:"NextYearForecastDividendPerShareFiscalYearEnd"`
	NextYearForecastDividendPerShareAnnual        string `json:"NextYearForecastDividendPerShareAnnual"`
	NextYearForecastDistributionsPerUnit          string `json:"NextYearForecastDistributionsPerUnit(REIT)"`
	NextYearForecastPayoutRatioAnnual             string `json:"NextYearForecastPayoutRatioAnnual"`
	ForecastNetSales2ndQuarter                    string `json:"ForecastNetSales2ndQuarter"`
	ForecastOperatingProfit2ndQuarter             string `json:"ForecastOperatingProfit2ndQuarter"`
	ForecastOrdinaryProfit2ndQuarter              string `json:"ForecastOrdinaryProfit2ndQuarter"`
	ForecastProfit2ndQuarter                      string `json:"ForecastProfit2ndQuarter"`
	ForecastEarningsPerShare2ndQuarter            string `json:"ForecastEarningsPerShare2ndQuarter"`
	NextYearForecastNetSales2ndQuarter            string `json:"NextYearForecastNetSales2ndQuarter"`
	NextYearForecastOperatingProfit2ndQuarter     string `json:"NextYearForecastOperatingProfit2ndQuarter"`
	NextYearForecastOrdinaryProfit2ndQuarter      string `json:"NextYearForecastOrdinaryProfit2ndQuarter"`
	NextYearForecastProfit2ndQuarter              string `json:"NextYearForecastProfit2ndQuarter"`
	NextYearForecastEarningsPerShare2ndQuarter      string `json:"NextYearForecastEarningsPerShare2ndQuarter"`
	ForecastNetSales                            string `json:"ForecastNetSales"`
	ForecastOperatingProfit                     string `json:"ForecastOperatingProfit"`
	ForecastOrdinaryProfit                      string `json:"ForecastOrdinaryProfit"`
	ForecastProfit                              string `json:"ForecastProfit"`
	ForecastEarningsPerShare                    string `json:"ForecastEarningsPerShare"`
	NextYearForecastNetSales                      string `json:"NextYearForecastNetSales"`
	NextYearForecastOperatingProfit               string `json:"NextYearForecastOperatingProfit"`
	NextYearForecastOrdinaryProfit                string `json:"NextYearForecastOrdinaryProfit"`
	NextYearForecastProfit                        string `json:"NextYearForecastProfit"`
	NextYearForecastEarningsPerShare              string `json:"NextYearForecastEarningsPerShare"`
	MaterialChangesInSubsidiaries                 string `json:"MaterialChangesInSubsidiaries"`
	SignificantChangesInTheScopeOfConsolidation   string `json:"SignificantChangesInTheScopeOfConsolidation"`
	ChangesBasedOnRevisionsOfAccountingStandard     string `json:"ChangesBasedOnRevisionsOfAccountingStandard"`
	ChangesOtherThanOnesBasedOnRevisionsOfAccountingStandard string `json:"ChangesOtherThanOnesBasedOnRevisionsOfAccountingStandard"`
	ChangesInAccountingEstimates                  string `json:"ChangesInAccountingEstimates"`
	RetrospectiveRestatement                      string `json:"RetrospectiveRestatement"`
	NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock string `json:"NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock"`
	NumberOfTreasuryStockAtTheEndOfFiscalYear     string `json:"NumberOfTreasuryStockAtTheEndOfFiscalYear"`
	AverageNumberOfShares                         string `json:"AverageNumberOfShares"`
	NonConsolidatedNetSales                       string `json:"NonConsolidatedNetSales"`
	NonConsolidatedOperatingProfit                string `json:"NonConsolidatedOperatingProfit"`
	NonConsolidatedOrdinaryProfit                 string `json:"NonConsolidatedOrdinaryProfit"`
	NonConsolidatedProfit                         string `json:"NonConsolidatedProfit"`
	NonConsolidatedEarningsPerShare               string `json:"NonConsolidatedEarningsPerShare"`
	NonConsolidatedTotalAssets                    string `json:"NonConsolidatedTotalAssets"`
	NonConsolidatedEquity                         string `json:"NonConsolidatedEquity"`
	NonConsolidatedEquityToAssetRatio             string `json:"NonConsolidatedEquityToAssetRatio"`
	NonConsolidatedBookValuePerShare              string `json:"NonConsolidatedBookValuePerShare"`
	ForecastNonConsolidatedNetSales2ndQuarter     string `json:"ForecastNonConsolidatedNetSales2ndQuarter"`
	ForecastNonConsolidatedOperatingProfit2ndQuarter string `json:"ForecastNonConsolidatedOperatingProfit2ndQuarter"`
	ForecastNonConsolidatedOrdinaryProfit2ndQuarter  string `json:"ForecastNonConsolidatedOrdinaryProfit2ndQuarter"`
	ForecastNonConsolidatedProfit2ndQuarter         string `json:"ForecastNonConsolidatedProfit2ndQuarter"`
	ForecastNonConsolidatedEarningsPerShare2ndQuarter string `json:"ForecastNonConsolidatedEarningsPerShare2ndQuarter"`
	NextYearForecastNonConsolidatedNetSales2ndQuarter string `json:"NextYearForecastNonConsolidatedNetSales2ndQuarter"`
	NextYearForecastNonConsolidatedOperatingProfit2ndQuarter string `json:"NextYearForecastNonConsolidatedOperatingProfit2ndQuarter"`
	NextYearForecastNonConsolidatedOrdinaryProfit2ndQuarter  string `json:"NextYearForecastNonConsolidatedOrdinaryProfit2ndQuarter"`
	NextYearForecastNonConsolidatedProfit2ndQuarter         string `json:"NextYearForecastNonConsolidatedProfit2ndQuarter"`
	NextYearForecastNonConsolidatedEarningsPerShare2ndQuarter string `json:"NextYearForecastNonConsolidatedEarningsPerShare2ndQuarter"`
	ForecastNonConsolidatedNetSales               string `json:"ForecastNonConsolidatedNetSales"`
	ForecastNonConsolidatedOperatingProfit        string `json:"ForecastNonConsolidatedOperatingProfit"`
	ForecastNonConsolidatedOrdinaryProfit         string `json:"ForecastNonConsolidatedOrdinaryProfit"`
	ForecastNonConsolidatedProfit                 string `json:"ForecastNonConsolidatedProfit"`
	ForecastNonConsolidatedEarningsPerShare         string `json:"ForecastNonConsolidatedEarningsPerShare"`
	NextYearForecastNonConsolidatedNetSales         string `json:"NextYearForecastNonConsolidatedNetSales"`
	NextYearForecastNonConsolidatedOperatingProfit  string `json:"NextYearForecastNonConsolidatedOperatingProfit"`
	NextYearForecastNonConsolidatedOrdinaryProfit   string `json:"NextYearForecastNonConsolidatedOrdinaryProfit"`
	NextYearForecastNonConsolidatedProfit           string `json:"NextYearForecastNonConsolidatedProfit"`
	NextYearForecastNonConsolidatedEarningsPerShare   string `json:"NextYearForecastNonConsolidatedEarningsPerShare"`
}

// statementsResponse は /fins/statements のレスポンス全体を表す構造体
type statementsResponse struct {
	Statements    []Statement `json:"statements"`
	PaginationKey string      `json:"pagination_key"`
}

// GetStatementsParams は、/fins/statements 取得時のクエリパラメータを定義する
type GetStatementsParams struct {
	Code string // 例: "86970" または "8697"
	Date string // 例: "20230130" または "2023-01-30"
}

// GetStatements は /fins/statements をページネーション対応で全件取得し、[]Statement を返す
func (c *JQuantsClient) GetStatements(params GetStatementsParams) ([]Statement, error) {
	baseURL := "https://api.jquants.com/v1/fins/statements"
	q := url.Values{}
	if params.Code != "" {
		q.Set("code", params.Code)
	}
	if params.Date != "" {
		q.Set("date", params.Date)
	}

	// extract 関数で JSON のレスポンスから []Statement と pagination_key を抽出
	extractor := func(respBytes []byte) ([]Statement, string, error) {
		var r statementsResponse
		if err := json.Unmarshal(respBytes, &r); err != nil {
			return nil, "", fmt.Errorf("failed to unmarshal statements: %w", err)
		}
		return r.Statements, r.PaginationKey, nil
	}

	return DoPaginatedGet[Statement](c, baseURL, q, extractor)
}

