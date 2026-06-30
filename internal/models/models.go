package models

import "time"

// CheckIn represents a check-in record
type CheckIn struct {
	ID          int       `json:"id"`
	UserType    string    `json:"user_type"` // student, staff, external
	Name        string    `json:"name"`
	StaffID     string    `json:"staff_id,omitempty"`
	Faculty     string    `json:"faculty,omitempty"`
	Position    string    `json:"position,omitempty"`
	Occupation  string    `json:"occupation,omitempty"`
	CheckInTime time.Time `json:"checkin_time"`
	CreatedAt   time.Time `json:"created_at"`
}

// QuizResult represents a quiz result record
type QuizResult struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Answers          string    `json:"answers"` // JSON string
	ResultType       string    `json:"result_type"`
	ResultDescription string   `json:"result_description"`
	CreatedAt        time.Time `json:"created_at"`
}

// CheckInRequest represents the request body for check-in
type CheckInRequest struct {
	UserType   string `json:"user_type"`
	Name       string `json:"name"`
	StaffID    string `json:"staff_id,omitempty"`
	Faculty    string `json:"faculty,omitempty"`
	Position   string `json:"position,omitempty"`
	Occupation string `json:"occupation,omitempty"`
}

// QuizRequest represents the request body for quiz submission
type QuizRequest struct {
	Name       string         `json:"name"`
	UserType   string         `json:"user_type"`
	StaffID    string         `json:"staff_id,omitempty"`
	Faculty    string         `json:"faculty,omitempty"`
	Position   string         `json:"position,omitempty"`
	Occupation string         `json:"occupation,omitempty"`
	Answers    map[string]int `json:"answers"`
}

// StarPersonality represents a star personality type
type StarPersonality struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Traits      []string `json:"traits"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
}
