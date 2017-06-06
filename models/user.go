package models

import "encoding/json"

// User represents a platform user
type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	GhubID    int    `json:"ghub_id"`
	AuthToken string `json:"auth_token"`
}

// ToByte converts the user in a byte array
func (u *User) ToByte() ([]byte, error) {
	ub, err := json.Marshal(u)
	if err != nil {
		return []byte{}, err
	}
	return ub, nil
}
