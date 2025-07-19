package tpcgo

type Session struct {
	FCPSession     *FCPSession
	CoreAPISession *CoreAPISession
	VATSIMSession  *VATSIMSession
}

type SessionConfig struct {
	FCPKey       string
	FCPEnv       string
	VATSIMAPIKey string
	CoreApiKey   string
}

type FCPSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
	Environment  string
}

type CoreAPISession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
}

type VATSIMSession struct {
	ApiKeyHeader string
	ApiKey       string
	UserAgent    string
}
