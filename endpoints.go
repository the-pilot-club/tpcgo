package tpcgo

var (
	VATSIMDataFeedVersion = "3"

	EndpointFCP        = "https://flightcrew.thepilotclub.org/api"
	EndpointVATSIMData = "https://data.vatsim.net/"

	// FCP API Endpoints

	ENDPOINTFCPUser            = func(userId string) string { return EndpointFCP + "/users/find/" + userId }
	ENDPOINTFCPUserCallsign    = func(userId string) string { return EndpointFCP + "/users/find/" + userId + "/callsign" }
	ENDPOINTFCPUserAuditLogAdd = func(userId string) string { return EndpointFCP + "/users/find/" + userId + "/audit-logs/new" }

	//VASTSIM API Endpoints

	EndpointVATSIMDataFeed = EndpointVATSIMData + "v" + VATSIMDataFeedVersion + "/vatsim-data.json"
)
