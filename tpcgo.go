package tpcgo

const VERSION = "0.0.7"

func NewSession(config SessionConfig) (s *Session, e error) {
	s = &Session{
		FCPSession: &FCPSession{
			ApiKeyHeader: "Authorization",
			ApiKey:       config.FCPKey,
			UserAgent:    "TPCGO (https://github.com/the-pilot-club/tpcgo, v" + VERSION + ")",
			Environment:  config.FCPEnv,
		},
		VATSIMSession: &VATSIMSession{
			ApiKeyHeader: "X-API-KEY",
			ApiKey:       config.VATSIMAPIKey,
			UserAgent:    "TPCGO (https://github.com/the-pilot-club/tpcgo, v" + VERSION + ")",
		},
		CoreAPISession: &CoreAPISession{
			ApiKeyHeader: "X-API-KEY",
			ApiKey:       config.CoreApiKey,
			UserAgent:    "TPCGO (https://github.com/the-pilot-club/tpcgo, v" + VERSION + ")",
		},
	}
	return
}
