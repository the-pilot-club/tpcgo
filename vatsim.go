package tpcgo

import (
	"encoding/json"
	"fmt"
)

// GetVatsimDataFeed fetches the latest VATSIM datafeed and decodes it into a DataFeed struct.
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

// GetUserDiscordID looks up a VATSIM member by Discord ID and returns the mapped member info.
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

// GetUserHours fetches a VATSIM member's hours/stats by VATSIM CID.
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
