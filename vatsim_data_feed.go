package tpcgo

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrJSONUnmarshal = errors.New("json unmarshal")
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
	QNH_I_Hg       float64     `json:"qnh_i_hg,omitempty"`
	QNH_MB         int         `json:"qnh_mb,omitempty"`
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
	Frequency   float64  `json:"frequency,omitempty"`
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
	Frequency   float64  `json:"frequency,omitempty"`
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

func GetVatsimDataFeed() (df *DataFeed, err error) {
	data, err := sendRequest("GET", "https://data.vatsim.net/v3/vatsim-data.json", "", "", "", "")
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSONUnmarshal, err)
	}
	err = json.NewDecoder(data.Body).Decode(&df)
	if err != nil {
		return nil, err
	}
	return df, nil
}
