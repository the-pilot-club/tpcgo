package tpcgo

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

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

type QuizQuestion struct {
	ID                string `json:"id,omitempty"`
	Question          string `json:"question,omitempty"`
	OptionA           string `json:"optionA,omitempty"`
	OptionB           string `json:"optionB,omitempty"`
	OptionC           string `json:"optionC,omitempty"`
	CorrectAnswer     string `json:"correctAnswer,omitempty"`
	AnswerDescription string `json:"answerDescription,omitempty"`
	IsActive          bool   `json:"isActive,omitempty"`
	IsOld             bool   `json:"isOld,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
}

type QuizQuestionForResponse struct {
	MessageID  string `json:"messageId,omitempty"`
	QuestionID string `json:"questionId,omitempty"`
}

type QuizUserResponseSet struct {
	MessageID string            `json:"messageId,omitempty"`
	UserID    string            `json:"userId,omitempty"`
	Answer    string            `json:"answer,omitempty"`
	User      *discordgo.Member `json:"user,omitempty"`
}

type QuizUserResponse struct {
	User *discordgo.Member `json:"user,omitempty"`
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

func (s *Session) GetAllQuizQuestions() (su []*QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPIAllQuizQuestions, "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&su)
	if err != nil {
		return nil, err
	}
	return su, nil
}

func (s *Session) GetCurrentQuizQuestions() (q *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPICurrentQuizQuestion, "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Session) GetQuizQuestionByID(id string) (q *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("GET", EndPointCoreAPIQuizQuestionByID(id), "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Session) GetNextQuizQuestion() (q *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPINextQuizQuestion, "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Session) AddNewQuizQuestion(i QuizQuestion) (q *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("POST", ENDPOINTCoreAPINewQuizQuestion, i)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Session) UpdateQuizQuestion(id string, i QuizQuestion) (q *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("PUT", EndPointCoreAPIQuizQuestionByID(id), i)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&q)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (s *Session) DeleteQuizQuestion(id string) (d *QuizQuestion, e error) {
	data, err := s.sendCoreAPIRequest("DELETE", EndPointCoreAPIQuizQuestionByID(id), "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Session) SetQuestionForResponse(messageID string, questionID string) (c bool, e error) {
	data, err := s.sendCoreAPIRequest("POST", ENDPOINTCoreAPIQuizSetQuestionForResponse, QuizQuestionForResponse{MessageID: messageID, QuestionID: questionID})
	if err != nil {
		return false, err
	}
	if data.StatusCode != 201 {
		return false, nil
	}
	return true, nil
}

func (s *Session) SetQuizUserResponse(i *QuizUserResponseSet) (c bool, e error) {
	data, err := s.sendCoreAPIRequest("POST", ENDPOINTCoreAPIQuizSetUserResponse, i)
	if err != nil {
		return false, err
	}
	if data.StatusCode != 201 {
		return false, nil
	}
	return true, nil
}

func (s *Session) DeleteQuizQuestionForResponse(id string) (d bool, e error) {
	data, err := s.sendCoreAPIRequest("DELETE", ENDPOINTCoreAPIQuizDeleteQuestionForResponse(id), "")
	if err != nil {
		return false, err
	}
	if data.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

func (s *Session) GetQuizUserResponses(id string, answer string) (r []*QuizUserResponse, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPIQuizGetResponses(id, answer), "")
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(data.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Session) ResetQuizUserResponses() (d bool, e error) {
	data, err := s.sendCoreAPIRequest("DELETE", ENDPOINTCoreAPIResetUserResponses, "")
	if err != nil {
		return false, err
	}
	if data.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}
func (s *Session) CheckUserQuizResponse(id string) (r bool, e error) {
	data, err := s.sendCoreAPIRequest("GET", ENDPOINTCoreAPICheckUserResponse(id), "")
	if err != nil {
		return false, err
	}
	err = json.NewDecoder(data.Body).Decode(&r)
	if err != nil {
		return false, err
	}
	return r, nil
}
