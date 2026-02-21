package tpcgo

const VERSION = "0.0.7"

// NewSession creates a new session that can be used anywhere in the project assuming you put it in the root of the directory.
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
