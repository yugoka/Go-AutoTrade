package jquants

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Helper function: Formats a numeric string as currency in JPY.
func formatCurrency(val string) string {
	if val == "" {
		return ""
	}
	return val + " JPY"
}

// Helper function: Formats a numeric string as "per share" text (e.g., EPS, dividend).
func formatPerShare(val string) string {
	if val == "" {
		return ""
	}
	return val + " JPY/share"
}

// Helper function: Converts a ratio (in decimal) to a percentage string.
func formatRatio(val string) string {
	if val == "" {
		return ""
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return val + " %"
	}
	return fmt.Sprintf("%.2f%%", f*100)
}

// Helper function: Formats share counts.
func formatShares(val string) string {
	if val == "" {
		return ""
	}
	return val + " shares"
}

func GenerateStatementsReport(statements []Statement) string {
	if len(statements) == 0 {
		return "[Error]No statement data found."
	}

	// Sort by disclosure date
	sort.Slice(statements, func(i, j int) bool {
		return statements[i].DisclosedDate < statements[j].DisclosedDate
	})

	var sb strings.Builder
	for i, st := range statements {
		sb.WriteString(fmt.Sprintf("Statement #%d\n", i+1))
		sb.WriteString(fmt.Sprintf("Disclosure date/time: %s %s\n", st.DisclosedDate, st.DisclosedTime))
		sb.WriteString(fmt.Sprintf("Document type: %s\n", st.TypeOfDocument))
		if st.TypeOfCurrentPeriod != "" {
			sb.WriteString(fmt.Sprintf("Fiscal period: %s (%s ~ %s)\n", st.TypeOfCurrentPeriod, st.CurrentPeriodStartDate, st.CurrentPeriodEndDate))
		}

		// Consolidated P&L indicators
		if st.NetSales != "" || st.OperatingProfit != "" || st.OrdinaryProfit != "" || st.Profit != "" {
			sb.WriteString("▼Consolidated P&L\n")
			if st.NetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales: %s\n", formatCurrency(st.NetSales)))
			}
			if st.OperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit: %s\n", formatCurrency(st.OperatingProfit)))
			}
			if st.OrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit: %s\n", formatCurrency(st.OrdinaryProfit)))
			}
			if st.Profit != "" {
				sb.WriteString(fmt.Sprintf("   Net income: %s\n", formatCurrency(st.Profit)))
			}
			if st.EarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS: %s\n", formatPerShare(st.EarningsPerShare)))
			}
			if st.DilutedEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   Diluted EPS: %s\n", formatPerShare(st.DilutedEarningsPerShare)))
			}
		}

		// Consolidated financial position
		if st.TotalAssets != "" || st.Equity != "" || st.EquityToAssetRatio != "" || st.BookValuePerShare != "" {
			sb.WriteString("▼Consolidated Financial Position\n")
			if st.TotalAssets != "" {
				sb.WriteString(fmt.Sprintf("   Total assets: %s\n", formatCurrency(st.TotalAssets)))
			}
			if st.Equity != "" {
				sb.WriteString(fmt.Sprintf("   Equity: %s\n", formatCurrency(st.Equity)))
			}
			if st.EquityToAssetRatio != "" {
				sb.WriteString(fmt.Sprintf("   Equity ratio: %s\n", formatRatio(st.EquityToAssetRatio)))
			}
			if st.BookValuePerShare != "" {
				sb.WriteString(fmt.Sprintf("   BPS: %s\n", formatPerShare(st.BookValuePerShare)))
			}
		}

		// Cash flows
		if st.CashFlowsFromOperatingActivities != "" ||
			st.CashFlowsFromInvestingActivities != "" ||
			st.CashFlowsFromFinancingActivities != "" ||
			st.CashAndEquivalents != "" {
			sb.WriteString("▼Cash Flows\n")
			if st.CashFlowsFromOperatingActivities != "" {
				sb.WriteString(fmt.Sprintf("   Operating CF: %s\n", formatCurrency(st.CashFlowsFromOperatingActivities)))
			}
			if st.CashFlowsFromInvestingActivities != "" {
				sb.WriteString(fmt.Sprintf("   Investing CF: %s\n", formatCurrency(st.CashFlowsFromInvestingActivities)))
			}
			if st.CashFlowsFromFinancingActivities != "" {
				sb.WriteString(fmt.Sprintf("   Financing CF: %s\n", formatCurrency(st.CashFlowsFromFinancingActivities)))
			}
			if st.CashAndEquivalents != "" {
				sb.WriteString(fmt.Sprintf("   Cash & equivalents: %s\n", formatCurrency(st.CashAndEquivalents)))
			}
		}

		// Dividend results
		if st.ResultDividendPerShare1stQuarter != "" ||
			st.ResultDividendPerShare2ndQuarter != "" ||
			st.ResultDividendPerShare3rdQuarter != "" ||
			st.ResultDividendPerShareFiscalYearEnd != "" ||
			st.ResultDividendPerShareAnnual != "" ||
			st.DistributionsPerUnit != "" ||
			st.ResultTotalDividendPaidAnnual != "" ||
			st.ResultPayoutRatioAnnual != "" {
			sb.WriteString("▼Dividend Results\n")
			if st.ResultDividendPerShare1stQuarter != "" {
				sb.WriteString(fmt.Sprintf("   1st quarter-end dividend: %s\n", formatPerShare(st.ResultDividendPerShare1stQuarter)))
			}
			if st.ResultDividendPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   2nd quarter-end dividend: %s\n", formatPerShare(st.ResultDividendPerShare2ndQuarter)))
			}
			if st.ResultDividendPerShare3rdQuarter != "" {
				sb.WriteString(fmt.Sprintf("   3rd quarter-end dividend: %s\n", formatPerShare(st.ResultDividendPerShare3rdQuarter)))
			}
			if st.ResultDividendPerShareFiscalYearEnd != "" {
				sb.WriteString(fmt.Sprintf("   Year-end dividend: %s\n", formatPerShare(st.ResultDividendPerShareFiscalYearEnd)))
			}
			if st.ResultDividendPerShareAnnual != "" {
				sb.WriteString(fmt.Sprintf("   Annual total dividend: %s\n", formatPerShare(st.ResultDividendPerShareAnnual)))
			}
			if st.DistributionsPerUnit != "" {
				// For REIT
				sb.WriteString(fmt.Sprintf("   Distribution per unit (REIT): %s JPY/unit\n", st.DistributionsPerUnit))
			}
			if st.ResultTotalDividendPaidAnnual != "" {
				sb.WriteString(fmt.Sprintf("   Total dividends paid: %s\n", formatCurrency(st.ResultTotalDividendPaidAnnual)))
			}
			if st.ResultPayoutRatioAnnual != "" {
				sb.WriteString(fmt.Sprintf("   Payout ratio: %s\n", formatRatio(st.ResultPayoutRatioAnnual)))
			}
		}

		// Dividend forecast
		if st.ForecastDividendPerShare1stQuarter != "" ||
			st.ForecastDividendPerShare2ndQuarter != "" ||
			st.ForecastDividendPerShare3rdQuarter != "" ||
			st.ForecastDividendPerShareFiscalYearEnd != "" ||
			st.ForecastDividendPerShareAnnual != "" ||
			st.ForecastPayoutRatioAnnual != "" ||
			st.ForecastDistributionsPerUnit != "" {
			sb.WriteString("▼Dividend Forecast\n")
			if st.ForecastDividendPerShare1stQuarter != "" {
				sb.WriteString(fmt.Sprintf("   1st quarter-end forecast: %s\n", formatPerShare(st.ForecastDividendPerShare1stQuarter)))
			}
			if st.ForecastDividendPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   2nd quarter-end forecast: %s\n", formatPerShare(st.ForecastDividendPerShare2ndQuarter)))
			}
			if st.ForecastDividendPerShare3rdQuarter != "" {
				sb.WriteString(fmt.Sprintf("   3rd quarter-end forecast: %s\n", formatPerShare(st.ForecastDividendPerShare3rdQuarter)))
			}
			if st.ForecastDividendPerShareFiscalYearEnd != "" {
				sb.WriteString(fmt.Sprintf("   Year-end forecast: %s\n", formatPerShare(st.ForecastDividendPerShareFiscalYearEnd)))
			}
			if st.ForecastDividendPerShareAnnual != "" {
				sb.WriteString(fmt.Sprintf("   Annual total forecast: %s\n", formatPerShare(st.ForecastDividendPerShareAnnual)))
			}
			if st.ForecastDistributionsPerUnit != "" {
				// For REIT
				sb.WriteString(fmt.Sprintf("   Distribution per unit forecast (REIT): %s JPY/unit\n", st.ForecastDistributionsPerUnit))
			}
			if st.ForecastPayoutRatioAnnual != "" {
				sb.WriteString(fmt.Sprintf("   (Forecast) payout ratio: %s\n", formatRatio(st.ForecastPayoutRatioAnnual)))
			}
		}

		// Earnings forecast (current FY)
		if st.ForecastNetSales != "" ||
			st.ForecastOperatingProfit != "" ||
			st.ForecastOrdinaryProfit != "" ||
			st.ForecastProfit != "" ||
			st.ForecastEarningsPerShare != "" ||
			st.ForecastNetSales2ndQuarter != "" ||
			st.ForecastOperatingProfit2ndQuarter != "" ||
			st.ForecastOrdinaryProfit2ndQuarter != "" ||
			st.ForecastProfit2ndQuarter != "" ||
			st.ForecastEarningsPerShare2ndQuarter != "" {
			sb.WriteString("▼Earnings Forecast (Current FY)\n")
			if st.ForecastNetSales2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (2Q forecast): %s\n", formatCurrency(st.ForecastNetSales2ndQuarter)))
			}
			if st.ForecastOperatingProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (2Q forecast): %s\n", formatCurrency(st.ForecastOperatingProfit2ndQuarter)))
			}
			if st.ForecastOrdinaryProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (2Q forecast): %s\n", formatCurrency(st.ForecastOrdinaryProfit2ndQuarter)))
			}
			if st.ForecastProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net income (2Q forecast): %s\n", formatCurrency(st.ForecastProfit2ndQuarter)))
			}
			if st.ForecastEarningsPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   EPS (2Q forecast): %s\n", formatPerShare(st.ForecastEarningsPerShare2ndQuarter)))
			}
			if st.ForecastNetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (full-year forecast): %s\n", formatCurrency(st.ForecastNetSales)))
			}
			if st.ForecastOperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (full-year): %s\n", formatCurrency(st.ForecastOperatingProfit)))
			}
			if st.ForecastOrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (full-year): %s\n", formatCurrency(st.ForecastOrdinaryProfit)))
			}
			if st.ForecastProfit != "" {
				sb.WriteString(fmt.Sprintf("   Net income (full-year): %s\n", formatCurrency(st.ForecastProfit)))
			}
			if st.ForecastEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS (full-year): %s\n", formatPerShare(st.ForecastEarningsPerShare)))
			}
		}

		// Earnings forecast (next FY)
		if st.NextYearForecastNetSales != "" ||
			st.NextYearForecastOperatingProfit != "" ||
			st.NextYearForecastOrdinaryProfit != "" ||
			st.NextYearForecastProfit != "" ||
			st.NextYearForecastEarningsPerShare != "" ||
			st.NextYearForecastNetSales2ndQuarter != "" ||
			st.NextYearForecastOperatingProfit2ndQuarter != "" ||
			st.NextYearForecastOrdinaryProfit2ndQuarter != "" ||
			st.NextYearForecastProfit2ndQuarter != "" ||
			st.NextYearForecastEarningsPerShare2ndQuarter != "" {
			sb.WriteString("▼Earnings Forecast (Next FY)\n")
			if st.NextYearForecastNetSales2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (2Q forecast): %s\n", formatCurrency(st.NextYearForecastNetSales2ndQuarter)))
			}
			if st.NextYearForecastOperatingProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (2Q forecast): %s\n", formatCurrency(st.NextYearForecastOperatingProfit2ndQuarter)))
			}
			if st.NextYearForecastOrdinaryProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (2Q forecast): %s\n", formatCurrency(st.NextYearForecastOrdinaryProfit2ndQuarter)))
			}
			if st.NextYearForecastProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net income (2Q forecast): %s\n", formatCurrency(st.NextYearForecastProfit2ndQuarter)))
			}
			if st.NextYearForecastEarningsPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   EPS (2Q forecast): %s\n", formatPerShare(st.NextYearForecastEarningsPerShare2ndQuarter)))
			}
			if st.NextYearForecastNetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (full-year forecast): %s\n", formatCurrency(st.NextYearForecastNetSales)))
			}
			if st.NextYearForecastOperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (full-year): %s\n", formatCurrency(st.NextYearForecastOperatingProfit)))
			}
			if st.NextYearForecastOrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (full-year): %s\n", formatCurrency(st.NextYearForecastOrdinaryProfit)))
			}
			if st.NextYearForecastProfit != "" {
				sb.WriteString(fmt.Sprintf("   Net income (full-year): %s\n", formatCurrency(st.NextYearForecastProfit)))
			}
			if st.NextYearForecastEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS (full-year): %s\n", formatPerShare(st.NextYearForecastEarningsPerShare)))
			}
		}

		// Non-consolidated results
		if st.NonConsolidatedNetSales != "" ||
			st.NonConsolidatedOperatingProfit != "" ||
			st.NonConsolidatedOrdinaryProfit != "" ||
			st.NonConsolidatedProfit != "" ||
			st.NonConsolidatedEarningsPerShare != "" {
			sb.WriteString("▼Non-Consolidated Results\n")
			if st.NonConsolidatedNetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (non-consolidated): %s\n", formatCurrency(st.NonConsolidatedNetSales)))
			}
			if st.NonConsolidatedOperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (non-consolidated): %s\n", formatCurrency(st.NonConsolidatedOperatingProfit)))
			}
			if st.NonConsolidatedOrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (non-consolidated): %s\n", formatCurrency(st.NonConsolidatedOrdinaryProfit)))
			}
			if st.NonConsolidatedProfit != "" {
				sb.WriteString(fmt.Sprintf("   Net income (non-consolidated): %s\n", formatCurrency(st.NonConsolidatedProfit)))
			}
			if st.NonConsolidatedEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS (non-consolidated): %s\n", formatPerShare(st.NonConsolidatedEarningsPerShare)))
			}
		}

		// Non-consolidated forecast (current FY)
		if st.ForecastNonConsolidatedNetSales != "" ||
			st.ForecastNonConsolidatedOperatingProfit != "" ||
			st.ForecastNonConsolidatedOrdinaryProfit != "" ||
			st.ForecastNonConsolidatedProfit != "" ||
			st.ForecastNonConsolidatedEarningsPerShare != "" ||
			st.ForecastNonConsolidatedNetSales2ndQuarter != "" ||
			st.ForecastNonConsolidatedOperatingProfit2ndQuarter != "" ||
			st.ForecastNonConsolidatedOrdinaryProfit2ndQuarter != "" ||
			st.ForecastNonConsolidatedProfit2ndQuarter != "" ||
			st.ForecastNonConsolidatedEarningsPerShare2ndQuarter != "" {
			sb.WriteString("▼Non-Consolidated Forecast (Current FY)\n")
			if st.ForecastNonConsolidatedNetSales2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (2Q forecast_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedNetSales2ndQuarter)))
			}
			if st.ForecastNonConsolidatedOperatingProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (2Q forecast_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedOperatingProfit2ndQuarter)))
			}
			if st.ForecastNonConsolidatedOrdinaryProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (2Q forecast_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedOrdinaryProfit2ndQuarter)))
			}
			if st.ForecastNonConsolidatedProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net income (2Q forecast_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedProfit2ndQuarter)))
			}
			if st.ForecastNonConsolidatedEarningsPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   EPS (2Q forecast_non-consolidated): %s\n", formatPerShare(st.ForecastNonConsolidatedEarningsPerShare2ndQuarter)))
			}
			if st.ForecastNonConsolidatedNetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (full-year forecast_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedNetSales)))
			}
			if st.ForecastNonConsolidatedOperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (full-year_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedOperatingProfit)))
			}
			if st.ForecastNonConsolidatedOrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (full-year_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedOrdinaryProfit)))
			}
			if st.ForecastNonConsolidatedProfit != "" {
				sb.WriteString(fmt.Sprintf("   Net income (full-year_non-consolidated): %s\n", formatCurrency(st.ForecastNonConsolidatedProfit)))
			}
			if st.ForecastNonConsolidatedEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS (full-year_non-consolidated): %s\n", formatPerShare(st.ForecastNonConsolidatedEarningsPerShare)))
			}
		}

		// Non-consolidated forecast (next FY)
		if st.NextYearForecastNonConsolidatedNetSales != "" ||
			st.NextYearForecastNonConsolidatedOperatingProfit != "" ||
			st.NextYearForecastNonConsolidatedOrdinaryProfit != "" ||
			st.NextYearForecastNonConsolidatedProfit != "" ||
			st.NextYearForecastNonConsolidatedEarningsPerShare != "" ||
			st.NextYearForecastNonConsolidatedNetSales2ndQuarter != "" ||
			st.NextYearForecastNonConsolidatedOperatingProfit2ndQuarter != "" ||
			st.NextYearForecastNonConsolidatedOrdinaryProfit2ndQuarter != "" ||
			st.NextYearForecastNonConsolidatedProfit2ndQuarter != "" ||
			st.NextYearForecastNonConsolidatedEarningsPerShare2ndQuarter != "" {
			sb.WriteString("▼Non-Consolidated Forecast (Next FY)\n")
			if st.NextYearForecastNonConsolidatedNetSales2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (2Q forecast_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedNetSales2ndQuarter)))
			}
			if st.NextYearForecastNonConsolidatedOperatingProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (2Q forecast_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedOperatingProfit2ndQuarter)))
			}
			if st.NextYearForecastNonConsolidatedOrdinaryProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (2Q forecast_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedOrdinaryProfit2ndQuarter)))
			}
			if st.NextYearForecastNonConsolidatedProfit2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   Net income (2Q forecast_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedProfit2ndQuarter)))
			}
			if st.NextYearForecastNonConsolidatedEarningsPerShare2ndQuarter != "" {
				sb.WriteString(fmt.Sprintf("   EPS (2Q forecast_non-consolidated): %s\n", formatPerShare(st.NextYearForecastNonConsolidatedEarningsPerShare2ndQuarter)))
			}
			if st.NextYearForecastNonConsolidatedNetSales != "" {
				sb.WriteString(fmt.Sprintf("   Net sales (full-year forecast_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedNetSales)))
			}
			if st.NextYearForecastNonConsolidatedOperatingProfit != "" {
				sb.WriteString(fmt.Sprintf("   Operating profit (full-year_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedOperatingProfit)))
			}
			if st.NextYearForecastNonConsolidatedOrdinaryProfit != "" {
				sb.WriteString(fmt.Sprintf("   Ordinary profit (full-year_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedOrdinaryProfit)))
			}
			if st.NextYearForecastNonConsolidatedProfit != "" {
				sb.WriteString(fmt.Sprintf("   Net income (full-year_non-consolidated): %s\n", formatCurrency(st.NextYearForecastNonConsolidatedProfit)))
			}
			if st.NextYearForecastNonConsolidatedEarningsPerShare != "" {
				sb.WriteString(fmt.Sprintf("   EPS (full-year_non-consolidated): %s\n", formatPerShare(st.NextYearForecastNonConsolidatedEarningsPerShare)))
			}
		}

		// Important accounting changes / notes
		if st.MaterialChangesInSubsidiaries == "true" ||
			st.SignificantChangesInTheScopeOfConsolidation == "true" ||
			st.ChangesBasedOnRevisionsOfAccountingStandard == "true" ||
			st.ChangesOtherThanOnesBasedOnRevisionsOfAccountingStandard == "true" ||
			st.ChangesInAccountingEstimates == "true" ||
			st.RetrospectiveRestatement == "true" {
			sb.WriteString("▼Important Accounting Changes / Notes\n")
			if st.MaterialChangesInSubsidiaries == "true" {
				sb.WriteString("   Significant changes in subsidiaries during the period\n")
			}
			if st.SignificantChangesInTheScopeOfConsolidation == "true" {
				sb.WriteString("   Significant changes in the scope of consolidation\n")
			}
			if st.ChangesBasedOnRevisionsOfAccountingStandard == "true" {
				sb.WriteString("   Changes in accounting policy due to revisions of accounting standards\n")
			}
			if st.ChangesOtherThanOnesBasedOnRevisionsOfAccountingStandard == "true" {
				sb.WriteString("   Changes in accounting policy other than those based on revisions\n")
			}
			if st.ChangesInAccountingEstimates == "true" {
				sb.WriteString("   Changes in accounting estimates\n")
			}
			if st.RetrospectiveRestatement == "true" {
				sb.WriteString("   Retrospective restatement\n")
			}
		}

		// Number of shares
		if st.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock != "" ||
			st.NumberOfTreasuryStockAtTheEndOfFiscalYear != "" ||
			st.AverageNumberOfShares != "" {
			sb.WriteString("▼Shares Outstanding\n")
			if st.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock != "" {
				sb.WriteString(fmt.Sprintf("   Issued shares at FY-end (incl. treasury): %s\n", formatShares(st.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock)))
			}
			if st.NumberOfTreasuryStockAtTheEndOfFiscalYear != "" {
				sb.WriteString(fmt.Sprintf("   Treasury shares at FY-end: %s\n", formatShares(st.NumberOfTreasuryStockAtTheEndOfFiscalYear)))
			}
			if st.AverageNumberOfShares != "" {
				sb.WriteString(fmt.Sprintf("   Average shares during the period: %s\n", formatShares(st.AverageNumberOfShares)))
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
