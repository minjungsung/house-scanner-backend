package models

type User struct {
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	Email          string   `json:"email"`
	Address        string   `json:"address"`
	Message        string   `json:"message"`
	ReferralSource []string `json:"referralSource"`
	CreatedAt      string   `json:"created_at,omitempty"`
}
