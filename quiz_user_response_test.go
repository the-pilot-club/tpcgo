package tpcgo

import (
	"encoding/json"
	"testing"
)

func TestQuizUserResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		wantNil  bool
		wantID   string
		wantNick string
		wantErr  bool
	}{
		{
			name:     "member object",
			payload:  `{"user":{"user":{"id":"123","username":"pilot"},"nick":"Ace"}}`,
			wantID:   "123",
			wantNick: "Ace",
		},
		{
			name:     "JSON encoded member string",
			payload:  `{"user":"{\"user\":{\"id\":\"123\",\"username\":\"pilot\"},\"nick\":\"Ace\"}"}`,
			wantID:   "123",
			wantNick: "Ace",
		},
		{
			name:    "null member",
			payload: `{"user":null}`,
			wantNil: true,
		},
		{
			name:    "missing member",
			payload: `{}`,
			wantNil: true,
		},
		{
			name:    "invalid encoded member",
			payload: `{"user":"not JSON"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got QuizUserResponse
			err := json.Unmarshal([]byte(tt.payload), &got)
			if tt.wantErr {
				if err == nil {
					t.Fatal("json.Unmarshal() error = nil, want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if tt.wantNil {
				if got.User != nil {
					t.Fatalf("User = %+v, want nil", got.User)
				}
				return
			}
			if got.User == nil || got.User.User == nil {
				t.Fatalf("User = %+v, want a Discord member and user", got.User)
			}
			if got.User.User.ID != tt.wantID {
				t.Errorf("User.User.ID = %q, want %q", got.User.User.ID, tt.wantID)
			}
			if got.User.Nick != tt.wantNick {
				t.Errorf("User.Nick = %q, want %q", got.User.Nick, tt.wantNick)
			}
		})
	}
}
