package services

import (
	"github.com/flosch/pongo2/v6"
)

func init() {
	// Generator Filters
	pongo2.RegisterFilter("generate_string", FilterGenerateString)
	pongo2.RegisterFilter("generate_alphanumeric", FilterGenerateAlphaNumeric)
	pongo2.RegisterFilter("generate_hex", FilterGenerateHex)
	pongo2.RegisterFilter("generate_digits", FilterGenerateDigits)
	pongo2.RegisterFilter("generate_digits2", FilterGenerateDigits2) // Excludes '0'
	pongo2.RegisterFilter("generate_alpha", FilterGenerateAlpha)
	pongo2.RegisterFilter("generate_alpha_upper", FilterGenerateAlphaUpper)
	pongo2.RegisterFilter("generate_alpha_lower", FilterGenerateAlphaLower)

	// Timestamp Filters
	pongo2.RegisterFilter("timestamp_year", FilterTimestampYear)
	pongo2.RegisterFilter("timestamp_month", FilterTimestampMonth)
	pongo2.RegisterFilter("timestamp_day", FilterTimestampDay)
	pongo2.RegisterFilter("timestamp_hour", FilterTimestampHour)
	pongo2.RegisterFilter("timestamp_minute", FilterTimestampMinute)
	pongo2.RegisterFilter("timestamp_second", FilterTimestampSecond)
	pongo2.RegisterFilter("timestamp_datetime", FilterTimestampDatetime)
	pongo2.RegisterFilter("timestamp_random", FilterTimestampRandom)

	// File Filters
	pongo2.RegisterFilter("rand_file", FilterRandFile)

	// Local Data Filters
	pongo2.RegisterFilter("rand_local_data", FilterRandLocalData)
	pongo2.RegisterFilter("local_data", FilterLocalData)

	// Article Data Filters
	pongo2.RegisterFilter("fetch_random_article", FilterFetchRandomArticle)
	pongo2.RegisterFilter("rand_article_title", FilterRandArticleTitle)
}
