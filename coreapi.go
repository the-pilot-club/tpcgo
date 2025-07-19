package tpcgo

import "encoding/json"

type NewSuggestion struct {
	StatusID  int    `json:"statusId"`
	DiscordID string `json:"discordId"`
	IdeaText  string `json:"ideaText"`
}

type Suggestions struct {
	ID           string `json:"id,omitempty"`
	StatsID      int    `json:"statsId,omitempty"`
	DiscordId    string `json:"discordId,omitempty"`
	IdeaText     string `json:"ideaText,omitempty"`
	Reason       string `json:"reason,omitempty"`
	MessageID    string `json:"messageId,omitempty"`
	ChannelID    string `json:"channelId,omitempty"`
	ActionUserID string `json:"actionUserId,omitempty"`
}

func (s *Session) GetAllSuggestions() (su []*Suggestions, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPIAllSuggestions, "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) AddSuggestion(entry *NewSuggestion) (su *Suggestions, e error) {
	data, err := s.sendCoreAPIRequest("POST", ENDPOINTCoreAPINewSuggestion, entry)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) GetSuggestion(id string) (su *Suggestions, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPISuggestion(id), "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) UpdateSuggestion(id string, entry *Suggestions) (su *Suggestions, e error) {
	data, err := s.sendCoreAPIRequest("PATCH", ENDPOINTCoreAPISuggestion(id), entry)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) DeleteSuggestion(id string) (d bool, e error) {
	_, err := s.sendCoreAPIRequest("DELETE", ENDPOINTCoreAPISuggestion(id), "")
	if err != nil {
		return false, err
	}
	return true, nil

}
