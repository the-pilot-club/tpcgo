package tpcgo

import (
	"encoding/json"
)

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
