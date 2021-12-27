package models

type LoginRequest struct {
	Token string `json:"token,omitempty"`
}
