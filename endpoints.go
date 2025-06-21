package tpcgo

var (
	VATSIMDataFeedVersion = "3"
	VATSIMAPIVersion      = "2"

	EndpointFCP        = "https://flightcrew.thepilotclub.org/api"
	EndpointVATSIMAPI  = "https://api.vatsim.net/"
	EndpointVATSIMData = "https://data.vatsim.net/"

	// FCP API Endpoints

	EndPointFCPALLFBOs = EndpointFCP + "/fbos"
	EndPointFCPFBO     = func(icao string) string { return EndpointFCP + "/fbo/" + icao }
	EndPointFCPSectors = EndpointFCP + "/sectors"

	/*
		FCP API  User Endpoints
	*/

	ENDPOINTFCPGetAllUsers     = EndpointFCP + "/users/get"
	ENDPOINTFCPUser            = func(userId string) string { return EndpointFCP + "/users/find/" + userId }
	ENDPOINTFCPUserBirthdays   = EndpointFCP + "/users/birthdays"
	ENDPOINTFCPUserAdd         = EndpointFCP + "/users/new"
	ENDPOINTFCPUserDelete      = func(userId string) string { return EndpointFCP + "/users/find/" + userId + "/delete" }
	ENDPOINTFCPUserCallsign    = func(userId string) string { return EndpointFCP + "/users/find/" + userId + "/callsign" }
	ENDPOINTFCPUserAuditLogAdd = func(userId string) string { return EndpointFCP + "/users/find/" + userId + "/audit-logs/new" }

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
