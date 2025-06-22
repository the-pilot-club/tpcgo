package tpcgo

type FCPSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
	Environment  string
}

type VATSIMSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
}
