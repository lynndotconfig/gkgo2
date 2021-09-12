package entity

type UserEntity struct {
	ID uint64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}
