package tpcgo

import (
	"encoding/json"
	"fmt"
)

type NotFoundResource struct {
	Detail string `json:"detail,omitempty"`
}

type AuditLogEntry struct {
	UserId      string `json:"user_id,omitempty"`
	SubmittedBy string `json:"submitted_by,omitempty"`
	Text        string `json:"text,omitempty"`
}

type FCPAuditLogAdd struct {
	UserID      string `json:"user_id"`
	SubmittedBy string `json:"submitted_by"`
	Text        string `json:"text"`
}

type AuditLogResponse struct {
	Success string        `json:"success,omitempty"`
	Entry   AuditLogEntry `json:"entry,omitempty"`
}

type FCPCallsignResponse struct {
	Callsign int `json:"tpcCallsign,omitempty"`
}

type FCPUser struct {
	ID             string            `json:"id,omitempty"`
	DiscordID      string            `json:"discordId,omitempty"`
	Callsign       int               `json:"callsign,omitempty"`
	VATSIMCid      int               `json:"vatsimCid,omitempty"`
	HomeAirport    string            `json:"homeAirport,omitempty"`
	ChartersCode   string            `json:"chartersCode,omitempty"`
	Bio            string            `json:"bio,omitempty"`
	AircraftHangar []FCPUserAircraft `json:"aircraftHangar,omitempty"`
}

type FCPUserAircraft struct {
	ICAO string `json:"icao,omitempty"`
	Name string `json:"aircraftName,omitempty"`
}

func GetFCPCallsign(UserId string, UserAgent string) (c *FCPCallsignResponse, e error) {

	url := fmt.Sprintf("https://flightcrew.thepilotclub.org/api/users/find/%v/callsign", UserId)
	data, err := sendRequest("GET", url, "", "", "", UserAgent)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetFCPUser(UserId string, UserAgent string) (u *FCPUser, e error) {
	url := fmt.Sprintf("https://flightcrew.thepilotclub.org/api/users/find/%v", UserId)
	data, err := sendRequest("GET", url, "", "", "", UserAgent)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func PostAuditLogEntry(UserId string, APIKeyHeader string, APIKey string, UserAgent string, entry AuditLogEntry) (a *AuditLogResponse, e error) {

	url := fmt.Sprintf("https://flightcrew.thepilotclub.org/api/users/find/%v/audit-logs/new", UserId)
	data, err := sendRequest("POST", url, entry, APIKeyHeader, APIKey, UserAgent)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
