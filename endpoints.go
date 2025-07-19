package tpcgo

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
	ENDPOINTCoreAPI    = "http://127.0.0.1:3000"
	EndpointVATSIMAPI  = "https://api.vatsim.net/"
	EndpointVATSIMData = "https://data.vatsim.net/"

	// FCP API Endpoints

	EndPointFCPALLFBOs = func(env string) string { return EndpointFCP(env) + "/fbos" }
	EndPointFCPFBO     = func(icao string, env string) string { return EndpointFCP(env) + "/fbo/" + icao }
	EndPointFCPSectors = func(env string) string { return EndpointFCP(env) + "/sectors" }

	/*
		FCP API  User Endpoints
	*/

	ENDPOINTFCPGetAllUsers   = func(env string) string { return EndpointFCP(env) + "/users/get" }
	ENDPOINTFCPUser          = func(userId string, env string) string { return EndpointFCP(env) + "/users/find/" + userId }
	ENDPOINTFCPUserBirthdays = func(env string) string { return EndpointFCP(env) + "/users/birthdays" }
	ENDPOINTFCPUserAdd       = func(env string) string { return EndpointFCP(env) + "/users/new" }
	ENDPOINTFCPUserDelete    = func(userId string, env string) string { return EndpointFCP(env) + "/users/find/" + userId + "/delete" }
	ENDPOINTFCPUserCallsign  = func(userId string, env string) string {
		return EndpointFCP(env) + "/users/find/" + userId + "/callsign"
	}
	ENDPOINTFCPUserAuditLogAdd = func(userId string, env string) string {
		return EndpointFCP(env) + "/users/find/" + userId + "/audit-logs/new"
	}

	/*
		CORE API Endpoints
	*/

	ENDPOINTCoreAPIAllSuggestions      = ENDPOINTCoreAPI + "/suggestions/all"
	ENDPOINTCoreAPINewSuggestion       = ENDPOINTCoreAPI + "/suggestions/new"
	ENDPOINTCoreAPISuggestion          = func(id string) string { return ENDPOINTCoreAPI + "/suggestions/" + id }
	ENDPOINTCoreAPIAllQuizQuestions    = ENDPOINTCoreAPI + "/quiz/all"
	ENDPOINTCoreAPINewQuizQuestion     = ENDPOINTCoreAPI + "/quiz/new"
	ENDPOINTCoreAPINextQuizQuestion    = ENDPOINTCoreAPI + "/quiz/next"
	ENDPOINTCoreAPICurrentQuizQuestion = ENDPOINTCoreAPI + "/quiz/current"
	EndPointCoreAPIQuizQuestionByID    = func(id string) string { return ENDPOINTCoreAPI + "/quiz/" + id }
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
