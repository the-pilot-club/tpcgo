package tpcgo

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// UnmarshalJSON accepts both CoreAPI response formats: a member encoded as a
// JSON object and a member encoded as a JSON string.
func (r *QuizUserResponse) UnmarshalJSON(data []byte) error {
	var payload struct {
		User json.RawMessage `json:"user"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if len(payload.User) == 0 || string(payload.User) == "null" {
		r.User = nil
		return nil
	}

	memberJSON := payload.User
	if payload.User[0] == '"' {
		var encoded string
		if err := json.Unmarshal(payload.User, &encoded); err != nil {
			return fmt.Errorf("decode quiz user response member string: %w", err)
		}
		memberJSON = []byte(encoded)
	}

	var member discordgo.Member
	if err := json.Unmarshal(memberJSON, &member); err != nil {
		return fmt.Errorf("decode quiz user response member: %w", err)
	}

	r.User = &member
	return nil
}
