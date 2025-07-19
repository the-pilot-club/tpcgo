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

type SuccessResponse struct {
	Success string `json:"success,omitempty"`
}

type FCPCallsignResponse struct {
	Callsign int `json:"tpcCallsign,omitempty"`
}

type FCPFullUser struct {
	ID              string            `json:"id"`
	DiscordID       string            `json:"discordId,omitempty"`
	DiscordUsername string            `json:"discordUsername,omitempty"`
	FirstName       string            `json:"firstName,omitempty"`
	LastName        string            `json:"lastName,omitempty"`
	Callsign        int               `json:"callsign,omitempty"`
	VATSIMCid       int               `json:"vatsimCid,omitempty"`
	Email           string            `json:"email,omitempty"`
	HomeAirport     string            `json:"homeAirport,omitempty"`
	ChartersCode    string            `json:"chartersCode,omitempty"`
	Bio             string            `json:"bio,omitempty"`
	Status          string            `json:"status,omitempty"`
	AircraftHangar  []FCPUserAircraft `json:"aircraftHangar,omitempty"`
}

type FCPLimitedUser struct {
	ID             string            `json:"id,omitempty"`
	DiscordID      string            `json:"discordId,omitempty"`
	Callsign       int               `json:"callsign,omitempty"`
	VATSIMCid      int               `json:"vatsimCid,omitempty"`
	HomeAirport    string            `json:"homeAirport,omitempty"`
	ChartersCode   string            `json:"chartersCode,omitempty"`
	Bio            string            `json:"bio,omitempty"`
	AircraftHangar []FCPUserAircraft `json:"aircraftHangar,omitempty"`
}

type FCPUserAdd struct {
	UserID string `json:"id"`
}

type FCPUserAircraft struct {
	ICAO string `json:"icao,omitempty"`
	Name string `json:"aircraftName,omitempty"`
}

type FCPFBOs struct {
	FBO []FBO `json:"fbos,omitempty"`
}
type FCPSectors struct {
	Sectors []Sectors `json:"sectors,omitempty"`
}
type FBO struct {
	ICAO        string    `json:"icao,omitempty"`
	Name        string    `json:"name,omitempty"`
	Region      string    `json:"region,omitempty"`
	Fuel        string    `json:"fuel,omitempty"`
	Focus       string    `json:"focus,omitempty"`
	PublicNotes string    `json:"publicNotes,omitempty"`
	Sectors     []Sectors `json:"sectors,omitempty"`
}

type Sectors struct {
	StartICAO    string `json:"startIcao,omitempty"`
	EndICAO      string `json:"endIcao,omitempty"`
	Size         string `json:"size,omitempty"`
	SectorNumber int    `json:"sectorNumber,omitempty"`
}

func (s *Session) GetFCPCallsign(UserId string) (c *FCPCallsignResponse, e error) {
	data, err := s.sendFCPRequest("GET", ENDPOINTFCPUserCallsign(UserId, s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Session) GetAllFCPUsers() (u []*FCPLimitedUser, e error) {
	data, err := s.sendFCPRequest("GET", ENDPOINTFCPGetAllUsers(s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Session) GetFCPUser(UserId string) (u *FCPLimitedUser, e error) {
	data, err := s.sendFCPRequest("GET", ENDPOINTFCPUser(UserId, s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Session) AddFCPUser(UserID *FCPUserAdd) (a *FCPFullUser, e error) {
	if s.FCPSession.ApiKey == "" {
		return nil, ErrNoKeyError
	}
	data, err := s.sendFCPRequest("POST", ENDPOINTFCPUserAdd(s.FCPSession.Environment), UserID)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *Session) DeleteFCPUser(UserID string) (su *SuccessResponse, e error) {
	if s.FCPSession.ApiKey == "" {
		return nil, ErrNoKeyError
	}
	data, err := s.sendFCPRequest("DELETE", ENDPOINTFCPUserDelete(UserID, s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) GetFCPUsersBirthdays() (u []*string, e error) {
	if s.FCPSession.ApiKey == "" {
		return nil, ErrNoKeyError
	}
	data, err := s.sendFCPRequest("GET", ENDPOINTFCPUserBirthdays(s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Session) PostAuditLogEntry(UserId string, entry AuditLogEntry) (a *AuditLogResponse, e error) {
	if s.FCPSession.ApiKey == "" {
		return nil, ErrNoKeyError
	}
	data, err := s.sendFCPRequest("POST", ENDPOINTFCPUserAuditLogAdd(UserId, s.FCPSession.Environment), entry)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Session) GetAllFBOs() (f *FCPFBOs, e error) {
	data, err := s.sendFCPRequest("GET", EndPointFCPALLFBOs(s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Session) GetFBO(Icao string) (f *FBO, e error) {
	data, err := s.sendFCPRequest("GET", EndPointFCPFBO(Icao, s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Session) GetSectors() (se *FCPSectors, e error) {
	data, err := s.sendFCPRequest("GET", EndPointFCPSectors(s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&se)
	if err != nil {
		return nil, err
	}
	return se, nil
}
