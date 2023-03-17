package models

type User struct {
	Id           int    `json:"id"`
	SocialSource string `json:"socialSource"`
	SocialUserId string `json:"socialUserId"`
	Username     string `json:"username"`
	AvatarUrl    string `json:"avatarUrl"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	LastLogin    MyTime `json:"lastLogin"`
}
