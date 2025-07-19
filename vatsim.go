package tpcgo

import (
	"encoding/json"
	"fmt"
)

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
