package tpcgo

const VERSION = "0.0.5"

func NewFCPSession(ApiKey string) (f *FCPSession, e error) {
	f = &FCPSession{
		ApiKeyHeader: "Authorization",
		ApiKey:       ApiKey,
		UserAgent:    "TPCGO (https://github.com/the-pilot-club/tpcgo, v" + VERSION + ")",
	}
	return
}

func NewVATSIMSession(ApiKey string) (v *VATSIMSession, e error) {
	v = &VATSIMSession{
		ApiKeyHeader: "X-API_KEY",
		ApiKey:       ApiKey,
		UserAgent:    "TPCGO (https://github.com/the-pilot-club/tpcgo, v" + VERSION + ")",
	}
	return
}
