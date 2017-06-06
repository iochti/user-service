package models

type User struct {
	ID        int
	Login     string
	AvatarURL string
	Name      string
	GhubID    int
	AuthToken string
}
