package tpcgo

import (
	"encoding/json"
	"fmt"
)

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
