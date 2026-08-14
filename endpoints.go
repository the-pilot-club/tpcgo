package tpcgo

// Environment identifies which deployment of a service to target (e.g. production, beta).
type Environment string

const (
	EnvProduction Environment = "production"
	EnvBeta       Environment = "beta"
)

var (
	VATSIMDataFeedVersion = "3"
	VATSIMAPIVersion      = "2"

	EndpointFCP = func(Env string) string {
		if Env == "production" {
			return "https://flightcrew.thepilotclub.org/api"
		} else if Env == "beta" {
			return "https://flightcrew-beta.thepilotclub.org/api"
		} else {
			return ""
		}
	}

	coreAPIBaseURLs = map[Environment]string{
		EnvProduction: "https://api.thepilotclub.org",
		EnvBeta:       "https://api.ejsteiner.com",
	}
	EndpointCoreAPI = func(env Environment) string {
		return coreAPIBaseURLs[env]
	}

	EndpointVATSIMAPI  = "https://api.vatsim.net/"
	EndpointVATSIMData = "https://data.vatsim.net/"

	// FCP API Endpoints

	EndPointFCPALLFBOs = func(env string) string { return EndpointFCP(env) + "/fbos" }
	EndPointFCPFBO     = func(icao string, env string) string { return EndpointFCP(env) + "/fbo/" + icao }
	EndPointFCPSectors = func(env string) string { return EndpointFCP(env) + "/sectors" }

	/*
		FCP API  User Endpoints
	*/

	ENDPOINTFCPGetAllUsers    = func(env string) string { return EndpointFCP(env) + "/users/get" }
	ENDPOINTFCPGetAllUsersCID = func(env string) string { return EndpointFCP(env) + "/users/get/cid" }
	ENDPOINTFCPUser           = func(userId string, env string) string { return EndpointFCP(env) + "/users/find/" + userId }
	ENDPOINTFCPUserBirthdays  = func(env string) string { return EndpointFCP(env) + "/users/birthdays" }
	ENDPOINTFCPUserAdd        = func(env string) string { return EndpointFCP(env) + "/users/new" }
	ENDPOINTFCPUserDelete     = func(userId string, env string) string { return EndpointFCP(env) + "/users/find/" + userId + "/delete" }
	ENDPOINTFCPUserCallsign   = func(userId string, env string) string {
		return EndpointFCP(env) + "/users/find/" + userId + "/callsign"
	}
	ENDPOINTFCPUserAuditLogAdd = func(userId string, env string) string {
		return EndpointFCP(env) + "/users/find/" + userId + "/audit-logs/new"
	}

	/*
		CORE API Endpoints
	*/

	ENDPOINTCoreAPIAllSuggestions                = func(env Environment) string { return EndpointCoreAPI(env) + "/suggestions/all" }
	ENDPOINTCoreAPINewSuggestion                 = func(env Environment) string { return EndpointCoreAPI(env) + "/suggestions/new" }
	ENDPOINTCoreAPISuggestion                    = func(id string, env Environment) string { return EndpointCoreAPI(env) + "/suggestions/" + id }
	ENDPOINTCoreAPIAllQuizQuestions              = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/all" }
	ENDPOINTCoreAPINewQuizQuestion               = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/new" }
	ENDPOINTCoreAPINextQuizQuestion              = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/next" }
	ENDPOINTCoreAPICurrentQuizQuestion           = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/current" }
	EndPointCoreAPIQuizQuestionByID              = func(id string, env Environment) string { return EndpointCoreAPI(env) + "/quiz/" + id }
	ENDPOINTCoreAPIQuizSetQuestionForResponse    = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/responses/question/set" }
	ENDPOINTCoreAPIQuizSetUserResponse           = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/responses/user/record" }
	ENDPOINTCoreAPIQuizDeleteQuestionForResponse = func(id string, env Environment) string {
		return EndpointCoreAPI(env) + "/quiz/responses/question/" + id
	}
	ENDPOINTCoreAPIQuizGetResponses = func(id string, answer string, env Environment) string {
		return EndpointCoreAPI(env) + "/quiz/responses/user/responses/" + id + "?answer=" + answer
	}
	ENDPOINTCoreAPIResetUserResponses = func(env Environment) string { return EndpointCoreAPI(env) + "/quiz/responses/user/responses" }
	ENDPOINTCoreAPICheckUserResponse  = func(id string, env Environment) string {
		return EndpointCoreAPI(env) + "/quiz/responses/user/check/" + id
	}

	/*
		VATSIM API Endpoints
	*/

	EndpointVATSIMDataFeed  = EndpointVATSIMData + "v" + VATSIMDataFeedVersion + "/vatsim-data.json"
	EndpointVATSIMDiscordId = func(userId string) string {
		return EndpointVATSIMAPI + "v" + VATSIMAPIVersion + "/members/discord/" + userId
	}
	EndpointVATSIMUserHours = func(userId string) string {
		return EndpointVATSIMAPI + "v" + VATSIMAPIVersion + "/members/" + userId + "/stats"
	}
)
