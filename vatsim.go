package tpcgo

import (
	"encoding/json"
	"fmt"
)

type DataFeed struct {
	General     *GeneralInfo `json:"general,omitempty"`
	Pilots      []Pilot      `json:"pilots,omitempty"`
	Controllers []Controller `json:"controllers,omitempty"`
	ATISs       []ATIS       `json:"atis,omitempty"`
}

type GeneralInfo struct {
	Version          int    `json:"version,omitempty"`
	Reload           int    `json:"reload,omitempty"`
	Update           string `json:"update,omitempty"`
	UpdateTimestamp  string `json:"update_timestamp,omitempty"`
	ConnectedClients int    `json:"connected_clients,omitempty"`
	UniqueUsers      int    `json:"unique_users,omitempty"`
}

type Pilot struct {
	CID            int         `json:"cid,omitempty"`
	Name           string      `json:"name,omitempty"`
	Callsign       string      `json:"callsign,omitempty"`
	Server         string      `json:"server,omitempty"`
	PilotRating    int         `json:"pilot_rating,omitempty"`
	MilitaryRating int         `json:"military_rating,omitempty"`
	Latitude       float64     `json:"latitude,omitempty"`
	Longitude      float64     `json:"longitude,omitempty"`
	Altitude       int         `json:"altitude,omitempty"`
	Groundspeed    int         `json:"groundspeed,omitempty"`
	Transponder    string      `json:"transponder,omitempty"`
	Heading        int         `json:"heading,omitempty"`
	QnhIHg         float64     `json:"qnh_i_hg,omitempty"`
	QnhMb          int         `json:"qnh_mb,omitempty"`
	FlightPlan     *FlightPlan `json:"flight_plan,omitempty"`
	LogonTime      string      `json:"logon_time,omitempty"`
	LastUpdated    string      `json:"last_updated,omitempty"`
}

type FlightPlan struct {
	FlightRules         string `json:"flight_rules,omitempty"`
	Aircraft            string `json:"aircraft,omitempty"`
	AircraftFAA         string `json:"aircraft_faa,omitempty"`
	AircraftShort       string `json:"aircraft_short,omitempty"`
	Departure           string `json:"departure,omitempty"`
	Arrival             string `json:"arrival,omitempty"`
	Alternate           string `json:"alternate,omitempty"`
	CruiseTAS           string `json:"cruise_tas,omitempty"`
	Altitude            string `json:"altitude,omitempty"`
	DepTime             string `json:"deptime,omitempty"`
	EnrouteTime         string `json:"enroute_time,omitempty"`
	FuelTime            string `json:"fuel_time,omitempty"`
	Remarks             string `json:"remarks,omitempty"`
	Route               string `json:"route,omitempty"`
	RevisionID          int    `json:"revision_id,omitempty"`
	AssignedTransponder string `json:"assigned_transponder,omitempty"`
}

type Controller struct {
	CID         int      `json:"cid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Callsign    string   `json:"callsign,omitempty"`
	Frequency   string   `json:"frequency,omitempty"`
	Facility    int      `json:"facility,omitempty"`
	Rating      int      `json:"rating,omitempty"`
	Server      string   `json:"server,omitempty"`
	VisualRange int      `json:"visual_range,omitempty"`
	TextAtis    []string `json:"text_atis,omitempty"`
	LogonTime   string   `json:"logon_time,omitempty"`
}

type ATIS struct {
	CID         int      `json:"cid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Callsign    string   `json:"callsign,omitempty"`
	Frequency   string   `json:"frequency,omitempty"`
	Facility    int      `json:"facility,omitempty"`
	Rating      int      `json:"rating,omitempty"`
	ATISCode    string   `json:"atis_code,omitempty"`
	Server      string   `json:"server,omitempty"`
	VisualRange int      `json:"visual_range,omitempty"`
	TextAtis    []string `json:"text_atis,omitempty"`
	LogonTime   string   `json:"logon_time,omitempty"`
}

type PreFile struct {
	CID        int         `json:"cid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Callsign   string      `json:"callsign,omitempty"`
	FlightPlan *FlightPlan `json:"flight_plan,omitempty"`
	LastUpdate string      `json:"last_updated,omitempty"`
}

type MemberDiscord struct {
	ID        string `json:"id,omitempty"`
	DiscordID string `json:"user_id,omitempty"`
}

type VATSIMUserHours struct {
	ID    int     `json:"id,omitempty"`
	ATC   float64 `json:"atc,omitempty"`
	Pilot float64 `json:"pilot,omitempty"`
	S1    float64 `json:"s1,omitempty"`
	S2    float64 `json:"s2,omitempty"`
	S3    float64 `json:"s3,omitempty"`
	C1    float64 `json:"c1,omitempty"`
	C2    float64 `json:"c2,omitempty"`
	C3    float64 `json:"c3,omitempty"`
	I1    float64 `json:"i1,omitempty"`
	I2    float64 `json:"i2,omitempty"`
	I3    float64 `json:"i3,omitempty"`
	SUP   float64 `json:"sup,omitempty"`
	ADM   float64 `json:"adm,omitempty"`
}

func (s *Session) GetVatsimDataFeed() (df *DataFeed, err error) {
	data, err := s.sendVATSIMRequest("GET", EndpointVATSIMDataFeed, "")
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&df)
	if err != nil {
		return nil, err
	}
	return df, nil
}

func (s *Session) GetUserDiscordID(UserID string) (m *MemberDiscord, err error) {
	data, err := s.sendVATSIMRequest("GET", EndpointVATSIMDiscordId(UserID), "")
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func (s *Session) GetUserHours(UserID string) (m *VATSIMUserHours, err error) {
	data, err := s.sendVATSIMRequest("GET", EndpointVATSIMUserHours(UserID), "")
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, err
}
