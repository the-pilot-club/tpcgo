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
		VASTSIM API Endpoints
	*/

	EndpointVATSIMDataFeed  = EndpointVATSIMData + "v" + VATSIMDataFeedVersion + "/vatsim-data.json"
	EndpointVATSIMDiscordId = func(userId string) string {
		return EndpointVATSIMAPI + "v" + VATSIMAPIVersion + "/members/discord/" + userId
	}
	EndpointVATSIMUserHours = func(userId string) string {
		return EndpointVATSIMAPI + "v" + VATSIMAPIVersion + "/members/" + userId + "/stats"
	}
)
