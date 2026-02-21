package tpcgo

import (
	"encoding/json"
	"fmt"
)

// GetFCPCallsign fetches a users TPC Callsign based on the Discord ID provided.
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

// GetAllFCPUsers fetches all FCP users (limited user fields) from the configured FCP environment.
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

// GetAllFCPUsersCID fetches only VATSIM CIDs for all FCP users from the configured FCP environment.
func (s *Session) GetAllFCPUsersCID() (u []*FCPCIDOnly, e error) {
	data, err := s.sendFCPRequest("GET", ENDPOINTFCPGetAllUsersCID(s.FCPSession.Environment), "")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.NewDecoder(data.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetFCPUser fetches a single FCP user (limited user fields) by user ID from the configured FCP environment.
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

// AddFCPUser creates/adds a new user in FCP (requires an FCP API key).
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

// DeleteFCPUser removes a user from FCP by user ID (requires an FCP API key).
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

// GetFCPUsersBirthdays fetches the FCP birthdays list (requires an FCP API key).
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

// PostAuditLogEntry creates a new audit log entry for the given user ID in FCP (requires an FCP API key).
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

// GetAllFBOs fetches all FBOs from the configured FCP environment.
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

// GetFBO fetches a single FBO by ICAO from the configured FCP environment.
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

// GetSectors fetches all sectors from the configured FCP environment.
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
